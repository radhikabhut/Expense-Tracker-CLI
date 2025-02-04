package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DeleteDef() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "delete",
		Short: "use for delete expense",
		Run:   deleteFunc,
	}

	addCmd.PersistentFlags().IntP("id", "", 0, "to delete description")

	return addCmd
}

func deleteFunc(cmd *cobra.Command, args []string) {

	flagValue := cmd.Flag("id")
	fmt.Println(flagValue.Value.String())
	fmt.Println(args)

	fmt.Println("delete")
}
