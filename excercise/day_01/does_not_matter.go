
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

//-----------------------------------------------------------------------------
// SEQUENCE DECODER
//-----------------------------------------------------------------------------

type St_instruction struct
{
    x_right bool
    n_step  int
}

type Fn_sequence_decoder struct
{
    i_as_line string
}

func Fn_decode_sequence(i_as_line []string) ([]St_instruction, error) {
{
    var out []St_instruction

    for _, s_line := range i_as_line {
    {
        if (len(s_line) < 2) {
        {
            return nil, fmt.Errorf("invalid token: %s", s_line)
        }}

        s_dir := s_line[0]
		if (s_dir != 'L' && s_dir != 'R') {
        {
            return nil, fmt.Errorf("invalid direction in token: %s", s_line)
        }}

        s_num := s_line[1:]

        n_num, e_error := strconv.Atoi(s_num)
        if (e_error != nil) {
        {
            return nil, fmt.Errorf("invalid number in token %s", s_line)
        }}

        step := St_instruction{
            x_right: (s_dir == 'R'),
            n_step:  n_num,
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
		I_s_path: "day_01/puzzle_cue.txt",
    	//I_s_path: "day_01/puzzle_input.txt",
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
		fmt.Printf("Sequence Length: %d\n", len(decoded))
		for _, d := range decoded {
		{
			fmt.Printf("Right=%v  Step=%d\n", d.x_right, d.n_step)
		}}
	}}

	//STEP3:


}}
