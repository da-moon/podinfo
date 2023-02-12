package encoding

import (
	"context"
	"unicode/utf8"

	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

// NewTransactional is the transactional version of the error
// injector
type NewTransactional struct {
	*StorageEncoding
	Transactional physical.Transactional
}

// Transaction is used to run a transactional (atomic) set of operations
func (e *NewTransactional) Transaction(ctx context.Context, txns []*physical.TxnEntry) []physical.ErrTransactionalOperations {
	retErr := make([]physical.ErrTransactionalOperations, 0)

	for _, txn := range txns {
		if !utf8.ValidString(txn.Entry.Key) {
			retErr = append(retErr, physical.ErrTransactionalOperation(*txn, physical.NonUTF8ErrorMessage)) //nolint:staticcheck // this is a complex function and I don't want to mess anything up by modifying this line
			continue
		}
		if e.containsNonPrintableChars(txn.Entry.Key) {
			retErr = append(retErr, physical.ErrTransactionalOperation(*txn, physical.NonPrintableErrorMessage)) //nolint:staticcheck // this is a complex function and I don't want to mess anything up by modifying this line
			continue
		}
	}
	return e.Transactional.Transaction(ctx, txns)
}
