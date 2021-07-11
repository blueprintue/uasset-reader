package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/blueprintue/uasset-reader/internal"
)

var CLI struct {
	Read struct {
		Filepath string `arg name:"filepath" help:"Filepaths to read." type:"path"`
	} `cmd help:"Read uasset file."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "read <filepath>":
		fmt.Println(internal.ReadUasset(CLI.Read.Filepath))
	default:
		panic(ctx.Command())
	}
}
