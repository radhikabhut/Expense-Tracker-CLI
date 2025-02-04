package cli

import (
	"encoding/json"
	"errors"
	"expense-tracker/model"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ListDef() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "use for list expence",
		Run:   listFunc,
	}
	listCmd.PersistentFlags().StringP("month", "m", "", "to list month expence")

	return listCmd
}
func listFunc(cmd *cobra.Command, args []string) {

	flagValue := cmd.Flag("month")
	month := flagValue.Value.String()

	file, err := os.ReadFile(`F:\workspace\ExpenceTracker\expense.json`)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exists")
		return
	} else if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var expenses map[string][]model.Expense
	err = json.Unmarshal(file, &expenses)
	if err != nil {
		fmt.Println(err)
		return
	}

	if month != "" {
		if len(expenses) != 0 {
			monthExpenses, exist := expenses[month]
			if exist {
				for _, exp := range monthExpenses {
					fmt.Println(exp)
				}
			}
		} else {
			fmt.Println("no expences found")
		}
	} else {
		for month, monthExpences := range expenses {
			fmt.Println("====================")
			fmt.Println(month)
			for _, exp := range monthExpences {
				fmt.Println(exp)
			}
			fmt.Println("====================")
		}
	}

}
