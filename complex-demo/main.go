package main

import (
	"complex-demo/m/v2/models"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/bwmarrin/snowflake"
)

func main() {
	ctx := kong.Parse(&models.CLI)

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch ctx.Command() {
	case "generate <count>":
	case "generateFile <count> <path>":
	default:
		id := node.Generate()
		fmt.Printf("String ID: %s\n", id)
		fmt.Println(ctx.Command())
	}
	fmt.Printf("Done")
}
