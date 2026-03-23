//go run demo_struct.go

package main

import "fmt"

//field is a metadata that is ancillary
type Shaka struct
{
	//visiblle from outside
	S_name string `json:"proper_name"`
	//not visible from outside
	s_password string `json:"hidden_password"`
}

func (i_st Shaka) Greet(){
{
	fmt.Printf("Hello: %s ", i_st.S_name)
}}


func main() {
{
   	st_king := Shaka{S_name:"Kemehameah"}
	st_king.Greet()
}}