package main

import "fmt"

type RockClimber struct {
    name string
}

func climb(rc *RockClimber) {
    fmt.Println(rc.name, "is climbing!")
}

func main() {
    rc := &RockClimber{name: "Alice"}
    climb(rc) // Output: Alice is climbing!
}
