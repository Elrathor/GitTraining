package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/bwmarrin/snowflake"
)

var CLI struct {
	Generate struct {
		Count int `arg name:"count" help:"Number of ids to generate`
	} `cmd help:"Generate SnowflakeIDs.`

	GenerateFile struct {
		Count int    `arg name:"count" help:"Number of ids to generate`
		Path  string `arg name:"path" help:"The location to generate the file to" type:"path"`
	} `cmd help:"Generate SnowflakeIDs to file.`
}

func main() {
	ctx := kong.Parse(&CLI)

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
