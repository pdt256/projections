package bank_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/projections/go/bank"
)

func TestProjections(t *testing.T) {
	// Given
	count := bank.NewCount()
	store := bank.NewEventStore(
		count,
	)

	// When
	store.Replay("../../data/events.json")

	// Then
	// Count the number of events
	assert.Equal(t, 50223, count.EventCount)

	// Count the number of deposits and withdrawals

	// Get the total balance for the bank across all accounts

	// Get the total running balance for the bank for each month

	// Find the top 5 accounts with the highest balance. Include account name and id.

	// Find the top 5 accounts with the highest balance by month.
}
