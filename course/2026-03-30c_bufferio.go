//go run course/2026-03-30c_bufferio.go

package main

//import do not have a post processor to brick the curly bracket, designers were not strong enough to implement it in the tokenizer
import
(
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
{
	// Open the file
	file, err := os.Open("course/shrek_script.txt")
	if (err != nil) {
	{
		log.Fatal(err)
	}}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for (scanner.Scan()) {
	{
		// Print each line
		fmt.Println(scanner.Text())
	}}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
	{
		log.Fatal(err)
	}}
}}