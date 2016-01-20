package main

import (
        "fmt"
        "test/stringutil"
)

func main() {
        fmt.Printf("hello, go world\n")
        fmt.Printf( stringutil.Reverse("hello, go world") )
        fmt.Printf("\n");
}
