//md5sum tool takes a file name as an input and gives the md5 hash of the file content as the output.
//If input file name is not specified, it will take an executable as an input file and return its md5 hash value as an output.

package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

// Exit codes.
const (
	Success = iota
	Failed
)

func main() {
	input := os.Stdin
	fileStat, _ := input.Stat()
	fm := fileStat.Mode()
	var fname string

	if fm&os.ModeNamedPipe == os.ModeNamedPipe {
		fname = input.Name()
	} else if fm&os.ModeCharDevice == os.ModeCharDevice {

		if len(os.Args) > 1 {
			fname = os.Args[1]
		} else {
			fname = os.Args[0]
		}

	} else {
		fname = input.Name()
	}

	inputFile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(Failed)
	}
	defer inputFile.Close()

	Buf, err := ioutil.ReadFile(inputFile.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(Failed)
	}

	fmt.Printf("md5 hash = %x\n", md5.Sum([]byte(Buf)))
}
