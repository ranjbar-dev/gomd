package types

// Permission holds information about a user permission
type Permission struct {
	// ID
	ID int32 `json:"i"`
	// FirstName
	FirstName string `json:"fn"`
	// LastName
	LastName string `json:"ln"`
	// Email
	Email string `json:"e"`
	// Balance
	Balance float64 `json:"b"`
	// Permissions
	Permissions []string `json:"p"`
	// Status
	Status bool `json:"s"`
}
