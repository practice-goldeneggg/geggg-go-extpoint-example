package main

import (
	"fmt"

	"github.com/practice-goldeneggg/geggg-go-extpoint-example/extpoints"
)

func init() {
	hoges.Register(new(xyz), "")
}

type xyz struct{}

var _ extpoints.Hoge = new(xyz)

func (g *xyz) Before(args []string) error {
	fmt.Println("Exec Before")
	return nil
}

func (g *xyz) Run(args []string) error {
	fmt.Println("Exec Run")
	return nil
}

func (g *xyz) After(args []string) error {
	fmt.Println("Exec After")
	return nil
}

func (g *xyz) String() string {
	return "I am xyz"
}
