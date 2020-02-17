package sqlite

import (
	"database/sql"
	"fmt"
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

	insertExpenseSQL = `INSERT INTO expense (label, value, account_type, time) VALUES (?, ?, ?, ?)`
	insertExMemSQL   = `INSERT INTO expense_member (expense_id, member_id) VALUES (?, ?)`

	updateExpenseSQL = `UPDATE expense SET label = ?, value = ?, account_type = ?, time = ? WHERE id = ? `

	deleteExMemSQL = `DELETE FROM expense_member WHERE expense_id = ? AND member_id = ?`

	db *sql.DB
)

type repo struct {
}

func init() {
	repository.Register(new(repo))
}

func (b *repo) Init() (err error) {
	db, err = sql.Open("sqlite3", DefaultDB)
	if err != nil {
		panic(err)
	}

	for _, s := range createSQLs {
		func() {
			stmt, err := db.Prepare(s)
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

func (b *repo) NewExpense(req book.NewExpenseRequest) (ret book.Expense, err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("[NewExpense] err: %s", err)
			// log
			_ = tx.Rollback()
		}
	}()

	// expense data
	result, err := tx.Exec(insertExpenseSQL, req.Expense.Label, req.Expense.Value, req.Expense.AccountType, req.Expense.Time)
	if err != nil {
		return
	}

	expenseId, err := result.LastInsertId()
	if err != nil {
		return
	}

	err = insertOrDeleteMembers(tx, insertExMemSQL, expenseId, req.Expense.Members)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func (b *repo) UpdateExpense(req book.UpdateExpenseRequest) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("[NewExpense] err: %s", err)
			// log
			_ = tx.Rollback()
		}
	}()

	// expense data
	_, err = tx.Exec(updateExpenseSQL, req.Expense.Label, req.Expense.Value, req.Expense.AccountType,
		req.Expense.Time, req.Expense.Id)
	if err != nil {
		return
	}

	// delete all old members
	err = insertOrDeleteMembers(tx, deleteExMemSQL, req.Expense.Id, req.Expense.Members)
	if err != nil {
		return
	}

	// insert the new members
	err = insertOrDeleteMembers(tx, insertExMemSQL, req.Expense.Id, req.Expense.Members)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func (b *repo) ListExpenses(req book.ListExpensesRequest) (ret page.Page, err error) {
	return
}

func (b *repo) RemoveExpense(id int64) (err error) {
	return
}

func insertOrDeleteMembers(tx *sql.Tx, insertOrDeleteSQL string, expenseId int64, memberIds []int64) (err error) {
	for _, id := range memberIds {
		err = func() (err error) {
			state, err := tx.Prepare(insertExMemSQL)
			if err != nil {
				return
			}
			defer state.Close()

			_, err = state.Exec(expenseId, id)
			if err != nil {
				return
			}
			return
		}()

		if err != nil {
			return
		}
	}

	return
}
