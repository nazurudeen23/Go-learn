package main

import "fmt"

type RockClimber struct {
    name string
}

func (rc *RockClimber) climb() {
    fmt.Println(rc.name, "is climbing!")
}

func main() {
    rc := &RockClimber{name: "Alice"}
    rc.climb() // Output: Alice is climbing!
}
