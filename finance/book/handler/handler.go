package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/printfcoder/home/finance/book/service"
	"github.com/printfcoder/home/proto/common/response"
	"github.com/printfcoder/home/proto/finance/book"
)

type BookHandler struct {
}

func (b BookHandler) AddExpense(ctx context.Context, req *book.AddExpenseRequest, rsp *response.Response) error {
	ret, err := service.AddExpense(*req)
	if err != nil {

	}

	rsp.Data, _ = ptypes.MarshalAny(&ret)

	return nil
}

func (b BookHandler) DeleteExpense(ctx context.Context, req *book.DeleteExpenseRequest, rsp *response.Response) error {
	panic("implement me")
}

func (b BookHandler) UpdateExpense(ctx context.Context, req *book.UpdateExpenseRequest, rsp *response.Response) error {
	panic("implement me")
}

func (b BookHandler) ListExpenses(ctx context.Context, req *book.ListExpensesRequest, rsp *response.Response) error {
	panic("implement me")
}
