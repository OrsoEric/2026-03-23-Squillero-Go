
//go run lesson_error.go


package main

import "fmt"



func div( i_n_num, i_n_den int ) ( float32, error ) {
{
	if i_n_den == 0 {
	{
		return 0.0, fmt.Errorf("divide by 0")
	}}

	return float32(i_n_num)/float32(i_n_den), nil
}}

func main() {
{
	fmt.Println( div(0,0) )

	fmt.Println( div(3,2) )

}}