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