package main

import (
	initialize2 "meng-admin-gin/initialize"
)

func main() {

	initialize2.InitViper()
	//initialize.InitRoute()

	initialize2.RunServer()
}
