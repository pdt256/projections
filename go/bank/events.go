package bank

type AccountWasOpened struct {
	Base
	AccountId string
	FirstName string
	LastName  string
}
type MoneyWasDeposited struct {
	Base
	AccountId  string
	Amount     int
	NewBalance int
}
type DepositFailed struct {
	Base
	AccountId string
	Amount    int
}
type MoneyWasWithdrawn struct {
	Base
	AccountId  string
	Amount     int
	NewBalance int
}
type WithdrawDenied struct {
	Base
	AccountId      string
	Amount         int
	CurrentBalance int
}
type AccountWasClosed struct {
	Base
	AccountId string
}
type FailedToCloseAccountWithBalance struct {
	Base
	AccountId string
	Balance   int
}
