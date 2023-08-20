package main

import (
    "fmt"
    "gitea.home/go-laravel/celeritas"
)

func main() {
    result := celeritas.TestFunc(1, 1)
    fmt.Println(result)

    result = celeritas.TestFunc2(2, 1)
    fmt.Println(result)
}
