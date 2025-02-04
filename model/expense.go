package model

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          string    `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

func (exp Expense) String() string {
	return fmt.Sprintf("ID: %s | Amount: %.2f | Discription: %s | CreatedAt: %s | UpdatedAt: %s\n",
		exp.Id, exp.Amount, exp.Description, exp.CreatedAt, exp.UpdatedAt)
}
