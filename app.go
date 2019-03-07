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
        ctx := context.Background()

        // Creates a client.
        client, err := translate.NewClient(ctx)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }

        // Gets the user-inputted text to translate.
        text := os.Args[1]
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
        fmt.Printf("%v\n", translations[0].Text)
}