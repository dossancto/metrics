package metrics

import (
	"github.com/lu-css/metrics/src/application/features/metrics/entities"
	gonanoid "github.com/matoous/go-nanoid"
)

func Save(entity entities.Metric) (entities.Metric, error) {
	id, err := gonanoid.Nanoid(38)

	if err != nil {
		return entities.Metric{}, err
	}

	entity.Id = id

	return entity, nil
}
