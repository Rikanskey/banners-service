package main

import "banners-service/internal/runner"

const configDir = "./config/"

func main() {
	runner.Start(configDir)
}
