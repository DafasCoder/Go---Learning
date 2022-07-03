package main

import ("testing"
		)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got , want)
		}
	}

	assetError := func (t testing.TB, got, want error){
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but wanted one")
		}
	
		if got != want  {
			t.Errorf("got %s but want %s", got , want)
		}

	}

	t.Run("deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		want := Bitcoin(20)
		assertBalance(t,wallet, want)
	
	
	})
	t.Run("withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(30)}
        wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(20)
		assertBalance(t,wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		
		assetError(t,err, ErrInsufficientFunds)
		assertBalance(t,wallet, startingBalance)
	})

}