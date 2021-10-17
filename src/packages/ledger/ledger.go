/**
BY: Deyana Atanasova, Henrik Tambo Buhl & Alexander Stæhr Johansen
DATE: 16-10-2021
COURSE: Distributed Systems and Security
DESCRIPTION: Distributed transaction system implemented as structured P2P flooding network.
**/

package ledger

import (
	"fmt"
	"strconv"
	"sync"
)

/* Signed transaction struct */
type SignedTransaction struct {
	Type      string // signedTransaction
	ID        string // ID of the transaction
	From      string // Sender of the transaction (public key)
	To        string // Receiver of the transaction (public key)
	Amount    int    // Amount to transfer
	Signature string // Signature of the transaction
}

/* Ledger struct */
type Ledger struct {
	Type     string
	Accounts map[string]int
	lock     sync.Mutex
}

/* Ledger constructor */
func MakeLedger() *Ledger {
	ledger := new(Ledger)
	ledger.Accounts = make(map[string]int)
	return ledger
}

/* Transaction method */
func (ledger *Ledger) Transaction(signedTransaction SignedTransaction) {
	ledger.lock.Lock()
	defer ledger.lock.Unlock()
	ledger.Accounts[signedTransaction.From] -= signedTransaction.Amount
	ledger.Accounts[signedTransaction.To] += signedTransaction.Amount
}

/* Print ledger method */
func (ledger *Ledger) PrintLedger() {
	ledger.lock.Lock()
	for account, amount := range ledger.Accounts {
		fmt.Println("Account name: " + account + " amount: " + strconv.Itoa(amount))
	}
	ledger.lock.Unlock()
}
