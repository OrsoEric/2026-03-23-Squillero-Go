package main

import "fmt"

func collatz(n int) {
{
	defer fmt.Println(n)
    if n == 1 {
	{
        return
    }}
    if n%2 == 0 {
	{
        collatz(n/2)
    }} else {
	{
        collatz(3*n+1)
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

    collatz(n)
}}
