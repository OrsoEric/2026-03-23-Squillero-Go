//go run lesson_panic_recover.go

package main

import "fmt"



func div( i_n_num, i_n_den int ) ( float32 ) {
{
	if i_n_den == 0 {
	{
		panic("Divisionbyzero")
	}}

	return float32(i_n_num)/float32(i_n_den)
}}

func trap() {
{
	if xxx := recover(); xxx != nil{
	{
		fmt.Println("recovered %v", xxx)
	}}


}}

func main() {
{
	defer trap()
	fmt.Println( div(0,0) )

	fmt.Println( div(3,2) )

}}