package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func hello(language,name string) string {
	if name == "" {
			name = "World"
	}

	return grettingPrefix(language) + name
	
}


func grettingPrefix(language string) (prefix string) {
	switch language {
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}


func main () {
	fmt.Println(hello("Spanish","Elliot"))
}

