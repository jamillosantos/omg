package cmd

import (
	"github.com/jamillosantos/omg/external/protodep"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Uses protodep to install all proto dependencies",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := protodep.Up()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
