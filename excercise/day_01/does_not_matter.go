
/*



*/

package day_01

import
(
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Fn_sequence_reader struct
{
    Path string
}

func (i_cl_file *Fn_sequence_reader) Read() ([]string, error) {
{
    i_st_file, e_error := os.Open(i_cl_file.Path)
    if (e_error != nil) {
    {
        return nil, e_error
    }}
    defer i_st_file.Close()

    var s_content []string

    cl_scanner := bufio.NewScanner(i_st_file)
    for cl_scanner.Scan() {
    {
        s_line := strings.TrimSpace(cl_scanner.Text())
        if (s_line != "") {
        {
            s_content = append(s_content, s_line)
        }}
    }}

    if (cl_scanner.Err() != nil) {
    {
        return nil, cl_scanner.Err()
    }}

    return s_content, nil
}}

// function definition
func Shaka() {
{
	fmt.Println("Shaka, when the walls fell")

	cl_sequence_reader := Fn_sequence_reader{
         Path: "day_01/puzzle_input.txt",
	}

    as_sequence, e_error := cl_sequence_reader.Read()
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

    for _, s := range as_sequence {
    {
        fmt.Println(s)
    }}
}}