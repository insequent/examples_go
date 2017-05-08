package main

import(
    "fmt"
    "reflect"
)

func main() {
    fmt.Println("Let's loop!")

    var i interface{}
    i = nil

    if m, ok := i.(map[string]interface{}); ok {
        for k, v := range m {
            fmt.Println("Shocking! We're inside a nil loop!", k, v)
        }
    } else {
        fmt.Println("Eek, we're nil!")
    }

    fmt.Println("The type of an nil map is", reflect.TypeOf(i))
}
