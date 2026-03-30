// Package day_02 implements several utilities for reading,
// tokenizing, and parsing puzzle input files containing numeric
// ranges in the form "A-B,C-D,E-F,...".
//
// This file intentionally uses a forced ANSI‑bracket style,
// where scope blocks are wrapped using:
//
//     {
//     {
//         ... code ...
//     }}
//
// This is a stylistic choice to emulate ANSI‑style bracketing
// and to bypass Go’s tokenizer restrictions on bracket placement.
// No logic is altered by this formatting.

package day_02

//go get strconv
import
(
    
    "fmt"
	"log"
    "os"
	"bufio"
    "strings"
	"strconv"
	"bytes"
	"time"
)

// -----------------------------------------------------------------------------
// FILE READER
// -----------------------------------------------------------------------------
//
// Fn_sequence_reader reads a text file line‑by‑line and returns a slice
// of trimmed, non‑empty strings. It is a classical scanner‑based reader.
//
// PARAMETERS:
//   i_s_file_path — path to the file to read.
//
// RETURNS:
//   []string — all non‑empty lines
//   error    — any file or scanner error encountered
//
// NOTES:
//   • Empty lines are skipped.
//   • Leading/trailing whitespace is removed.
//   • The function uses ANSI‑style double‑bracket scoping.
//

func Fn_sequence_reader(i_s_file_path string)(
	[]string,
	error) {
{
    i_st_file, e_error := os.Open(i_s_file_path)
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

// -----------------------------------------------------------------------------
// STRUCTURE PAIR 
// -----------------------------------------------------------------------------
//
// St_pair represents a numeric range "A-B". The two integers are stored
// in a fixed‑size array for compactness and predictable memory layout.
//
// Example:
//     St_pair{ n_value: [2]int{11, 22} }
//
// The String() method prints the pair in "A-B" format.
//

type St_pair struct
{
    n_value [2]int
}

func (i_st_pair St_pair) String() string {
{
    return fmt.Sprintf("%d-%d", i_st_pair.n_value[0], i_st_pair.n_value[1])
}}

// -----------------------------------------------------------------------------
// Fn_find_pair_in_content
// -----------------------------------------------------------------------------
//
// Fn_find_pair_in_content parses a slice of strings where each string
// contains comma‑separated numeric ranges, e.g.:
//
//     "11-22,95-115,998-1012"
//
// The function extracts all pairs across all lines.
//
// PARAMETERS:
//   i_as_content — slice of input lines
//
// RETURNS:
//   []St_pair — all parsed numeric pairs
//   error     — malformed input or conversion error
//
// BEHAVIOR:
//   • Splits each line by commas.
//   • Splits each chunk by '-'.
//   • Converts both sides to integers.
//   • Appends a St_pair for each valid range.
//

func Fn_find_pair_in_content(i_as_content []string)(
	[]St_pair,
	error) {
{
	//allocate a array of pairs
	var ast_pair []St_pair

    // Process each line
    for _, s_line := range i_as_content {
    {
        // Split by comma
        as_chunks := strings.Split(s_line, ",")

        for _, s_chunk := range as_chunks {
        {
            s_chunk = strings.TrimSpace(s_chunk)
            if s_chunk == "" {
                continue
            }

            // Split by dash
            as_range := strings.Split(s_chunk, "-")
            if len(as_range) != 2 {
                return nil, fmt.Errorf("invalid pair format: %s", s_chunk)
            }


			n_left, e_error_left := strconv.Atoi(as_range[0])
			n_right, e_error_right := strconv.Atoi(as_range[1])
			if e_error_left != nil || e_error_right != nil {
			{
				return nil, fmt.Errorf("invalid number in token: %s", s_chunk)
			}}

            // Append pair
            ast_pair = append(ast_pair, St_pair{
                n_value: [2]int{n_left, n_right},
            })
        }}
    }}

    return ast_pair, nil
}}

//-----------------------------------------------------------------------------
// TOKEN READER
//-----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// TOKEN READER (Scanner Split Function)
// -----------------------------------------------------------------------------
//
// Fn_split_scanner reads a file and uses a custom bufio.SplitFunc to
// tokenize the input based on commas. Each token is expected to be a
// numeric range "A-B".
//
// This approach allows streaming tokenization without loading entire
// lines, and without requiring commas to align with newline boundaries.
//
// PARAMETERS:
//   i_s_file_path — path to the puzzle input file
//
// RETURNS:
//   []St_pair — parsed numeric ranges
//   error     — file, scanner, or parsing error
//
// CUSTOM SPLIT FUNCTION:
//   • Searches for ',' in the byte buffer.
//   • Returns the token before the comma.
//   • At EOF, returns remaining data as final token.
//   • Requests more data when needed.
//
// This is useful for extremely large files or continuous streams.
//

func Fn_split_scanner(i_s_file_path string) ([]St_pair, error) {
{
    st_file, e_error := os.Open(i_s_file_path)
    if e_error != nil {
	{
        return nil, e_error
    }}
    defer st_file.Close()

    cl_scanner := bufio.NewScanner(st_file)

	// Custom split function: split on commas
	var fn_split_comma bufio.SplitFunc = func(
		i_s_data []byte,
		i_x_end_of_file bool) (
		o_n_index_advance int,
		o_s_tokens []byte,
		o_e_error error) {
	{
			// Look for a comma
			if n_index := bytes.IndexByte(i_s_data, ','); n_index >= 0 {
			{
				// We found a full token ending at comma
				return n_index + 1, bytes.TrimSpace(i_s_data[:n_index]), nil
			}}

			// If we're at EOF, return the remaining data
			if i_x_end_of_file && len(i_s_data) > 0 {
			{
				return len(i_s_data), bytes.TrimSpace(i_s_data), nil
			}}

			// Request more data
			return 0, nil, nil
    }}

    // Install custom tokenizer
    cl_scanner.Split( fn_split_comma )

    var ast_pair []St_pair

    for cl_scanner.Scan() {
	{
        i_s_tokens := cl_scanner.Text()
        if len(i_s_tokens) <= 0 {
		{
            continue
        }}

        i_as_part := strings.Split(i_s_tokens, "-")
        if len(i_as_part) != 2 {
		{
            return nil, fmt.Errorf("invalid token: %s", i_s_tokens)
        }}

		n_left, e_error_left := strconv.Atoi(i_as_part[0])
        n_right, e_error_right := strconv.Atoi(i_as_part[1])
        if e_error_left != nil || e_error_right != nil {
		{
            return nil, fmt.Errorf("invalid number in token: %s", i_s_tokens)
        }}

		ast_pair = append(ast_pair, St_pair{n_value: [2]int{n_left, n_right}} )
    }}

    if e_error := cl_scanner.Err(); e_error != nil {
	{
        return nil, e_error
    }}

    return ast_pair, nil
}}

// -----------------------------------------------------------------------------
// SCAN PAIRS
// -----------------------------------------------------------------------------
// This function scan pairs


func Fn_scan_pair(
	i_ast_pair []St_pair) (
	o_n_num_invalid int,
	o_n_sum_invalid int,
	o_e_error error) {
{
	t_start := time.Now()

	var n_num_invalid int = 0
	var n_sum_invalid int = 0

	//For each pair
	for _, st_pair := range(i_ast_pair) {
	{
		//for each value in the range
		for n_cnt := st_pair.n_value[0]; n_cnt <= st_pair.n_value[1]; n_cnt++ {
		{
			x_invalid, e_error := Fn_is_half_number_same( n_cnt ) 
			if e_error != nil {
			{
				return 0, 0, fmt.Errorf("invalid invalid chec on number: %n", n_cnt)
			}}
			if x_invalid == true {
			{
				n_num_invalid += 1
				n_sum_invalid += n_cnt
			}}
		}}
	}}

	t_elapsed := time.Since(t_start)
	log.Printf("Elapsed: %d\n", t_elapsed.Microseconds())

	return n_num_invalid, n_sum_invalid, nil
}}

// -----------------------------------------------------------------------------
// PARALLEL CHECK MIRROR
// -----------------------------------------------------------------------------
// Fn_scan_pair_parallel launches one goroutine per St_pair.
// Each goroutine scans its numeric range and reports:
//
//   • number of invalid values
//   • sum of invalid values
//   • error (if any)
//
// The parent function aggregates all results.
//
// This allows ALL ranges to be processed concurrently.
//
// NOTE:
//   The ANSI‑bracket style is preserved exactly as requested.
//
func Fn_scan_pair_parallel(
    i_ast_pair []St_pair) (
    o_n_num_invalid int,
    o_n_sum_invalid int,
    o_e_error error) {
{
	t_start := time.Now()

    // Result structure for channel communication
    type St_result struct {
        n_num_invalid int
        n_sum_invalid int
        e_error       error
    }

    // Create channel with buffer = number of pairs
    ch_result := make(chan St_result, len(i_ast_pair))

    // Launch one goroutine per pair
    for _, st_pair := range(i_ast_pair) {
    {
        go func(i_st_pair St_pair) {
        {
            var n_local_invalid int = 0
            var n_local_sum int = 0

            // Scan the numeric range
            for n_cnt := i_st_pair.n_value[0]; n_cnt <= i_st_pair.n_value[1]; n_cnt++ {
            {
                x_invalid, e_error := Fn_is_half_number_same(n_cnt)
                if e_error != nil {
                {
                    ch_result <- St_result{
                        n_num_invalid: 0,
                        n_sum_invalid: 0,
                        e_error:       fmt.Errorf("invalid check on number: %d", n_cnt),
                    }
                    return
                }}

                if x_invalid == true {
                {
                    n_local_invalid += 1
                    n_local_sum += n_cnt
                }}
            }}

            // Send result back
            ch_result <- St_result{
                n_num_invalid: n_local_invalid,
                n_sum_invalid: n_local_sum,
                e_error:       nil,
            }
        }}(st_pair)
    }}

    // Collect results
    var n_total_invalid int = 0
    var n_total_sum int = 0

    for n_cnt := 0; n_cnt < len(i_ast_pair); n_cnt++ {
    {
        st_res := <-ch_result

        if st_res.e_error != nil {
        {
            return 0, 0, st_res.e_error
        }}

        n_total_invalid += st_res.n_num_invalid
        n_total_sum += st_res.n_sum_invalid
    }}

	t_elapsed := time.Since(t_start)
	log.Printf("Elapsed: %d\n", t_elapsed.Microseconds())

    return n_total_invalid, n_total_sum, nil
}}


// -----------------------------------------------------------------------------
// CHECK MIRROR
// -----------------------------------------------------------------------------
// Fn_is_half_number_same determines whether an integer can be split into two
// equal-length halves AND whether those halves contain identical digit sequences.
//
// RULES:
//   • Odd number of digits  → invalid (o_x_invalid = true)
//   • Even digits, halves differ → valid format but NOT same (o_x_invalid = true)
//   • Even digits, halves equal → valid format AND same (o_x_invalid = false)
//
// RETURNS:
//   o_x_invalid = true  → number cannot be split evenly (odd digit count)
//   o_x_invalid = false → number is splittable; halves may or may not match
//   o_e_error   = nil   → always nil unless future logic adds errors
//
// EXAMPLES:
//   12345   → odd digits → invalid
//   123123  → even digits, halves equal → valid
//   123321  → even digits, halves differ → valid
//
func Fn_is_half_number_same(
	i_n_value int) (
	o_x_invalid bool,
	o_e_error error) {
{
    // Convert number to string for digit inspection
    s_value := strconv.Itoa(i_n_value)

    // Check if digit count is even
    if (len(s_value) % 2 != 0) {
    {
        // Cannot split evenly
        return false, nil
    }}

    // Compute midpoint
    n_mid := len(s_value) / 2

    // Extract halves
    s_left  := s_value[:n_mid]
    s_right := s_value[n_mid:]

    // Compare halves
    if (s_left == s_right) {
    {
		//log.Printf("SAME %s %s", s_left, s_right)
        // Halves match
        return true, nil
    }} else {
	{
		//log.Printf("NOT SAME %s %s", s_left, s_right)
	}}


    // Halves do not match
    return false, nil
}}

// -----------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------
//
// Part1 is the entry point for this puzzle stage. It demonstrates:
//
//   • Reading puzzle input
//   • Extracting numeric ranges
//   • Logging results
//
// The function currently uses Fn_split_scanner, but includes commented
// code showing how to use Fn_sequence_reader + Fn_find_pair_in_content.
//
// The ANSI‑bracket style is preserved throughout.
//

func Part1() {
{
	fmt.Printf("Shaka, when the walls fell.\n")

	log.Printf("START LOG")

	//-----------------------------------------------------------------------------
	// CLASSICAL READ FILE AND PRICESS
	//-----------------------------------------------------------------------------

	//var s_puzzle_path string = "day_02/puzzle_cue.txt"
	var s_puzzle_path string = "day_02/puzzle_input.txt"

	/*
	as_content, e_error := Fn_sequence_reader( s_puzzle_path )
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	if (false) {
	{
		fmt.Printf("=========Read from file=======\n")
		fmt.Printf("Lines: %d\n", len(as_content))
		for _, s_line := range as_content {
		{
			log.Printf(s_line)
		}}
	}}



	ast_pair, e_error := Fn_find_pair_in_content( as_content )
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	if (true) {
	{
		fmt.Printf("=========Extract Pairs=======\n")
		fmt.Printf("Pairs: %d\n", len(ast_pair))
		for _, st_pair := range ast_pair {
		{
			//log.Printf("%d-%d", st_pair.n_value[0], st_pair.n_value[1] )
			log.Printf("%s", st_pair )
		}}
	}}
	*/

	//-----------------------------------------------------------------------------
	// SCANNER SPLITTER THAT USES SPLIT LAMBDA AND IS INCOMPREHENSIBLE
	//-----------------------------------------------------------------------------

	ast_pair, e_error := Fn_split_scanner( s_puzzle_path )
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	if (true) {
	{
		fmt.Printf("=========Extract Pairs=======\n")
		fmt.Printf("Pairs: %d\n", len(ast_pair))
		for _, st_pair := range ast_pair {
		{
			//log.Printf("%d-%d", st_pair.n_value[0], st_pair.n_value[1] )
			log.Printf("%s", st_pair )
		}}
	}}
	

	//-----------------------------------------------------------------------------
	// CHECK PAIRS SERIAL
	//-----------------------------------------------------------------------------

	fmt.Printf("=========SERIAL=======\n")
	fmt.Printf("Pairs: %d\n", len(ast_pair))
	n_num_invalid, n_sum_invalid, e_error := Fn_scan_pair( ast_pair )
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	
	log.Printf("Num Invalid: %d", n_num_invalid )
	log.Printf("Sum Invalid: %d", n_sum_invalid )

	//-----------------------------------------------------------------------------
	// CHECK PAIRS PARALLEL
	//-----------------------------------------------------------------------------

	fmt.Printf("=========PARALLEL=======\n")
	fmt.Printf("Pairs: %d\n", len(ast_pair))
	n_num_invalid, n_sum_invalid, e_error = Fn_scan_pair_parallel( ast_pair )
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	log.Printf("Num Invalid: %d", n_num_invalid )
	log.Printf("Sum Invalid: %d", n_sum_invalid )

	log.Printf("STOP LOG")
}}