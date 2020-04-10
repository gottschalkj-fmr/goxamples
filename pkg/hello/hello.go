package hello

import (
	"fmt"
	"time"
)

// Greeting of Hello World in Golang
func Greeting() string {
	return fmt.Sprintf("Hello, World! %s", time.Now())
}
