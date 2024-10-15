
# Go Localization Example

This demonstrates how to implement localization (L10n) in a Go application using the `go-i18n` package. The example shows how to load translations from language-specific resource files (JSON) and display localized messages based on the user's preferred language.

## Features

- Support for multiple languages (English, Spanish in this example).
- Dynamically load and switch between languages.
- Store all translations in external JSON files.
- Easily extendable to support more languages.
- Demonstrates handling translation for user-facing messages.

## Getting Started

### Prerequisites

- [Go 1.16+](https://golang.org/dl/) installed on your machine.

### Installing

1. Clone the repository:
   ```bash
   git clone https://github.com/balasl342/go-i18n-localization-example
   ```

2. Install the required Go package:
   ```bash
   go get -u github.com/nicksnyder/go-i18n/v2/goi18n
   ```

3. Prepare the translation files in JSON format (already included in the repository):

- **en.json** (English):
   ```json
   {
    "welcome_message": "Welcome",
    "goodbye_message": "Goodbye"
  }
   ```

- **es.json** (Spanish):
   ```json
    {
      "welcome_message": "¡Bienvenido",
      "goodbye_message": "¡Gracias"
    }
   ```

### Running the Example

1. Open `main.go` and set the preferred language to either "en" (English) or "es" (Spanish):

   ```go
   preferredLang := "es" // For Spanish; change to "en" for English.
   ```

2. Run the Go application:

   ```bash
   go run main.go
   ```

3. The application will display localized messages depending on the selected language.

### Extending

To add support for more languages:

1. Create a new JSON file for the desired language (e.g., `fr.json` for French).
2. Load the new translation file in the Go application using `LoadMessageFile`.
3. Set the `preferredLang` variable to the language code (e.g., "fr").

### License

This project is licensed under the MIT License.
