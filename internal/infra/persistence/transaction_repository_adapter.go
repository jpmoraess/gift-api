package persistence

import (
	"context"
	"fmt"
	db "github.com/jpmoraess/gift-api/db/sqlc"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type TransactionRepositoryAdapter struct {
	store db.Store
}

func NewTransactionRepositoryAdapter(store db.Store) *TransactionRepositoryAdapter {
	return &TransactionRepositoryAdapter{store: store}
}

func (t *TransactionRepositoryAdapter) Save(ctx context.Context, transaction *domain.Transaction) (err error) {
	arg := db.InsertTransactionParams{
		ID:     transaction.ID(),
		GiftID: transaction.GiftID(),
		Amount: transaction.Amount(),
		Date:   transaction.Date(),
		Status: db.TransactionStatus(transaction.Status()),
	}

	_, err = t.store.InsertTransaction(ctx, arg)
	if err != nil {
		fmt.Println("failed to save transaction to db", err)
		return err
	}

	return err
}
