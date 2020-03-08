package listing

import (
	"fmt"
	"sort"
)

// Repository interface
type Repository interface {
	GetAllRenewableResources() ([]*RenewableResource, error) // get all renewable resources
}

// Service interface
type Service interface {
	GetRenewableResources() ([]*RenewableResource, error) // get all renewable resources
}

// service struct is a Service implementation
type service struct {
	r Repository
}

// GetRenewableResources retrieves all renewable resources through
// the service repository passed to the constructor argument
func (s *service) GetRenewableResources() ([]*RenewableResource, error) {
	res, err := s.r.GetAllRenewableResources()
	if err != nil {
		return nil, fmt.Errorf("Error while retrieving data from repository (GetAllRenewableResources):%v", err)
	}
	sort.Slice(res[:], func(i, j int) bool {
		return res[i].CalendarDate < res[j].CalendarDate
	})

	return res, nil
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
