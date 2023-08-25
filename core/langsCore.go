package core

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Vanilla OS Contributors <https://github.com/vanilla-os/>
	Copyright: 2023
	Description:
		Chronos is a simple, fast and lightweight documentation server written in Go.
*/

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/text/language"
)

// getPreferredLanguage retrieves the preferred language from the request headers.
func getPreferredLanguage(r *http.Request) string {
	acceptLanguage := r.Header.Get("Accept-Language")
	if acceptLanguage != "" {
		tags, _, err := language.ParseAcceptLanguage(acceptLanguage)
		if err == nil && len(tags) > 0 {
			lang := tags[0].String()
			lang = strings.Split(lang, "-")[0]
			if isLanguageSupported(lang) {
				return lang
			}
		}
	}
	return "en"
}

// PopulateSupportedLanguages populates the list of supported languages based on the articles directory.
func PopulateSupportedLanguages() error {
	dirEntries, err := os.ReadDir(articlesDir)
	if err != nil {
		return err
	}

	fmt.Println("Loading supported languages...")
	for _, entry := range dirEntries {
		if entry.IsDir() {
			SupportedLang = append(SupportedLang, entry.Name())
			fmt.Printf("- [%s] found\n", entry.Name())
		}
	}
	fmt.Println("Supported languages loaded.")

	return nil
}

// isLanguageSupported checks if a given language is in the list of supported languages.
func isLanguageSupported(lang string) bool {
	for _, l := range SupportedLang {
		if l == lang {
			return true
		}
	}
	return false
}
