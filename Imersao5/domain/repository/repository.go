package repository

type TransactionRepository interface {
	Insert(id string, account string, amount float64, stauts string, errorMessage string) error
}
