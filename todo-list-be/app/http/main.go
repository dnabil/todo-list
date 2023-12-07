package main

import "todo-list-be/config"

func main() {
	// dependencies (configs, log, db, app):
	config.LoadConfig() // .env (dev only)
	log := config.NewLogger()
	app := config.NewGin()
	sql := config.NewSQL(log)

	config.Bootstrap(&config.BootstrapConfig{
		DB: sql,
		App: app,
		Log: log,
	})

	if err := app.Run(); err != nil{
		log.Fatalln("app failed to run :", err)
	}
}