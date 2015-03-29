package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		error_output := "File argument not present"
		fmt.Println(error_output)
		os.Exit(1)
	}
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		error_output := "'" + file + "' File does not exist"
		fmt.Println(error_output)
		os.Exit(1)
	}
	bytes := sha256.Sum256(data)
	sha := hex.EncodeToString(bytes[:])
	output := sha + " " + file
	fmt.Println(output)
}


