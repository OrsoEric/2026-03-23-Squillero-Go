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

