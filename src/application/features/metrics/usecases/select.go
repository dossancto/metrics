package usecases

import (
	. "github.com/lu-css/metrics/src/application/features/metrics/entities"
	repository "github.com/lu-css/metrics/src/infra/database/features/metrics"
)

func GetAll() ([]Metric, error) {
	repo := repository.NewMetricRepository()

	m, err := repo.FindAll()

	return m, err
}
