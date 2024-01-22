package usecases

import (
	"github.com/lu-css/metrics/src/application/features/metrics/entities"
	repo "github.com/lu-css/metrics/src/infra/database/features/metrics"
)

func Create(name string) (entities.Metric, error) {
	metric := entities.Metric{
		Name: name,
	}

	return repo.Save(metric)
}
