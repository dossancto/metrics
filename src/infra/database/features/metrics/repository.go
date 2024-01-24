package metrics

import (
	"github.com/lu-css/metrics/src/application/features/metrics/entities"
	gonanoid "github.com/matoous/go-nanoid"
	"gorm.io/gorm"
)

type GormMetricRepository struct {
	db *gorm.DB
}

func NewMetricRepository(db *gorm.DB) *GormMetricRepository { return &GormMetricRepository{db}
}

func (repo *MetricRepository) FindAll() ([]entities.Metric, error) {
var metrics []entities.Metric
	if err := repo.db.Find(&metrics).Error; err != nil {
		return nil, err
	}
	return metrics, nil
}

func (repo *MetricRepository) FindById(id uint) (*entities.Metric, error) {
	var metric entities.Metric
	if err := repo.db.First(&metric, id).Error; err != nil {
		return nil, err
	}
	return &metric, nil
}

func (repo *MetricRepository) Save(metric entities.Metric) error {
	return repo.db.Save(&metric).Error
}

func (repo *MetricRepository) Delete(id uint) error {
	return repo.db.Delete(&entities.Metric{}, id).Error
}

func Save(entity entities.Metric) (entities.Metric, error) {
	id, err := gonanoid.Nanoid(38)

	if err != nil {
		return entities.Metric{}, err
	}

	entity.Id = id

	return entity, nil
}
