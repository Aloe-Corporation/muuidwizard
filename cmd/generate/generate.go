package generate

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var generateCmd = cobra.Command{
	Use:   "generate <number>",
	Short: "generate <number> random uuid and display it",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		count, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("argument is not a valid number")
			os.Exit(1)
		}
		for i := 0; i < count; i++ {
			uuidToEncode := uuid.New()
			fmt.Println(uuidToEncode)
		}
	},
}

func Init(parent *cobra.Command) {
	parent.AddCommand(&generateCmd)
}
