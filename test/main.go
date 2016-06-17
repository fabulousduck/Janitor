package main

import (
	janitor "github.com/fabulousduck/janitor"
)

func main() {
	jan := janitor.NewJanitor()

	jan.CleanDir("/Users/ryanvlaming/Desktop")

}
