package main

import "fmt"

func collatz(n int, sequence *[]int) {
{
	defer fmt.Println(n)
    *sequence = append(*sequence, n)
    if n == 1 {
	{
        return
    }}
    if n%2 == 0 {
	{
        collatz(n/2, sequence)
    }} else {
	{
        collatz(3*n+1, sequence)
    }}
}}

func main() {
{
    var n int
    fmt.Print("Enter a positive integer: ")
    _, err := fmt.Scan(&n)
    if err != nil || n <= 0 {
	{
        fmt.Println("Invalid input. Please enter a positive integer.")
        return
    }}

    var sequence []int
    collatz(n, &sequence)
}}
