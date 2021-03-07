// Package gke implements helper methods to create and list GKE cluster
package gke

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate validates the CreateClusterRequest model and return error if the validation fails
// Returns error if validation fails
func (val CreateClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		validation.Field(&val.Name, validation.Required),
		validation.Field(&val.Description, validation.Required),
	)
}
