package main

import (
    "fmt"
    "os"
)

const World = "world"

func main() {
    var myname = os.Getenv("HELLO_GO_NAME")
    if myname == "" {
        myname = World
    }
    myname += "!"
    fmt.Printf("Hello, %x.\n",myname)
}