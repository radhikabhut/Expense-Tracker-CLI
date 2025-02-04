package cli

import (
	"encoding/json"
	"errors"
	"expense-tracker/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func AddDef() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "use for add expense",
		Run:   addFunc,
	}

	addCmd.PersistentFlags().StringP("description", "d", "", "to add description")
	addCmd.PersistentFlags().Float32P("amount", "a", 0, "to add amount")

	return addCmd
}

func addFunc(cmd *cobra.Command, args []string) {

	flagValue := cmd.Flag("description")
	des := flagValue.Value.String()

	flagValue = cmd.Flag("amount")
	amountStr := flagValue.Value.String()
	floatValue, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	createdAt := time.Now()
	month := createdAt.Month().String()
	updatedAt := time.Now()

	id := uuid.NewString()

	exp := model.Expense{
		Id:          id,
		Description: des,
		Amount:      floatValue,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	fmt.Println(exp)

	file, err := os.ReadFile(`F:\workspace\ExpenceTracker\expense.json`)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exists")
		return
	} else if err != nil {
		return
	}

	var expenses map[string][]model.Expense
	err = json.Unmarshal(file, &expenses)
	if err != nil {
		fmt.Println(err)
		return
	}

	monthExpenses := expenses[month]
	if len(monthExpenses) == 0 {
		expenses[month] = []model.Expense{exp}
	} else {
		expenses[month] = append(expenses[month], exp)
	}

	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(`F:\workspace\ExpenceTracker\expense.json`, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Expense added successfully!")

}
