package simple

type SimpleRepository struct {
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	return &SimpleService{
		SimpleRepository: repository,
	}
}