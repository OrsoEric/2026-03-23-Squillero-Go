//go run func_that_makes_func.go

package main

import "fmt"

// fn_make_inc creates and returns a closure that "remembers" the initial offset.
// Each time the returned function is called, the offset is doubled and added to the input.
// This demonstrates how closures in Go capture and retain the state of their outer scope.
//
// Parameters:
//   i_n_offset int: The initial offset value to be remembered and modified.
//
// Returns:
//   func(int) int: A function that takes an integer and returns the sum of the doubled offset and the input.
func fn_make_inc(i_n_offset int) func(int) int {
{
    // The returned function is a closure: it captures and retains the `i_n_offset` variable.
    return func(i_n int) int {
	{
        // Each call doubles the captured `i_n_offset` and adds the input `i_n`.
        i_n_offset *= 2
        return i_n_offset + i_n
    }}
}}

func main() {
{
    // Create a closure with an initial offset of 10.
    // The closure "remembers" this offset between calls.
    fn_remember_ten := fn_make_inc(10)

    // First call: offset is 10, doubled to 20, then 20 + 1 = 21
    fmt.Println(fn_remember_ten(1))
    // Second call: offset is now 20, doubled to 40, then 40 + 1 = 41
    fmt.Println(fn_remember_ten(1))

    // Create another closure with an initial offset of 100.
    // This closure is independent and has its own captured offset.
    fn_remember_hundred := fn_make_inc(100)

    // First call: offset is 100, doubled to 200, then 200 + 1 = 201
    fmt.Println(fn_remember_hundred(1))
    // Second call: offset is now 200, doubled to 400, then 400 + 1 = 401
    fmt.Println(fn_remember_hundred(1))
}}
