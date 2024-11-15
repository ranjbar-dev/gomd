package types

// User holds information about a user
type User struct {
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
	// Usernames
	Usernames []string `json:"u"`
	// Maps
	Maps map[int32]string `json:"m"`
	// Permissions
	Permissions Permission `json:"p"`
	// Status
	Status bool `json:"s"`
}
