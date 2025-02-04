package cli

import (
	"encoding/json"
	"errors"
	"expense-tracker/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func UpdateDef() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "use for update expence",
		Run:   updateFunc,
	}
	updateCmd.PersistentFlags().StringP("id", "i", "", "to add id")
	updateCmd.PersistentFlags().StringP("description", "d", "", "to add description")
	updateCmd.PersistentFlags().Float32P("amount", "a", 0, "to add amount")

	return updateCmd
}
func updateFunc(cmd *cobra.Command, args []string) {

	flagValue := cmd.Flag("id")
	id := flagValue.Value.String()

	flagValue = cmd.Flag("description")
	des := flagValue.Value.String()

	flagValue = cmd.Flag("amount")
	amountStr := flagValue.Value.String()
	// fmt.Println("Amount received:", amountStr) // Debugging output
	floatValue, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	
	if id == "" {
		fmt.Println("pleae provide Id to update")
		return
	}

	file, err := os.ReadFile(`F:\workspace\ExpenceTracker\expense.json`)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exists")
		return
	} else if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var expences map[string][]model.Expense
	err = json.Unmarshal(file, &expences)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}


	found := false
	for month, monthExpences := range expences {
		for index, exp := range monthExpences {
			if exp.Id == id {
				if des != "" {
					exp.Description = des
				}
				if floatValue != 0.0 {
					exp.Amount = floatValue
				}
				exp.UpdatedAt = time.Now()
				expences[month][index] = exp
				found = true
				break
			}
		}

	}
	if !found {
		fmt.Println("expense with given id is not found")
		return
	}
	fmt.Println("Updated expenses:", expences)
	data, err := json.MarshalIndent(expences, "", " ")
	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	err = os.WriteFile(`F:\workspace\ExpenceTracker\expense.json`, data, 0644)
	if err != nil {
		fmt.Println("Error saving updated expenses:", err)
		return
	}

	fmt.Println("updated sucessfully")
}
