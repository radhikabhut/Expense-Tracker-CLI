package cli

import (
	"github.com/spf13/cobra"
)

func InitCli() {
	rootCommand := &cobra.Command{
		Use: "expense-tracker",
	}

	// deleteSubCommand := &cobra.Command{
	// 	Use:   "delete",
	// 	Short: "use for delete expense",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("delete")
	// 	},
	// }

	rootCommand.AddCommand(AddDef(), DeleteDef(), ListDef(), UpdateDef(), SummaryDef())
	// rootCommand.AddCommand(deleteSubCommand)
	rootCommand.Execute()
}
