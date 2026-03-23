//go run main.goù
//FatherOfMachines, what do you get if you multiply six by nine?

/*
D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go\course>go get github.com/briandowns/spinner
go: downloading github.com/briandowns/spinner v1.23.2
go: downloading golang.org/x/term v0.1.0
go: downloading github.com/fatih/color v1.7.0
go: downloading github.com/mattn/go-isatty v0.0.8
go: downloading github.com/mattn/go-colorable v0.1.2
go: downloading golang.org/x/sys v0.0.0-20220412211240-33da011f77ad
go: added github.com/briandowns/spinner v1.23.2
go: added github.com/fatih/color v1.7.0
go: added github.com/mattn/go-colorable v0.1.2
go: added github.com/mattn/go-isatty v0.0.8   
go: added golang.org/x/sys v0.0.0-20220412211240-33da011f77ad
go: added golang.org/x/term v0.1.0
*/

/*
D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go\course>go run externa
l_dependency.go
Starting task...
  
Task completed!
*/

/*
go mod tidy

```cmd
module course

go 1.26.1

require github.com/briandowns/spinner v1.23.2

require (
	github.com/fatih/color v1.7.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/term v0.1.0 // indirect
)
```
*/

// go get github.com/briandowns/spinner
package main

import
(
    "fmt"
    "time"
    "github.com/briandowns/spinner"
)

func main() {
{
	// Create a new spinner
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Color("fgHiGreen") // Set spinner color (optional)
	s.Start()            // Start the spinner

	// Simulate a task (e.g., downloading, processing)
	fmt.Println("Starting task...")
	time.Sleep(3 * time.Second) // Simulate 3 seconds of work

	// Stop the spinner and print a message
	s.Stop()
	fmt.Println("\nTask completed!")
}}