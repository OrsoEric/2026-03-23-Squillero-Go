//go run lesson_interface.go

package main

import "fmt"

// Shaka is a struct with a public and a private field
type Shaka struct
{
	S_name     string `json:"proper_name"`
	s_password string `json:"hidden_password"`
}

// Serf is a struct with a public field
type Serf struct
{
	S_name string `json:"proper_name"`
}

// Greeter is an interface with a Greet method
type Greeter interface
{
	Greet(verbose bool)
}

// Greet method for Shaka
func (s Shaka) Greet(verbose bool) {
{
	if verbose {
	{
		fmt.Printf("Hi %s, your password is %s\n", s.S_name, s.s_password)
	}} else {
	{
		fmt.Printf("Hi %s\n", s.S_name)
	}}
}}

// Greet method for Serf
func (s Serf) Greet(verbose bool) {
{
	fmt.Printf("Hi %s\n", s.S_name)
}}

func main() {
{
	list := make([]Greeter, 0)

	st_king := Shaka{S_name: "Kemehameah", s_password: "secret"}
	st_serf := Serf{S_name: "Umbe"}

	list = append(list, st_king)
	list = append(list, st_serf)

	for _, greeter := range list {
	{
		greeter.Greet(false)
	}}
}}
