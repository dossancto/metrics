package usecases

import (
	. "github.com/lu-css/metrics/src/application/features/metrics/entities"
	repository "github.com/lu-css/metrics/src/infra/database/features/metrics"
)

func Create(name string) (Metric, error) {
	repo := repository.NewMetricRepository()

	metric := Metric{
		Name: name,
	}

	m, err := repo.Save(metric)

	return m, err
}
