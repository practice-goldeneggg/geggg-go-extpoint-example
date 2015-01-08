package main

import (
	"fmt"

	"github.com/goldeneggg/geggg-go-extpoint-example/extpoints"
)

func init() {
	hoges.Register(new(gekiteki), "")
}

type gekiteki struct{}

var _ extpoints.Hoge = new(gekiteki)

func (g *gekiteki) Before(args []string) error {
	fmt.Println("ビフォー")
	return nil
}

func (g *gekiteki) Run(args []string) error {
	fmt.Println("匠の劇的な技！")
	return nil
}

func (g *gekiteki) After(args []string) error {
	fmt.Println("アフター")
	return nil
}

func (g *gekiteki) String() string {
	return "劇的！ビフォーアフター"
}
