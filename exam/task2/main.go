package main

import (
    "fmt"
    "sort"
)

type BrainrotMeme struct {
    Name       string
    TrendLevel int
    Category   string
    Views      float64
}

func createMemes() []BrainrotMeme {
    return []BrainrotMeme{
        {"Ohio Gyatt", 9, "Subo Bratik", 120.5},
        {"Fanum Tax", 8, "Sigma", 85.2},
        {"Skibidi Toilet", 10, "Skibidi", 350.7},
        {"Mewing Challenge", 7, "Mewing", 45.3},
        {"Rizzler", 6, "Sigma", 65.8},
        {"Tuntuntunsahur", 9, "TUNTUNTUNSAHUR", 200.1},
        {"What the Sigma", 5, "Sigma", 30.4},
        {"Ohio Rizz", 8, "Other", 75.6},
    }
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
    var result []BrainrotMeme
    
    for _, meme := range memes {
        if meme.Views > minViews {
            result = append(result, meme)
        }
    }
    
    sort.Slice(result, func(i, j int) bool {
        return result[i].TrendLevel > result[j].TrendLevel
    })
    
    return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
    impact := make(map[string]float64)
    
    for _, meme := range memes {
        impact[meme.Category] += meme.Views
    }
    
    return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
    var result []string
    
    for _, meme := range memes {
        if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
            result = append(result, meme.Name)
        }
    }
    
    return result
}

func main() {
    memes := createMemes()
    
    fmt.Println("=== Все мемы ===")
    for i, meme := range memes {
        fmt.Printf("%d. %s | Тренд: %d | Категория: %s | Просмотры: %.1fM\n",
            i+1, meme.Name, meme.TrendLevel, meme.Category, meme.Views)
    }
    
    fmt.Println("\n=== Топовые тренды (просмотры > 70M) ===")
    topTrending := FindTopTrending(memes, 70)
    for i, meme := range topTrending {
        fmt.Printf("%d. %s | Тренд: %d | Просмотры: %.1fM\n",
            i+1, meme.Name, meme.TrendLevel, meme.Views)
    }
    
    fmt.Println("\n=== Влияние категорий ===")
    categoryImpact := CalculateCategoryImpact(memes)
    for category, views := range categoryImpact {
        fmt.Printf("Категория %s: %.1fM просмотров\n", category, views)
    }
    
    fmt.Println("\n=== Мемы по сложному условию ===")
    filtered := FilterByComplexCondition(memes)
    fmt.Println("Названия мемов:")
    for _, name := range filtered {
        fmt.Printf("- %s\n", name)
    }
}