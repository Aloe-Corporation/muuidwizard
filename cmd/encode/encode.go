package encode

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var encodeCmd = cobra.Command{
	Use:   "encode <uuid>",
	Short: "Display the base64 representation of the uuid.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uuidToEncode, err := uuid.Parse(args[0])
		if err != nil {
			fmt.Println("argument is not a valid uuid")
			os.Exit(1)
		}

		fmt.Println(base64.StdEncoding.EncodeToString(uuidToEncode[:]))
	},
}

func Init(parent *cobra.Command) {
	parent.AddCommand(&encodeCmd)
}
