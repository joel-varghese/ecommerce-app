package main

import (
	"api_ecom/app"
)

func main() {
	App := &app.App{}
	App.Initialize()
	App.Run(":3000")

}