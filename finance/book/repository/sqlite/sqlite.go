package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/printfcoder/home/finance/book/repository"
	"github.com/printfcoder/home/proto/common/page"
	"github.com/printfcoder/home/proto/finance/book"
)

var (
	DefaultDB = "/Users/shuxian/workspace/go/src/github.com/printfcoder/home/.data/finbook.db"
)

type repo struct {
}

func init() {
	repository.Register(new(repo))
}

func (b *repo) Init() (err error) {
	database, err := sql.Open("sqlite3", DefaultDB)
	if err != nil {
		panic(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	return
}

func (b *repo) String() (name string) {
	return "sqlite"
}

func (b *repo) NewExpense(expense book.NewExpenseRequest) (ret book.Expense, err error) {
	return
}

func (b *repo) UpdateExpense(expense book.UpdateExpenseRequest) (err error) {
	return
}

func (b *repo) ListExpenses(req book.ListExpensesRequest) (ret page.Page, err error) {
	return
}

func (b *repo) RemoveExpense(id int64) (err error) {
	return
}
