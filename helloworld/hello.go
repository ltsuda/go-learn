package main

import (
	"fmt"
)

const (
	defaultName         string = "World"
	spanish             string = "Spanish"
	french              string = "French"
	brazilianPortuguese string = "Brazilian"

	englishHelloPrefix             string = "Hello, "
	spanishHelloPrefix             string = "Hola, "
	frenchHelloPrefix              string = "Bonjour, "
	brazilianPortugueseHelloPrefix string = "Olá, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case brazilianPortuguese:
		prefix = brazilianPortugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
