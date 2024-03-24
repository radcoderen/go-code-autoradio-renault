package main

import (
	"fmt"
	"os"

	"github.com/radcoderen/go-code-autoradio-renault"
)

const help string = `renault code generate a validation code for Renault radios.

Usage:
renaultcode <precode> [<precode> ...]
renaultcode -a | --all
renaultcode -h | --help

Options:
<precode>  The precode to generate the radio code.
-a --all   Generate all the radio codes.
-h --help  Show this screen.

The precode must be a string of 4 characters. 
The first character must be an upper case letter 
and the following 3 characters must be digits.
The precode must not start with A0.

Examples:
renaultcode A123 B456 C789 Z999
renaultcode --all
`

func main() {
	// Check if the user provided an argument.
	if len(os.Args) < 2 {
		fmt.Println(help)
		return
	}
	var arg string
	// Check if the user asked for help.
	arg = os.Args[1]
	if arg == "-h" || arg == "--help" {
		fmt.Println(help)
		return
	}
	// Check if the user asked for all the codes.
	if arg == "-a" || arg == "--all" {
		for i := 'A'; i <= 'Z'; i++ {
			for j := '0'; j <= '9'; j++ {
				for k := '0'; k <= '9'; k++ {
					for l := '0'; l <= '9'; l++ {
						precode := fmt.Sprintf("%c%c%c%c", i, j, k, l)
						code, ok := codeautoradio.GetRadioCode(precode)
						if ok {
							fmt.Printf("%s %s\n", precode, code)
						}
					}
				}
			}
		}
		return
	}
	// Generate the radio code for the given precodes.
	for _, arg = range os.Args[1:] {
		code, ok := codeautoradio.GetRadioCode(arg)
		if ok {
			fmt.Println(code)
		} else {
			fmt.Println(fmt.Sprintf("The precode %s is not valid.", arg))
		}
	}
}
