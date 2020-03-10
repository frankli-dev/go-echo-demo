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
	// FIXME: not sorting properly
	sort.SliceStable(res[:], func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02T15:04:05", res[i].CalendarDate)
		ji, _ := time.Parse("2006-01-02T15:04:05", res[j].CalendarDate)

		return ti.Before(ji)
	})

	return res, nil
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
