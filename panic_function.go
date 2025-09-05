// Package declaration - defines this file as part of the main package
// which creates an executable program
package main

// Import statement - brings in the fmt package for formatted I/O operations
// fmt provides functions like Println for console output
import "fmt"

// endApp function - a cleanup function that will be called when the program ends
// This function has no parameters and no return values
func endApp() {
	// Print a message to indicate the application is ending
	// This will help us see when cleanup occurs
	fmt.Println("End App")

	message := recover()

	fmt.Println("Terjadi error:", message)
	fmt.Println("App Running")
}

// runApp function - simulates running an application with potential error handling
// Parameters:
//   - error: a boolean flag that determines whether to trigger a panic
func runApp(error bool) {

	// Note: This function has no explicit return statement
	// It will return normally (with no values) when it reaches the end
	defer endApp()

	// Conditional check - if error is true, trigger a panic
	// panic() stops normal execution and begins panicking
	// The string "Error" will be the panic message
	if error {
		panic("Error")
	}

	// defer statement - schedules endApp() to run when runApp() exits
	// Key points about defer:
	// - Executes even if a panic occurs (before the panic propagates)
	// - Runs in LIFO (Last In, First Out) order if multiple defers exist
	// - Executes after return statements but before the function actually returns

}

// main function - the entry point of the program
// Every executable Go program must have a main function
func main() {
	// Call runApp with false - this will NOT trigger a panic
	// Expected execution flow:
	// 1. runApp(false) is called
	// 2. error check fails (false), so no panic
	// 3. defer endApp() is scheduled
	// 4. runApp() completes normally
	// 5. defer executes: "End App" is printed
	// 6. Program exits normally
	runApp(true)
	fmt.Println("Program selesai")
}
