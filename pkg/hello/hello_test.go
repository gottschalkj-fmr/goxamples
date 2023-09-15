/*
 * Copyright (c) 2020 gottschalkj-fmr.
 * Licensed under the Apache License, Version 2.0
 * http://www.apache.org/licenses/LICENSE-2.0
 */

package hello

import (
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	if g := Greeting(); !strings.HasPrefix(g, "Hello, World!") {
		t.Errorf("Greeting() is %v", g)
	}
}
