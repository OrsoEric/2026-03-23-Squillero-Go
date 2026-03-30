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
	//"strconv"
	"bytes"
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
// SPLITTER
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
    s_value [2]string
}

func (i_st_pair St_pair) String() string {
{
    return fmt.Sprintf("%s-%s", i_st_pair.s_value[0], i_st_pair.s_value[1])
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

            // Append pair
            ast_pair = append(ast_pair, St_pair{
                s_value: [2]string{as_range[0], as_range[1]},
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
		n_index_advance int,
		i_s_tokens []byte,
		e_error error) {
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

		ast_pair = append(ast_pair, St_pair{s_value: [2]string{i_as_part[0], i_as_part[1]}} )
    }}

    if e_error := cl_scanner.Err(); e_error != nil {
	{
        return nil, e_error
    }}

    return ast_pair, nil
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

	var s_puzzle_path string = "day_02/puzzle_cue.txt"

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

	if (false) {
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
	// 
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

	log.Printf("STOP LOG")
}}