// lesson_interface.go
// Demonstrates the use of interfaces, structs, and methods in Go,
// with a focus on custom bracket formatting for visual clarity.

package main

import "fmt"

// Shaka represents a user with both public and private fields.
// The struct uses JSON tags for serialization and custom bracket formatting.
type Shaka struct
{
    // S_name is a public field representing the user's name.
    S_name     string `json:"proper_name"`
    // s_password is a private field representing the user's password.
    s_password string `json:"hidden_password"`
}

// Serf represents a user with only a public field.
// The struct uses custom bracket formatting for visual consistency.
type Serf struct
{
    // S_name is a public field representing the user's name.
    S_name string `json:"proper_name"`
}

// Greeter defines an interface for types that can greet.
// The interface uses custom bracket formatting for visual clarity.
type Greeter interface
{
    // Greet is a method that prints a greeting, optionally with verbose details.
    Greet(verbose bool)
}

// Greet implements the Greeter interface for Shaka.
// The method uses nested braces for ANSI/C-style visual structure.
func (s Shaka) Greet(verbose bool) {
{
    // If verbose is true, print the name and password.
    if verbose {
    {
        fmt.Printf("Hi %s, your password is %s\n", s.S_name, s.s_password)
	// Otherwise, print only the name.
    }} else {
    {
        fmt.Printf("Hi %s\n", s.S_name)
    }}
}}

// Greet implements the Greeter interface for Serf.
// The method uses nested braces for ANSI/C-style visual structure.
func (s Serf) Greet(verbose bool) {
{
    // Print the name of the Serf.
    fmt.Printf("Hi %s\n", s.S_name)
}}

func main() {
{
    // Create a slice to hold Greeter objects.
    list := make([]Greeter, 0)

    // Initialize a Shaka instance with a name and password.
    st_king := Shaka{S_name: "Kemehameah", s_password: "secret"}
    // Initialize a Serf instance with a name.
    st_serf := Serf{S_name: "Umbe"}

    // Append both instances to the Greeter slice.
    list = append(list, st_king)
    list = append(list, st_serf)

    // Iterate over the Greeter slice and call Greet on each object.
    for _, greeter := range list {
    {
        // Call Greet with verbose set to false.
        greeter.Greet(false)
    }}
}}
