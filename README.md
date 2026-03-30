install go 
https://go.dev/doc/install

REBOOT

open cmd

where go

set PATH=%PATH%;D:\Programs\go\bin

go version

D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go>go version
go version go1.26.1 windows/amd64

```go
package main

import "fmt"

func main() { 
{
    fmt.Println("Shaka, when the walls fell") 
}}

```

go run test.go

go build test.go


```cmd
D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go>test.exe
Shaka, when the walls fell
```

# ?????

mkdir course 

cd course

go mod init course

D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go\course>go run main.go

What do you get if you multiply six by nine?

# MOD

```cmd
go mod verify

D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go\course>go mod verify
all modules verified
```

# external MOD

go get github will add a dependency

go.mod list the version

go.sum has the hashes

this command clean up go mod to put dependencies in direct and indirect
go mod tidy

```cmd
module course

go 1.26.1

require github.com/briandowns/spinner v1.23.2

require (
	github.com/fatih/color v1.7.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/term v0.1.0 // indirect
)
```

```cmd
D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go\course>go vet
# course
# [course]
vet.exe: .\main.go:12:6: main redeclared in this block
```

# Visibility

local name lowercse

public names Uppercase

# any

data type any will bypass the static type check

# pointer

used to give large structures between function, and modify back

unsafe keyword to do pointer math nonsense

```go
a := 41
b := &a
*b++

a = 42
```

# garbage collector

# _ blank identifier

it's a dummy variable write only that is discarded

to fool the compiler 

# defer

do something before the last curly bracket

in recursive code will unfurl in reverse direction

# LLM

GO is incompetently designed and doesn't allow ANSI bracket style. You are to use forced ansi brackets using scope brackets properly as per ansi style. Example below

```go
//import do not have a post processor to brick the curly bracket, designers were not strong enough to implement it in the tokenizer
import
(
	"log"
	"github.com/pterm/pterm"
)
//functions and if have a post processor to brick ansi curly brackets, but proper indentation is restored with scope brackets
//Note how its bracket, new line bracket, indented code, double close bracket
func my_function() {
{
	if (true) {
	{
	}}
}}
```

# pterm dependency

<details>
<summary>XXX</summary>

```
D:\Data\Project\Project Programming\Project Go\2026-03-23 Squillero Go>go run course/2026-03-30a.go
 INFO  Shaka
Please select your options [type to search]:
> [✗] Poo
Please select your options [type to search]:                    
> [✗] Poo                       
  [✗] Maa                       
  [✗] Taa                       
  [✗] Dii                       
enter: select | tab: confirm | left: none | right: all | type to
 filter
```

</details>

# Go generics

compile time parametric polymorphism, it will generate a different function for each type

c++ will go above and beyont to monomorphize, zero cost executable

go will pack object of same size, some cost in execution time

```go
//go mod init 2026-03-30b_constraint.go
//go get "golang.org/x/exp/constraints"
//go run course/2026-03-30b_constraint.go

package main

import 
(
	"fmt"
	"golang.org/x/exp/constraints"
)

func my_function[constraints.Float](value T) {
{
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}}

func my_function[T constraints.Integer](value T) {
{
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}}

func main() {
{
   	my_function(42)
	my_function(3.14)
}}
```

```cmd
_constraint.go
Value: 42, Type: int
Value: 3.14, Type: float64
```

# CONCURRENCY GORUNTIME

Concurrency doing things at the same time
	e.g. mouse and mp3 on single thread machine

Parallelism executing many timgs at the same time

Go claim a million gorutines


