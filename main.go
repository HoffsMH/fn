package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"prepend"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fn",
	Short: "file name utility",
	Long: `stuff`,
	Run: mv,
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
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
}

func init() {
	rootCmd.AddCommand(prependDateCmd)
}

func mv(cmd *cobra.Command, args []string) {
	fmt.Println("hi")
}

