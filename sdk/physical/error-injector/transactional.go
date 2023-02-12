package errorinjector

import (
	"context"

	"github.com/da-moon/northern-labs-interview/internal/logger"
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

// TransactionalErrorInjector is the transactional version of the error
// injector
type TransactionalErrorInjector struct {
	*ErrorInjector
	Transactional physical.Transactional
}

// NewTransactional creates a new transactional ErrorInjector
func NewTransactional(b physical.Backend, errorPercent int, log *logger.WrappedLogger) *TransactionalErrorInjector {
	return &TransactionalErrorInjector{
		ErrorInjector: New(b, errorPercent, log),
		Transactional: b.(physical.Transactional),
	}
}

// Transaction runs a transactional (atomic) error injector
func (e *TransactionalErrorInjector) Transaction(ctx context.Context, txns []*physical.TxnEntry) []physical.ErrTransactionalOperations {
	retErr := make([]physical.ErrTransactionalOperations, 0)
	actualTxn := make([]*physical.TxnEntry, 0)
	for _, txn := range txns {
		err := e.addError()
		if err != nil {
			retErr = append(retErr, physical.ErrTransactionalOperation(*txn, err.Error()))
			continue
		}
		actualTxn = append(actualTxn, txn)
	}
	retErr = append(retErr, e.Transactional.Transaction(ctx, actualTxn)...)
	return retErr
}
