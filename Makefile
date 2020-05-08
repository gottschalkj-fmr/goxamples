#
# Copyright (c) 2020 krautbax.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
image := github.com/krautbax/goxamples
image_tags := alpine debian ubuntu amazonlinux
runnable_images := $(image_tags:%=$(image)\:%)
published_port := 9090
exposed_port := 9090

module := github.com/krautbax/goxamples
build_args := MODULE=$(module)
build_options := --no-cache --force-rm $(build_args:%=--build-arg %)

.PHONY: all images test clean clobber $(image_tags)

all: images test

images: $(image_tags)

test:
	@port=$(published_port); \
	for image in $(runnable_images); do \
		if [ -n "$$(docker images $$image -q)" ]; then \
			tries=0; \
			command="docker run --publish $$port:$(exposed_port) $$image"; echo $$command && eval "($$command &)"; \
			while [ $$tries -lt 5 ] && [ $$(curl -o /dev/null -s -w "%{http_code}" http://localhost:$$port/health) -ne 200 ]; do \
				tries=$$(expr $$tries + 1); \
				sleep 1; \
			done; \
			if [ $$tries -eq 5 ]; then exit 1; fi; \
			greeting=$$(curl -s http://localhost:$$port/greet); \
			port=$$(expr $$port + 1); echo $$greeting && echo; \
		fi; \
	done

# Using date command from BSD UNIX.
# You can always just check to see if the image exists and skip the modification check on the Dockerfile.
# No secondary expansion and no dependency on the Dockerfile.
# i.e. @if [ -z "$$(docker images $(image):$@ -q)" ]; then \ ...
.SECONDEXPANSION:
$(image_tags): build/$$@/Dockerfile
	@dockerfile_epochtime=$$(date -j $$(date -ur $< +%m%d%H%M%Y.%S) +%s); \
	image_epochtime=$$(date -j -f "%FT%T" $$(docker image inspect --format "{{.Created}}" $(image):$@ 2>/dev/null) +%s 2>/dev/null); \
	if [ -z "$$image_epochtime" ] || [ $$dockerfile_epochtime -gt $$image_epochtime ]; then \
		command="docker build $(build_options) --file $< --tag $(image):$@ ."; \
		echo $$command && eval $$command && echo; \
	fi

clean:
	-docker stop $$(docker ps -q) 2>/dev/null
	docker container prune --force
	docker image prune --force

clobber: clean
	docker rmi --force $(runnable_images)
