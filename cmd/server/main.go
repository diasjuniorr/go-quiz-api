package main

func main() {
	app := App{}

	err := app.Initialize("teste")
	if err != nil {
		panic("application failed to initialize")
	}
	defer app.Db.Close()

	app.Run(":3000")
}
