// AddLongString - add every singel digit in a long number
// Exempel 123456789 return 45
package main

import (
	"fmt"
	"math/big"
	"os"
)

func usage() {
	fmt.Printf("\nSum of every singel digit in a long number\n")
	fmt.Printf("%s [integer]\n", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("ERR:Missing argument\n")
		usage()
		os.Exit(1)
	}

	// TODO more error check it can delete space and remove ascii
	if os.Args[1][0] == '-' {
		fmt.Printf("ERR:Converting %s to positive number", os.Args[1])
		usage()
		os.Exit(1)
	}

	numStr := os.Args[1]

	sumBig := big.NewInt(int64(0))
	for _, c := range numStr {
		s := string(c)
		addBig := new(big.Int)
		addBig, ok := addBig.SetString(s, 10)
		if !ok || os.Args[1][0] == '-' {
			fmt.Printf("ERR:Converting %s to positive number", os.Args[1])
			usage()
			os.Exit(1)
		}
		sumBig.Add(sumBig, addBig)
	}

	fmt.Println(sumBig.String())

}
