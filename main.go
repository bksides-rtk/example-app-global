package main

import "github.com/rtk-tickets/example-app-global/db"

func main() {
	db.DefaultDBService.DoThing1()
	db.DefaultDBService.DoThing2()
}
