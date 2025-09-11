package main

import (
    "fmt"
)

func main() {
    stringList := []string{"spaceship", "sugar", "football", "sun", "paper"}
    fmt.Println(filterWords(stringList, "su"))
}

func filterWords(stringList []string, prefix string) []string {
    newStringList := []string{}

    // Outer loop: go through each word in the list
    for _, word := range stringList {

        // 2. [KEY TO PREVENT PANIC]
        // Make sure prefix is not longer than the word
        // (otherwise word[i] would panic because we go out of range)
        if len(prefix) > len(word) {
            continue
        }

        match := true

        // 1. Nested for loop: go through each character of the prefix
        for i := 0; i < len(prefix); i++ {

            // 1.2. Compare each character of prefix with the word
            if word[i] != prefix[i] {
                match = false
                break // stop checking this word
            }
        }

        // If all prefix characters matched, add word to results
        if match {
            newStringList = append(newStringList, word)
        }
    }

    return newStringList
}


