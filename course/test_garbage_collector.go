//go run  test_garbage_collector.go
package main

import
(
	"fmt"
	"runtime"
	"time"
)

type MyObject struct
{
	name string
}

func main() {
{
	runGarbageCollectorDemo()
}}

func runGarbageCollectorDemo() {
{
	obj := &MyObject{name: "example"}

	runtime.SetFinalizer(
		obj,
		func(o *MyObject) {
		{
			fmt.Println("Object garbage collected:", o.name)
	}})

	fmt.Println("Object created")

	time.Sleep(2 * time.Second)

	obj = nil // remove reference

	fmt.Println("Forcing garbage collection...")
	runtime.GC()

	time.Sleep(2 * time.Second)

	fmt.Println("Done")
}}