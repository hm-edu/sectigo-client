package domain

type ValidationStatus string

const (
	Validated    ValidationStatus = "VALIDATED"
	NotValidated ValidationStatus = "NOT_VALIDATED"
	Expired      ValidationStatus = "EXPIRED"
)
