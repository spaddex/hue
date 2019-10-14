package hue

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func init() {
	enable4Windows()
}

func enable4Windows() {
	if runtime.GOOS == "windows" {
		k, err := registry.OpenKey(registry.CURRENT_USER, `Console`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			log.Fatal("Error when trying to open the registry: ", err)
		}
		_, typ, _ := k.GetBinaryValue("VirtualTerminalLevel")
		if typ != 4 {

			if err := k.SetDWordValue("VirtualTerminalLevel", 1); err != nil {
				log.Fatal(`Error when setting "VirtualTerminalLevel"`)
			} else {
				fmt.Println(`We've set the value key "VirtualTerminalLevel" to 1 in "HKEY_CURRENT_USER\Console" to enable colors for windows.`)
				fmt.Println(`Please restart the terminal and you should have colors working`)
				os.Exit(0)
			}
		}
		if err := k.Close(); err != nil {
			log.Fatal("Error on closing registry: ", err)
		}
	}
}

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
