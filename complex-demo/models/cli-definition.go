package models

var CLI struct {
	Generate struct {
		Count int `arg name:"count" help:"Number of ids to generate`
	} `cmd help:"Generate SnowflakeIDs.`

	GenerateFile struct {
		Count int    `arg name:"count" help:"Number of ids to generate`
		Path  string `arg name:"path" help:"The location to generate the file to" type:"path"`
	} `cmd help:"Generate SnowflakeIDs to file.`
}
