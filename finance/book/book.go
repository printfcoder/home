package book

import (
	"github.com/printfcoder/home/finance/book/repository"
	_ "github.com/printfcoder/home/finance/book/repository/sqlite"
)

func Init() {
	repository.Init()
}
