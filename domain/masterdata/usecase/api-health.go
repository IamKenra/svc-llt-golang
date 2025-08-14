package usecase

func (masterdata masterdataUsecase) HealthCheck() error {
	return masterdata.repository.HealthCheck()
}