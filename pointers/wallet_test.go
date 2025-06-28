package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(25)}
		err := wallet.WithDraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.WithDraw(Bitcoin(50))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()

	result := wallet.Balance()

	if result != expected {
		t.Errorf("expected %s got %s", expected, result)
	}
}

func assertError(t testing.TB, result error, expected error) {
	t.Helper()

	if result == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func assertNoError(t testing.TB, result error) {
	t.Helper()
	if result != nil {
		t.Fatal("got an error but didn't expect one")
	}
}
