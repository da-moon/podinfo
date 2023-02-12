package latency

import (
	"context"
	"time"

	"github.com/da-moon/northern-labs-interview/internal/logger"
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

// TransactionalLatencyInjector is the transactional version of the latency
// injector
type TransactionalLatencyInjector struct {
	*Injector
	Transactional physical.Transactional
}

// NewTransactional creates a new transactional Injector
func NewTransactional(b physical.Backend, latency time.Duration, jitter int, log *logger.WrappedLogger) *TransactionalLatencyInjector {
	return &TransactionalLatencyInjector{
		Injector:      New(b, latency, jitter, log),
		Transactional: b.(physical.Transactional),
	}
}

// Transaction is a latent transaction request
func (l *TransactionalLatencyInjector) Transaction(ctx context.Context, txns []*physical.TxnEntry) []physical.ErrTransactionalOperations {
	l.addLatency()
	return l.Transactional.Transaction(ctx, txns)
}
