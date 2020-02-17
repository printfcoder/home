package repository

import (
	"sync"

	"github.com/printfcoder/home/proto/common/page"
	"github.com/printfcoder/home/proto/finance/book"
)

type BookRepository interface {
	Init() (err error)
	String() (name string)
	NewExpense(expense book.NewExpenseRequest) (ret book.Expense, err error)
	UpdateExpense(expense book.UpdateExpenseRequest) (err error)
	ListExpenses(req book.ListExpensesRequest) (ret page.Page, err error)
	RemoveExpense(id int64) (err error)
}

var (
	DefaultRepo = "sqlite"
	mux         sync.Mutex
	repoMap     = map[string]BookRepository{}
)

func Register(rep BookRepository) {
	mux.Lock()
	defer mux.Unlock()

	repoMap[rep.String()] = rep
}

func Init() {
	repoMap[DefaultRepo].Init()
}
