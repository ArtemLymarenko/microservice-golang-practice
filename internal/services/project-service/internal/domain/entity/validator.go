package entity

// Validator interface shows how concrete implementation must be used to validate the entity.
type Validator interface {
	// Struct method receives any struct and return an error if the validation failed.
	Struct(interface{}) error
}
