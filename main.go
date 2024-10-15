package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Struct to store localized messages
type Localizer struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

// NewLocalizer
func NewLocalizer(lang string) *Localizer {
	// Create a new i18n bundle for the default language (English in this case)
	bundle := i18n.NewBundle(language.English)

	// Register the unmarshal function for JSON (using the standard encoding/json package)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load the translation files for each language
	_, err := bundle.LoadMessageFile("locales/en.json")
	if err != nil {
		log.Fatalf("Error loading English messages: %v", err)
	}
	_, err = bundle.LoadMessageFile("locales/es.json")
	if err != nil {
		log.Fatalf("Error loading Spanish messages: %v", err)
	}

	// Create a new localizer for the selected language
	localizer := i18n.NewLocalizer(bundle, lang)
	return &Localizer{
		bundle:    bundle,
		localizer: localizer,
	}
}

// GetMessage
func (l *Localizer) GetMessage(messageID string) string {
	localizedMessage, err := l.localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})

	if err != nil {
		log.Printf("Error: Message '%s' not found: %v", messageID, err)
		return fmt.Sprintf("Message '%s' not found", messageID)
	}
	return localizedMessage
}

func main() {
	preferredLang := "es"
	localizer := NewLocalizer(preferredLang)

	// Display localized messages
	fmt.Println(localizer.GetMessage("welcome_message"))
}
