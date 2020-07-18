package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

func main() {
	var filename = flag.String("file", "", "File to be shared")
	var shareNumber = flag.Int64("share", 0, "Number of people to share the file with")
	var minNumberToRecover = flag.Int64("recover", 0, "Minimum number of people to recover the shared file")

	flag.Parse()

	if *filename == "" {
		fmt.Fprint(os.Stderr, "a file to be shared must be provided")
		os.Exit(1)
	}

	if *shareNumber < 2 {
		fmt.Fprint(os.Stderr, "number of people to share the file with must be at least 2")
		os.Exit(1)
	}

	if *minNumberToRecover < 2 {
		fmt.Fprint(os.Stderr, "minimum number of people to recover the shared file must be at least 2")
		os.Exit(1)
	}

	if *shareNumber < *minNumberToRecover {
		fmt.Fprint(os.Stderr, "minimum number of people to recover the shared file cannot be bigger than the number of people that the file is being shared with")
		os.Exit(1)
	}

	fmt.Println("Congrats!")
	var bs, err = ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problems found reading the file: %v", err)
		os.Exit(1)
	}

	var base = big.Int{}
	base.SetBytes(bs)
	fmt.Printf("Data read: %v\n", base)

	coefficients := make([]*big.Int, *minNumberToRecover - int64(1))
	for i := range coefficients {
		tmp, err := randomBigInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to generate random number")
			os.Exit(1)
		}
		coefficients[i] = tmp
	}

	for i := range coefficients {
		fmt.Printf("%v\n", coefficients[i])
	}
}
