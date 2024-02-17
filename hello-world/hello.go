package main

import (
	"fmt"
)

const defaultName string = "World"
const spanish string = "Spanish"
const french string = "French"
const brazilianPortuguese string = "Brazilian"
const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "
const brazilianPortugueseHelloPrefix string = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case brazilianPortuguese:
		prefix = brazilianPortugueseHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
