package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	result := wallet.Balance()
	expected := Bitcoin(10)

	if result != expected {
		t.Errorf("expected %s got %s", expected, result)
	}
}
