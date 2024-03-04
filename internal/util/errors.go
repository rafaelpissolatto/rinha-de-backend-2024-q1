package util

const (
	simpleError                                        = "an error occurred"
	debitTransactionInconsistentBalanceError           = "debit transaction inconsistent balance"
	creditTransactionInconsistentBalanceError          = "credit transaction inconsistent balance"
	limitMustBeGreaterOrEqualToZeroError               = "limit must be greater than or equal to zero"
	balanceMustBeGreaterOrEqualToTheNegativeLimitError = "balance must be greater than or equal to the negative limit"
	amountMustBeGreaterOrEqualToZeroError              = "amount must be greater than or equal to zero"
	descriptionCannotBeEmptyError                      = "description cannot be empty"
	customerIDCannotBeEmptyError                       = "customer id cannot be empty"
	invalidOperatorTypeError                           = "invalid operator type"
)

// Error returns the error message
func transactionError(err string) string {
	switch err {
	case "debitTransactionInconsistentBalance":
		return debitTransactionInconsistentBalanceError
	case "creditTransactionInconsistentBalance":
		return creditTransactionInconsistentBalanceError
	case "limitMustBeGreaterOrEqualToZero":
		return limitMustBeGreaterOrEqualToZeroError
	case "balanceMustBeGreaterOrEqualToTheNegativeLimit":
		return balanceMustBeGreaterOrEqualToTheNegativeLimitError
	case "amountMustBeGreaterOrEqualToZero":
		return amountMustBeGreaterOrEqualToZeroError
	case "descriptionCannotBeEmpty":
		return descriptionCannotBeEmptyError
	case "customerIDCannotBeEmpty":
		return customerIDCannotBeEmptyError
	case "invalidOperatorType":
		return invalidOperatorTypeError
	default:
		return simpleError
	}
}
