package main

// Importing a package which contains the Println function that we use to print
import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello!, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// In our function signature we have made a named return value (prefix string)
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Jay", "Spanish"))
	fmt.Println(Hello("Jay", "French"))
	fmt.Println(Hello("Jay", ""))
	fmt.Println(Hello("", ""))
}
