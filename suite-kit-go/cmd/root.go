// File: cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "suitectl",
	Short: "itsma suite-kit",
	Long:  `Suitectl is a tool written in Golang, for init and deploy ITSMA services, and CDF as well.`,
}

func Execute(version string) {
	v = version
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(SilentInstallCmd)
	RootCmd.AddCommand(VersionCmd)
}
