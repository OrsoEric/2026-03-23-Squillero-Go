
/*



*/

package day_01

//go get strconv
import
(
    "bufio"
    "fmt"
    "os"
    "strings"
	"strconv"
)



//-----------------------------------------------------------------------------
// FILE READER
//-----------------------------------------------------------------------------

type Fn_sequence_reader struct
{
    I_s_path string
}

func (i_cl_file *Fn_sequence_reader) Read() ([]string, error) {
{
    i_st_file, e_error := os.Open(i_cl_file.I_s_path)
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

type St_instruction struct
{
    x_right bool
    n_step  int
}

func Fn_decode_sequence(seq []string) ([]St_instruction, error) {
{
    var out []St_instruction

    for _, s := range seq {
    {
        if (len(s) < 2) {
        {
            return nil, fmt.Errorf("invalid token: %s", s)
        }}

        dir := s[0]
        numStr := s[1:]

        amt, err := strconv.Atoi(numStr)
        if (err != nil) {
        {
            return nil, fmt.Errorf("invalid number in token %s", s)
        }}

        step := St_instruction{
            x_right: (dir == 'R'),
            n_step:  amt,
        }

        out = append(out, step)
    }}

    return out, nil
}}


// function definition
func Shaka() {
{
	fmt.Println("Shaka, when the walls fell")

	//STEP1: read the instructions

	cl_sequence_reader := Fn_sequence_reader{
    	I_s_path: "day_01/puzzle_input.txt",
	}

    as_sequence, e_error := cl_sequence_reader.Read()
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	if (false) {
	{
		fmt.Println(len(as_sequence))
		for _, s := range as_sequence {
		{
			fmt.Println(s)
		}}
	}}

	//STEP2: decode the sequence into a structure

	decoded, err := Fn_decode_sequence(as_sequence)
    if (err != nil) {
    {
        fmt.Println("Error:", err)
        return
    }}

	if (true) {
	{
		fmt.Println("Sequence Length: %d", len(decoded))
		for _, d := range decoded {
		{
			fmt.Printf("Right=%v  Step=%d\n", d.x_right, d.n_step)
		}}
	}}
}}
