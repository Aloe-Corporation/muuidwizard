package decode

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var decodeCmd = cobra.Command{
	Use:   "decode <base64>",
	Args:  cobra.ExactArgs(1),
	Short: "Display the uuid encoded with base64.",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := base64.StdEncoding.DecodeString(args[0])
		if err != nil {
			fmt.Println("can't decode base64 data")
		}

		resultingUUID, err := uuid.FromBytes(res)
		if err != nil {
			fmt.Println("input base64 is not an uuid")
		}

		fmt.Println(resultingUUID)
	},
}

func Init(parent *cobra.Command) {
	parent.AddCommand(&decodeCmd)
}
