package hue

import "strings"

func Info(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;33m[!]\033[1;m " + allString)
}

func Que(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;34m[?]\033[1;m " + allString)
}

func Bad(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;31m[-]\033[1;m " + allString)
}

func Good(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;32m[+]\033[1;m " + allString)
}

func Run(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;97m[~]\033[1;m " + allString)
}

// Sub status etc

func SubBad(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;31m[->]\033[1;m " + allString)
}

func SubInfo(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;33m[->]\033[1;m " + allString)
}

func SubGood(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[1;32m[->]\033[1;m " + allString)
}

func SubGenInfo(s ...string) string { // # cyan + "->", no judgemental color
	allString := strings.Join(s, " ")
	return ("\033[36m[->]\033[0m " + allString)
}

func Comment(s ...string) string { // ligthblue([*]) +(text)
	allString := strings.Join(s, " ")
	return ("\033[94m[*]\033[0m " + allString)
}

func CommentHighlight(s ...string) string { // lightblue([*] + (text) )
	allString := strings.Join(s, " ")
	return ("\033[94m[*] " + allString + "\033[0m ")
}

// Colors

func Green(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[32m" + allString + "\033[0m")
}

func LightGreen(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[92m" + allString + "\033[0m")
}

func Grey(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[37m" + allString + "\033[0m")
}

func Black(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[30m" + allString + "\033[0m")
}

func Red(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[31m" + allString + "\033[0m")
}

func LightRed(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[91m" + allString + "\033[0m")
}

func Cyan(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[36m" + allString + "\033[0m")
}

func LightCyan(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[96m" + allString + "\033[0m")
}

func Blue(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[34m" + allString + "\033[0m")
}

func LightBlue(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[94m" + allString + "\033[0m")
}

func Purple(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[35m" + allString + "\033[0m")
}

func LightPurple(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[95m" + allString + "\033[0m")
}

func Yellow(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[93m" + allString + "\033[0m")
}

func White(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[97m" + allString + "\033[0m")
}

func Orange(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[33m" + allString + "\033[0m")
}

// Styles

func Bg(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[;7m" + allString + "\033[0m")
}

func Bold(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[;1m" + allString + "\033[0m")
}

func Italic(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[3m" + allString + "\033[0m")
}

func Under(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[4m" + allString + "\033[0m")
}
func Strike(s ...string) string {
	allString := strings.Join(s, " ")
	return ("\033[09m" + allString + "\033[0m")
}
