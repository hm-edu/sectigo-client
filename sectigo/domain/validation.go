package domain

// ValidationStatus represents the staus of the domain validation.
type ValidationStatus string

const (
	// Validated means the DCV was completed successfully.
	Validated ValidationStatus = "VALIDATED"
	// NotValidated means the DCV was either not completed successfully or not initiated.
	NotValidated ValidationStatus = "NOT_VALIDATED"
	// Expired means the DCV is expired without any renewal.
	Expired ValidationStatus = "EXPIRED"
)
