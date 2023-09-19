package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version info.",
	Long:  `A very detailed version of this tool, including build date, go version, kit version...`,
	Run: func(cmd *cobra.Command, _ []string) {
		listVersion()
	},
}

func listVersion() {
	fmt.Println("Tool name:\t", toolName)
	fmt.Println("Kit version:\t", v)
	fmt.Println("Go version:\t", goVersion)
	fmt.Println("Build date:\t", buildDate)
	fmt.Println("Git repo:\t", repo)
}
