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
)

//-----------------------------------------------------------------------------
// FILE READER
//-----------------------------------------------------------------------------

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

//-----------------------------------------------------------------------------
// SPLITTER
//-----------------------------------------------------------------------------
//I want to parse the slice of strings to fetch all the pairs
// 11-22,95-115,998-1012,1188511880-1188511890,222220-222224,

type St_pair struct
{
    n_value [2]int
}

func (i_st_pair St_pair) String() string {
{
    return fmt.Sprintf("%d-%d", i_st_pair.n_value[0], i_st_pair.n_value[1])
}}


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

            // Convert to ints
            n_a, errA := strconv.Atoi(as_range[0])
            n_b, errB := strconv.Atoi(as_range[1])
            if errA != nil || errB != nil {
                return nil, fmt.Errorf("invalid number in pair: %s", s_chunk)
            }

            // Append pair
            ast_pair = append(ast_pair, St_pair{
                n_value: [2]int{n_a, n_b},
            })
        }}
    }}

    return ast_pair, nil
}}


//-----------------------------------------------------------------------------
// MAIN
//-----------------------------------------------------------------------------

func Part1() {
{
	fmt.Printf("Shaka, when the walls fell.\n")

	log.Printf("START LOG")

	as_content, e_error := Fn_sequence_reader( "day_02/puzzle_cue.txt")
    if (e_error != nil) {
    {
        fmt.Println("Error:", e_error)
        return
    }}

	if (true) {
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

	log.Printf("STOP LOG")
}}