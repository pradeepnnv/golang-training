package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("want error. got none")
		}
	}

	t.Run("Deposit Bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw Bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		wallet.Withdraw(5)
		want := Bitcoin(15)

		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(25)
		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
