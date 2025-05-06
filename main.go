package main

import (
	"ucc-gorm/db"
	"ucc-gorm/app"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}