package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
)

var cost = flag.Int("c", 10, "bcrypt cost")
var content = flag.String("p", "", "password to be crypt or from stdin")
var hashedPassword = flag.String("h", "", "hashed password")

func main() {
	flag.Parse()

	var password = *content
	if len(password) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "fail to read strings from stdin")
			os.Exit(1)
		}
		password = string(b)
	}
	if len(*hashedPassword) > 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(password)); err == nil {
			fmt.Fprintln(os.Stderr, "It's right")
		} else {
			fmt.Fprintln(os.Stderr, "It's wrong: "+err.Error())
			os.Exit(1)
		}
	} else {
		bHash, err := bcrypt.GenerateFromPassword([]byte(password), *cost)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bcrypt err: "+err.Error())
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%v\n", string(bHash))
	}
}
