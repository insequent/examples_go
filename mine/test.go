package main

import (
    "fmt"
    "reflect"
)

func main() {
    blah := make(map[string]interface{})
    blah["stuff"] = 2
    blah["thing"] = 3

    var v interface{}
    var ok bool
    v = blah["not"]
    fmt.Println("v:", v, "Type:", reflect.TypeOf(v))
    fmt.Println("ok:", ok)

    if v != nil {
        for x, y := range v.(map[string]interface{}) {
            fmt.Println(x, y)
        }
    }
}
