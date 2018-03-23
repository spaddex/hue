package hue

func Info(s string) string {
	return ("\033[1;33m[!]\033[1;m " + s)
}

func Que(s string) string {
	return ("\033[1;34m[?]\033[1;m " + s)
}

func Bad(s string) string {
	return ("\033[1;31m[-]\033[1;m " + s)
}

func Good(s string) string {
	return ("\033[1;32m[+]\033[1;m " + s)
}

func Run(s string) string {
	return ("\033[1;97m[~]\033[1;m " + s)
}

// Colors

func Green(s string) string {
	return ("\033[32m" + s + "\033[0m")
}

func LightGreen(s string) string {
	return ("\033[92m" + s + "\033[0m")
}

func Grey(s string) string {
	return ("\033[37m" + s + "\033[0m")
}

func Black(s string) string {
	return ("\033[30m" + s + "\033[0m")
}

func Red(s string) string {
	return ("\033[31m" + s + "\033[0m")
}

func LightRed(s string) string {
	return ("\033[91m" + s + "\033[0m")
}

func Cyan(s string) string {
	return ("\033[36m" + s + "\033[0m")
}

func LightCyan(s string) string {
	return ("\033[96m" + s + "\033[0m")
}

func Blue(s string) string {
	return ("\033[34m" + s + "\033[0m")
}

func LightBlue(s string) string {
	return ("\033[94m" + s + "\033[0m")
}

func Purple(s string) string {
	return ("\033[35m" + s + "\033[0m")
}

func LightPurple(s string) string {
	return ("\033[95m" + s + "\033[0m")
}

func Yellow(s string) string {
	return ("\033[93m" + s + "\033[0m")
}

func White(s string) string {
	return ("\033[97m" + s + "\033[0m")
}

func Orange(s string) string {
	return ("\033[33m" + s + "\033[0m")
}

// Styles

func Bg(s string) string {
	return ("\033[;7m" + s + "\033[0m")
}

func Bold(s string) string {
	return ("\033[;1m" + s + "\033[0m")
}

func Italic(s string) string {
	return ("\033[3m" + s + "\033[0m")
}

func Under(s string) string {
	return ("\033[4m" + s + "\033[0m")
}
func Strike(s string) string {
	return ("\033[09m" + s + "\033[0m")
}
