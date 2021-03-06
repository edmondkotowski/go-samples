package main

import (
    "fmt"
)

func main() {
    perm := isPermutation("abcabc", "aabbcc")

    fmt.Printf("%v\n", perm)
}

func isPermutation(value1 string, value2 string) (bool) {
    if len(value1) != len(value2) {
        return false
    }

    lookup := make(map[string]int)

    for i := range value1 {
        temp := string(value1[i])
        if lookup[temp] == 0 {
            lookup[temp] = 1
            continue
        }

        lookup[temp]++
    }

    for x := range value2 {
        temp := string(value2[x])
        if lookup[temp] == 0 {
            return false
        }

        lookup[temp]--

        if lookup[temp] < 0 {
            return false
        }
    }

    return true
}