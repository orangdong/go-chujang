package main

import (
	"fmt"

	"github.com/orangdong/go-chujang/config"
)

func main() {
	err := config.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
	fmt.Println("App is running")
}
