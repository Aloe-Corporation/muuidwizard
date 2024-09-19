package generate

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var generateCmd = cobra.Command{
	Use:   "generate",
	Short: "generate a random uuid and display it",
	Run: func(cmd *cobra.Command, args []string) {
		uuidToEncode := uuid.New()
		fmt.Println(uuidToEncode)
		fmt.Println(base64.StdEncoding.EncodeToString(uuidToEncode[:]))
	},
}

func Init(parent *cobra.Command) {
	parent.AddCommand(&generateCmd)
}
