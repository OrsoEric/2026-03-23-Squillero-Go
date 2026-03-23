// go run demo_array.go

package main

import "fmt"

func main() {
{
   	// Array
    var as_teletubbies_a = [4]string{"poo", "moo", "a", "dii"}
    fmt.Println(as_teletubbies_a)
    fmt.Println("Element: ", as_teletubbies_a[0])
    fmt.Println("Capacity: ", cap(as_teletubbies_a))
    fmt.Printf("TYPE: %T\n", as_teletubbies_a)

    // Slice
    as_teletubbies_b := []string{"poo", "moo", "a", "dii"}
    fmt.Println(as_teletubbies_b)
    fmt.Println("Element: ", as_teletubbies_b[0])
    fmt.Println("Capacity: ", cap(as_teletubbies_b))

    // Append to slice
    as_teletubbies_b = append(as_teletubbies_b, "mee", "raa")
	//as_teletubbies_b = append( as_teletubbies_b, [...]string{"mee", "raa"} )
    fmt.Println(as_teletubbies_b)

}}