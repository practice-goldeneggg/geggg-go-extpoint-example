package main

import (
	"fmt"
	"os"

	_ "github.com/practice-goldeneggg/additionalextpoint"
	"github.com/practice-goldeneggg/geggg-go-extpoint-example/extpoints"
)

var (
	hoges = extpoints.Hoges
)

func main() {
	var s int
	defer os.Exit(s)

	for _, registered := range hoges.All() {
		fmt.Println("Registered object: ", registered)
	}

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Required name argument")
		s = 1
		return
	}

	name := args[0]
	hoge, ok := hoges.Lookup(name)
	if !ok {
		fmt.Fprintf(os.Stderr, "Name %s is not registered\n", name)
		s = 1
		return
	}

	if err := hoge.Before(args); err != nil {
		fmt.Fprintln(os.Stderr, "Before error", err)
		s = 1
		return
	}

	if err := hoge.Run(args); err != nil {
		fmt.Fprintln(os.Stderr, "Run error", err)
		s = 1
		return
	}

	if err := hoge.After(args); err != nil {
		fmt.Fprintln(os.Stderr, "After error", err)
		s = 1
		return
	}
}
