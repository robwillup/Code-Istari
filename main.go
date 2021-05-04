// This app simply displays 'Hello world'
package main

import (
	"fmt"

	"golang.org/x/mobile/app"

)

func main() {
	app.Main(func(a app.App) {
		fmt.Println("Hello world")
	})
}