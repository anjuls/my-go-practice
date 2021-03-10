package main

const englishPrefixSalute = "Hello, "
const spanishPrefixSalute = "Hola, "

func Chris(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "" {
		lang = "English"
	}

	prefix := englishPrefixSalute
	switch lang {
	case "English":
		prefix = englishPrefixSalute
	case "Spanish":
		prefix = spanishPrefixSalute
	default:
		prefix = "not implemented. "
	}

	return prefix + name
}
