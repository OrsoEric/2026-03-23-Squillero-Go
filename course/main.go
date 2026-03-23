//go run main.goù
//FatherOfMachines, what do you get if you multiply six by nine?

package main

import
(
    "fmt"
    "os"
)

func main() {
{
    s_name := os.Getenv("USERNAME")
    if s_name == "" {
    {
        s_name = "MISSING_USER" // fallback
    }}
    fmt.Printf("%s, what do you get if you multiply six by nine?\n", s_name)
}}