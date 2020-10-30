package cmd

import (
	"os"

	"github.com/jamillosantos/omg/external/buf"
	"github.com/spf13/cobra"
)

var typeFlag string

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Uses buf to run a lint on the configured files",
	Long: `For now, there is no way to configure what lints are going to be active while running this command.
	
But, it is possible to create a 'buf.yaml' at the root of your project and the buf itself will load it. Check buf
documentation at https://buf.build.`,
	Run: func(cmd *cobra.Command, args []string) {
		ec, err := buf.Lint(typeFlag)
		if err != nil {
			panic(err)
		}
		os.Exit(ec)
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)

	lintCmd.Flags().StringVar(&typeFlag, "type", "text", "Define the output type for the errors. (text,plain)")
}
