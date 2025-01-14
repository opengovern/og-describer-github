package main

import (
	"github.com/opengovern/og-describer-github/plugin/cloudql/github"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: github.Plugin})
}
