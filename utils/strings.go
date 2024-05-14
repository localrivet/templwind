package utils

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// toCamel converts string to camelCase
func ToCamel(s string) string {
	words := splitIntoWords(s)
	caser := cases.Title(language.English)
	for i := 1; i < len(words); i++ {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}

// toKebab converts string to kebab-case
func ToKebab(s string) string {
	return strings.Join(splitIntoWords(s), "-")
}

// toTitle converts string to Title Case
func ToTitle(s string) string {
	words := splitIntoWords(s)
	caser := cases.Title(language.English)
	for i := range words {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, " ")
}

// toSnake converts string to snake_case
func ToSnake(s string) string {
	return strings.Join(splitIntoWords(s), "_")
}

// toPascal converts string to PascalCase
func ToPascal(s string) string {
	words := splitIntoWords(s)
	caser := cases.Title(language.English)
	for i := range words {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}

// toConstant converts string to CONSTANT_CASE
func ToConstant(s string) string {
	return strings.ToUpper(strings.Join(splitIntoWords(s), "_"))
}

// splitIntoWords splits the string into words based on non-alphanumeric characters
func splitIntoWords(s string) []string {
	return regexp.MustCompile("[^a-zA-Z0-9]+").Split(strings.ToLower(s), -1)
}
