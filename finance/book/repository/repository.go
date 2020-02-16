package repository

import (
	"github.com/printfcoder/home/proto/common/page"
	"github.com/printfcoder/home/proto/finance/book"
)

type BookRepository interface {
	Init() (err error)
	NewExpense(expense book.NewExpenseRequest) (ret book.Expense, err error)
	UpdateExpense(expense book.UpdateExpenseRequest) (err error)
	ListExpenses(req book.ListExpensesRequest) (ret page.Page, err error)
	RemoveExpense(id int64) (err error)
}