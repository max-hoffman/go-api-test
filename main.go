package main

import (
	"fmt"
	"os"
)

func main() {
	a := App{}
	// a.Initialize(
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"))

	a.Initialize("postgres", "", "reactapp")

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	a.Run(address)
}
