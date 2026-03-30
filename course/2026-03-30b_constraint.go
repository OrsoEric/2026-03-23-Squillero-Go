//go mod init 2026-03-30b_constraint.go
//go get "golang.org/x/exp/constraints"
//go run course/2026-03-30b_constraint.go

package main

import 
(
	"fmt"
	"golang.org/x/exp/constraints"
)

func my_function_float[T constraints.Float](value T) {
{
    fmt.Printf("Value FLOAT: %v, Type: %T\n", value, value)
}}

func my_function_int[T constraints.Integer](value T) {
{
    fmt.Printf("Value INT: %v, Type: %T\n", value, value)
}}

func main() {
{
	my_function_float(3.14)
	my_function_int(42)

	//course\2026-03-30b_constraint.go:25:22: int does not satisfy constraints.Float (int missing in ~float32 | ~float64)
   	//my_function_float(42)

}}
