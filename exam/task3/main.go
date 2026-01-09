package main

import (
    "fmt"
    "strings"
)

func encryptWord(word string) string {
    if len(word) <= 1 {
        return word
    }
    
    firstChar := string(word[0])
    
    rest := word[1:]
    reversedRest := reverseString(rest)
    
    return firstChar + reversedRest
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func encryptPhrase(phrase string) string {
    words := strings.Fields(phrase)
    
    encryptedWords := make([]string, len(words))
    for i, word := range words {
        encryptedWords[i] = encryptWord(word)
    }
    
    return strings.Join(encryptedWords, " ")
}

func main() {
    phrases := []string{
        "Pepe Schnele is a legend",
        "Hello World",
        "Go programming is fun",
        "Brainrot meme culture",
    }
    
    fmt.Println("=== Шифратор фраз ===")
    for _, phrase := range phrases {
        encrypted := encryptPhrase(phrase)
        fmt.Printf("Исходная:  %s\n", phrase)
        fmt.Printf("Шифровка:  %s\n\n", encrypted)
    }
}