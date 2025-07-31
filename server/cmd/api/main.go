package main

import (
	"fmt"

	"github.com/Samrat-Mukherjee/file_share/internals/api"
)

func main() {
	router := api.NewRouter()
	fmt.Println("port is running on :8080")
	router.Run(":8080")
}
