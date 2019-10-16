package hue

import (
	"fmt"
	"strings"
	"time"
)

func Info(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;33m[!]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Que(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;34m[?]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Bad(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;31m[-]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Good(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;32m[+]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Run(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;97m[~]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

// Sub status etc

func SubBad(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;31m[->]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func SubInfo(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;33m[->]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func SubGood(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[1;32m[->]\033[1;m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func SubGenInfo(s ...interface{}) string { // # cyan + "->", no judgemental color
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[36m[->]\033[0m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Comment(s ...interface{}) string { // ligthblue([*]) +(text)
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[94m[*]\033[0m %v", strings.Trim(fmt.Sprintf(str), "[]"))
}

func CommentHighlight(s ...interface{}) string { // lightblue([*] + (text) )
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[94m[*] %v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func PrependTime(s ...string) string {
	allString := strings.Join(s, " ")
	currentTime := time.Now()
	return (fmt.Sprintf("%v - %v", currentTime.Format("2006-01-02 15:04:05"), allString))
}

// Colors

func Green(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[32m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func LightGreen(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[92m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Grey(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[37m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Black(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[30m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Red(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[31m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func LightRed(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[91m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Cyan(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[36m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func LightCyan(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[96m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Blue(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[34m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func LightBlue(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[94m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Purple(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[35m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func LightPurple(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[95m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Yellow(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[93m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func White(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[97m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Orange(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[33m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

// Styles

func Bg(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[;7m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Bold(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[;1m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Italic(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[3m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}

func Under(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[4m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}
func Strike(s ...interface{}) string {
	str := fmt.Sprintf("%v", s)
	return fmt.Sprintf("\033[09m%v\033[0m", strings.Trim(fmt.Sprintf(str), "[]"))
}
