//go run func_that_makes_func.go

package main

import "fmt"

func fn_make_inc(i_n_offset int) func(int) int {
{
    return func(i_n int) int {
	{
        i_n_offset *= 2
        return i_n_offset + i_n
    }}
}}

func main() {
{
    fn_remember_ten := fn_make_inc(10)

    fmt.Println(fn_remember_ten(1)) 
	fmt.Println(fn_remember_ten(1)) 

    fn_remember_hundred := fn_make_inc(100)

    fmt.Println(fn_remember_hundred(1)) 
	fmt.Println(fn_remember_hundred(1)) 
}}
