package main

import (
        "context"
        "fmt"
        "log"
        "os"
        "cloud.google.com/go/translate"
        "golang.org/x/text/language"
)

func main() {
        if len(os.Args) < 2 {
            fmt.Println("Usage: ./app \"Text to translate.\" [--show-arabic]")
            return
        }

        ctx := context.Background()

        // Creates a client.
        client, err := translate.NewClient(ctx)
        if err != nil {
            log.Fatalf("Failed to create client: %v", err)
        }

        // Gets the user-inputted text to translate.
        text := os.Args[1]
        // Gets flag (if it exists)
        showArabic := false
        if len(os.Args) > 2 {
            flag := os.Args[2]
            if flag == "--show-arabic" {
                showArabic = true
            } else {
                fmt.Println("Flag not found: %v", flag)
                return
            }
        }
        // Sets the target language to Sindhi.
        target, err := language.Parse("sd")
        if err != nil {
            log.Fatalf("Failed to parse target language: %v", err)
        }

        // Translates the text into Sindhi.
        translations, err := client.Translate(ctx, []string{text}, target, nil)
        if err != nil {
            log.Fatalf("Failed to translate text: %v", err)
        }
        if showArabic {
            fmt.Printf("Arabic: %v\n", translations[0].Text)
        }

        // Sets the Sindhi text to translate.
        text = translations[0].Text
        // Sets the target language to Arabic.
        source, _ := language.Parse("ar")
        target, err = language.Parse("en")
        if err != nil {
            log.Fatalf("Failed to parse target language: %v", err)
        }

        // Translates the text into English.
        translations, err = client.Translate(ctx, []string{text}, target, &translate.Options{
            Source: source,
        })
        if err != nil {
            log.Fatalf("Failed to translate text: %v", err)
        }
        // Prints the phonetic Sindhi translation of the user-inputted English text.
        fmt.Printf("Phonetic: %v\n", translations[0].Text)
}