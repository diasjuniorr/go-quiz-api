package main

import "fmt"

func main() {
	app := App{}

	err := app.Initialize("teste")
	if err != nil {
		fmt.Printf("error initializing: %v\n", err)
		panic("application failed to initialize")
	}
	defer app.Db.Close()

	app.Run(":3000")
}
