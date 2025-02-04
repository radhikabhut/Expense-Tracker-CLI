package cli

import (
	"encoding/json"
	"errors"
	"expense-tracker/model"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func SummaryDef() *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "use for summary of expence",
		Run:   summaryFunc,
	}
	summaryCmd.PersistentFlags().StringP("month", "m", "", "to list month wise expence")
	return summaryCmd
}
func summaryFunc(cmd *cobra.Command, args []string) {

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
		totalAmount := 0.0
		if len(expenses) != 0 {
			monthExpenses, exist := expenses[month]
			if exist {
				for _, exp := range monthExpenses {
					totalAmount = totalAmount + exp.Amount
				}
			}
		} else {
			fmt.Println("no expences found")
		}
		
		fmt.Println(totalAmount)

	} else {
		totalAmount := 0.0
		for _, monthExpences := range expenses {
			// fmt.Println("====================")
			// fmt.Println(month)
			for _, exp := range monthExpences {
				totalAmount = totalAmount + exp.Amount

			}
			// fmt.Println("====================")
		}

		fmt.Println(totalAmount)
	}
}
