package main

import (
	"fmt"
	"time"
)

func logging() {
	fmt.Println("Logging")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
	time.Sleep(2 * time.Second)
	fmt.Println("Application Finished")
}
func main() {
	runApplication()
}
