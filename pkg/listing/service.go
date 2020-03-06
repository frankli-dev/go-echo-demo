package listing

// TODO: add comment
type Repository interface {
	GetAllRenewableResources() []*RenewableResource
}

// TODO: add comment
type Service interface {
	GetRenewableResources() []*RenewableResource
}

// TODO: add comment
type service struct {
	r Repository
}

// TODO: add comment
// TODO: add error handling
func (s *service) GetRenewableResources() []*RenewableResource {
	return s.r.GetAllRenewableResources()
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
