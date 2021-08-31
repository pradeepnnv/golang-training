package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Check Bitcoin string representation", func(t *testing.T) {
		b := Bitcoin(5)
		got := b.String()
		want := fmt.Sprintf("%d BTC", int(b))
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("Deposit Bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw Bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(5)
		assertNoError(t, err)
		want := Bitcoin(15)

		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(25)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
