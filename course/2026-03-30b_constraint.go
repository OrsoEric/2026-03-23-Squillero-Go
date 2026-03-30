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



func my_function[T constraints.Float | constraints.Integer](value T) {
{
	switch v := any(value).(type) {
	case float32:
	{
		my_function_float(v)
	}
	case float64:
	{
		my_function_float(v)
	}
	case int:
	{
		my_function_int(v)
	}
	case int8:
	{
		my_function_int(v)
	}
	case int16:
	{
		my_function_int(v)
	}
	case int32:
	{
		my_function_int(v)
	}
	case int64:
	{
		my_function_int(v)
	}
	case uint:
	{
		my_function_int(v)
	}
	case uint8:
	{
		my_function_int(v)
	}
	case uint16:
	{
		my_function_int(v)
	}
	case uint32:
	{
		my_function_int(v)
	}
	case uint64:
	{
		my_function_int(v)
	}
	case uintptr:
	{
		my_function_int(v)
	}
	}
}}

func main() {
{
	my_function_int(42)
	my_function_float(3.14)

	//course\2026-03-30b_constraint.go:25:22: int does not satisfy constraints.Float (int missing in ~float32 | ~float64)
   	//my_function_float(42)

	my_function(42)
	my_function(3.14)
}}
