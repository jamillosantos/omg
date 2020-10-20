package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jamillosantos/omg/config"
	"github.com/jamillosantos/omg/internal"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "omg",
	Short: "A helper for dealing with gRPC code generation",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Verbose = verbose

		configFile, err := os.Open("omg.yaml")
		if err != nil {
			internal.Fatal(1, "error opening omg.yaml: ", err)
		}
		defer configFile.Close()

		configData, err := ioutil.ReadAll(configFile)
		if err != nil {
			internal.Fatal(2, "error reading config data: ", err)
		}

		if err := yaml.Unmarshal(configData, &config.Config); err != nil {
			internal.Fatal(3, "error parsing omg.yaml: ", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/omg.yaml)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "set verbose mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
