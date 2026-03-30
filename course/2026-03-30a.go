//go mod init 2026-03-30a
//go get github.com/pterm/pterm
//go run course/2026-03-30a.go
// this will brick the go.mod because of the many go programs using main
//go mod tidy

package main

import
(
	"log"
	"github.com/pterm/pterm"
)

func main() {
{
	var e_error error
	pterm.Info.Println("Shaka")

	ls_teletubbies := []string{"Poo", "Maa", "Taa", "Dii"}
	s_formatted_options, e_error := pterm.DefaultInteractiveMultiselect.WithOptions(ls_teletubbies).Show()
	if e_error != nil {
	{
		log.Panicf("Damnation: %v\n", e_error.Error() )
	}}

	pterm.Info.Println("Options: %s", pterm.Green(s_formatted_options) )
}}
