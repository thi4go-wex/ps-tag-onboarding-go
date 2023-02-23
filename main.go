package main

import (
	"ps-tag-onboarding-go/server"
)

func main() {
	sv := server.New()

	sv.Run()
}
