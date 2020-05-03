#
# Copyright (c) 2020 krautbax.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
image := github.com/krautbax/goxamples
image_tags := alpine buster amazonlinux
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
	@for image in $(runnable_images); do \
		if [ -n "$$(docker images $$image -q)" ]; then \
			command="(docker run --publish $(published_port):$(exposed_port) $$image &) && sleep 2 && curl http://localhost:$(published_port)"; \
			echo $$command && eval $$command && echo; \
		fi; \
	done

$(image_tags):
	@if [ -z "$$(docker images $(image):$@ -q)" ]; then \
		command="docker build $(build_options) --file build/$@/Dockerfile --tag $(image):$@ ."; \
		echo $$command && eval $$command && echo; \
	fi; \

clean:
	-docker stop $$(docker ps -q) 2>/dev/null
	docker container prune --force
	docker image prune --force

clobber: clean
	docker rmi --force $(runnable_images)
