package prepend

import (
	"github.com/spf13/cobra"
	"fmt"
)

var prependDateCmd = &cobra.Command{
	Use:     "prependdate",
	Aliases: []string{"pd"},
	Short:   "Downloads an episode of a podcast.",
	Long: ``,
	Args: cobra.MinimumNArgs(1),
	Run:  prependDate,
}

func prependDate(cmd *cobra.Command, args []string) {
	fmt.Println("prependDate");
	fmt.Println(args[0]);
}
