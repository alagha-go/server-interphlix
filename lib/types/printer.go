package types

import "fmt"

var (
	colorReset = "\033[0m"

    colorRed = "\033[31m"
    colorGreen = "\033[32m"
    colorYellow = "\033[33m"
    colorBlue = "\033[34m"
    colorPurple = "\033[35m"
    colorCyan = "\033[36m"
    colorWhite = "\033[37m"
)



func PrintGreen(value interface{}) {
	fmt.Printf("%s%v%s\n",colorGreen, value, colorReset)
}


func PrintRed(value interface{}) {
	fmt.Printf("%s%v%s\n",colorRed, value, colorReset)
}


func PrintYellow(value interface{}) {
	fmt.Printf("%s%v%s\n",colorYellow, value, colorReset)
}

func PrintBlue(value interface{}) {
	fmt.Printf("%s%v%s\n",colorBlue, value, colorReset)
}

func PrintPurple(value interface{}) {
	fmt.Printf("%s%v%s\n",colorPurple, value, colorReset)
}

func PrintCyan(value interface{}) {
	fmt.Printf("%s%v%s\n",colorCyan, value, colorReset)
}

func PrintWhite(value interface{}) {
	fmt.Printf("%s%v%s\n",colorWhite, value, colorReset)
}