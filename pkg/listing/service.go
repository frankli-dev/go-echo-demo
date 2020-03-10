package listing

import (
	"fmt"
	"sort"
	"time"
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
// the repository passed as an argument to NewService
func (s *service) GetRenewableResources() ([]*RenewableResource, error) {
	res, err := s.r.GetAllRenewableResources()
	if err != nil {
		return nil, fmt.Errorf("Error while retrieving data from repository (GetAllRenewableResources):%v", err)
	}

	sort.SliceStable(res, func(i, j int) bool {
		ti, err := time.Parse("01/02/2006", res[i].CalendarDate)
		tj, err := time.Parse("01/02/2006", res[j].CalendarDate)

		if err != nil {
			panic(err)
		}

		return ti.Before(tj)
	})

	return res, nil
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
