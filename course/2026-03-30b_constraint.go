//go mod init 2026-03-30b_constraint.go
//go get "golang.org/x/exp/constraints"
//go run course/2026-03-30b_constraint.go

package main

import 
(
	"fmt"
	"golang.org/x/exp/constraints"
)

func my_function[T constraints.Integer | constraints.Float](value T) {
{
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}}

func main() {
{
   	my_function(42)
	my_function(3.14)
}}
