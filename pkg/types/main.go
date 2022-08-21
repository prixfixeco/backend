package types

import (
	"fmt"
)

const (
	// sortAscendingString is the pre-determined Ascending sortType for external use.
	sortAscendingString = "asc"
	// sortDescendingString is the pre-determined Descending sortType for external use.
	sortDescendingString = "desc"
)

var (
	// SortAscending is the pre-determined Ascending string for external use.
	SortAscending = func(x string) *string { return &x }(sortAscendingString)
	// SortDescending is the pre-determined Descending string for external use.
	SortDescending = func(x string) *string { return &x }(sortDescendingString)
)

type (
	// ContextKey represents strings to be used in Context objects. From the docs:
	// 	"The provided key must be comparable and should not be of type string or
	// 	 any other built-in type to avoid collisions between packages using context."
	ContextKey string

	// Pagination represents a pagination request.
	Pagination struct {
		_ struct{}

		Page          uint64 `json:"page"`
		Limit         uint8  `json:"limit"`
		FilteredCount uint64 `json:"filteredCount"`
		TotalCount    uint64 `json:"totalCount"`
	}

	// ErrorResponse represents a response we might send to the User in the event of an error.
	ErrorResponse struct {
		_ struct{}

		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	// PreWriteResponse is what we respond with when the data requested to be written has not yet been written.
	PreWriteResponse struct {
		_ struct{}

		ID string `json:"id"`
	}
)

var _ error = (*ErrorResponse)(nil)

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", er.Code, er.Message)
}
