package main

func colors() (string, string, string) {
	return "red", "yellow", "blue"
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}
