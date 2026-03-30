
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

func (
	i_cl_file *Fn_sequence_reader) Read() (
	[]string,
	error) {
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

func Fn_decode_sequence(
	i_as_line []string) (
	[]St_instruction,
	error) {
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

//-----------------------------------------------------------------------------
// ZERO COUNTER
//-----------------------------------------------------------------------------

func Fn_zero_counter_from_instruction(
	i_n_start int,
	i_ast_instruction []St_instruction,
	i_x_measure_zero_during_passage bool) (
	int,
	error) {
{
	var n_cnt_zero int = 0
	var n_dial int = i_n_start

	for _, st_instruction := range i_ast_instruction {
	{
		var n_step int = st_instruction.n_step

		for {
		{
			if n_step >= 100 {
			{
				//I do a full rotation that does not change the dial position
				n_step -= 100

				//a full rotation HAS to transit through zero
				if i_x_measure_zero_during_passage == true {	
				{
					n_cnt_zero += 1
				}}
			}} else if n_step > 0 {
			{
				//If I have steps left
				n_step -= 1

				if st_instruction.x_right == true {
				{
					n_dial += 1
					if n_dial > 99 {
					{
						n_dial = 0
					}}
				}} else {
				{
					n_dial -= 1
					if n_dial < 0 {
					{
						n_dial = 99
					}}

				}}

				//measure during rotation
				if n_dial == 0 && i_x_measure_zero_during_passage == true {	
				{
					n_cnt_zero += 1
				}}
			}} else {
			{
				break
			}}
		}}

		//measure zero at the end of the dial 
		if n_dial == 0 && i_x_measure_zero_during_passage == false {	
		{
			n_cnt_zero += 1
		}}

		//fmt.Printf("Direction: %v | Step: %d | Dial: %d\n", st_instruction.x_right , st_instruction.n_step, n_dial)
	}}

	return n_cnt_zero, nil
}}



//-----------------------------------------------------------------------------
// PART 1
//-----------------------------------------------------------------------------

func Part12(
	i_x_measure_zero_during_passage bool ) {
{
	fmt.Println("Shaka, when the walls fell")

	//STEP1: read the instructions

	cl_sequence_reader := Fn_sequence_reader{
		//I_s_path: "day_01/puzzle_cue.txt",
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

	ast_instruction, e_error := Fn_decode_sequence(as_sequence)
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	if (false) {
	{
		fmt.Printf("Sequence Length: %d\n", len(ast_instruction))
		for _, st_instruction := range ast_instruction {
		{
			fmt.Printf("Right: %v | Step: %d |\n", st_instruction.x_right, st_instruction.n_step)
		}}
	}}

	//STEP3: count the zeros following the instructuibns
	n_cnt_zero, e_error := Fn_zero_counter_from_instruction(50, ast_instruction, i_x_measure_zero_during_passage)


	fmt.Printf("Dials Zero : %d\n", n_cnt_zero )

}}
