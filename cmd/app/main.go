package main

import "github.com/andrew-nino/atm_v1/internal/app"

const pathConfig = "config/config.yaml"

func main() {
	app.Run(pathConfig)
}
