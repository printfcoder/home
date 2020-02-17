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

	expenseCreateSQL = `create table IF NOT EXISTS expense
(
    id           integer
        constraint expense_pk
            primary key autoincrement,
    label        integer,
    value        NUMERIC not null,
    account_type text,
    time         NUMERIC
);`
	memberCreateSQL = `create table IF NOT EXISTS member
(
    id   integer
        constraint member_pk
            primary key autoincrement,
    name text not null
);`
	exMemCreateSQL = `create table IF NOT EXISTS expense_member
(
    id         integer
        constraint expense_members_pk
            primary key autoincrement,
    expense_id integer not null,
    member_id  integer not null
);
`
	createSQLs = []string{expenseCreateSQL, memberCreateSQL, exMemCreateSQL}
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

	for _, s := range createSQLs {
		func() {
			stmt, err := database.Prepare(s)
			if err != nil {
				panic(err)
			}
			defer stmt.Close()

			stmt.Exec()
		}()
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
