/*
 * Copyright (c) 2020 krautbax.
 * Licensed under the Apache License, Version 2.0
 * http://www.apache.org/licenses/LICENSE-2.0
 */

package hello

import (
	"fmt"
	"time"
)

// Greeting of Hello World in Golang
func Greeting() string {
	return fmt.Sprintf("Hello, World! %s", time.Now())
}
