package service

import (
	"github.com/printfcoder/home/finance/book/repository"
	"github.com/printfcoder/home/proto/common/page"
	"github.com/printfcoder/home/proto/finance/book"
)

var (
	repo = repository.Repo()
)

func AddExpense(req book.AddExpenseRequest) (ret book.Expense, err error) {
	ret, err = repo.NewExpense(req)
	if err != nil {

	}

	return
}

func DeleteExpense(req book.DeleteExpenseRequest) (err error) {
	err = repo.DeleteExpense(req.Id)
	if err != nil {

	}

	return
}

func UpdateExpense(req book.UpdateExpenseRequest) (err error) {
	err = repo.UpdateExpense(req)
	if err != nil {

	}

	return
}

func ListExpenses(req book.ListExpensesRequest) (ret page.Page, err error) {
	ret, err = repo.ListExpenses(req)
	if err != nil {

	}

	return
}
