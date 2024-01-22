package usecases

import "github.com/lu-css/metrics/src/application/features/metrics/entities"

func Create(name string) entities.Metric {
	return entities.Metric{
		Name: name,
	}
}
