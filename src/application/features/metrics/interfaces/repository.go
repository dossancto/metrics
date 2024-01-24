package interfaces

import "github.com/lu-css/metrics/src/application/features/metrics/entities"

type MetricRepository interface {
	Create(metric entities.Metric) error
	Get(id string) (*entities.Metric, error)
	Delete(id string) error
	List(limit, offset uint) ([]entities.Metric, error)
}
