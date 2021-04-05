package main

import "fmt"

func main() {
	app := App{}

	//TODO use fmt.Printf for dynamic variables
	connStr := `port=5432 user=postgres dbname=postgres sslmode=disable password=superpass@123`

	err := app.Initialize(connStr)
	if err != nil {
		fmt.Printf("error initializing: %v\n", err)
		panic("application failed to initialize")
	}
	defer app.Db.Close()

	app.Run(":3000")
}
