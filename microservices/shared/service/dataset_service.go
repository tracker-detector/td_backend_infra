package service

import (
	"context"
	"tds/shared/models"
	"tds/shared/payload"
)

type IDatasetService interface {
	Save(ctx context.Context, dataset *models.Dataset) (*models.Dataset, error)
	CreateDataset(ctx context.Context, datasetPayload *payload.CreateDatasetPayload) (*models.Dataset, error)
	SaveAll(ctx context.Context, datasets []*models.Dataset) ([]*models.Dataset, error)
	GetAllDatasets() []*models.Dataset
	ReloadCache(ctx context.Context)
	IsValidDataset(ctx context.Context, id string) bool
	IsLabelValid(label string) bool
}

type DatasetService struct {
	datasetRepo  models.DatasetRepository
	datasetCache []*models.Dataset
}

func NewDatasetService(datasetRepo models.DatasetRepository) *DatasetService {
	service := &DatasetService{
		datasetRepo: datasetRepo,
	}
	service.ReloadCache(context.Background())
	return service
}

func (s *DatasetService) SaveAll(ctx context.Context, datasets []*models.Dataset) ([]*models.Dataset, error) {
	res, err := s.datasetRepo.SaveAll(ctx, datasets)
	s.ReloadCache(ctx)
	return res, err
}

func (s *DatasetService) CreateDataset(ctx context.Context, datasetPayload *payload.CreateDatasetPayload) (*models.Dataset, error) {
	dataset := &models.Dataset{
		Name:        datasetPayload.Name,
		Description: datasetPayload.Description,
		Label:       datasetPayload.Label,
	}
	return s.Save(ctx, dataset)
}

func (s *DatasetService) Save(ctx context.Context, dataset *models.Dataset) (*models.Dataset, error) {
	res, err := s.datasetRepo.Save(ctx, dataset)
	s.ReloadCache(ctx)
	return res, err
}

func (s *DatasetService) GetAllDatasets() []*models.Dataset {
	return s.datasetCache
}

func (s *DatasetService) IsValidDataset(ctx context.Context, id string) bool {
	dataset, err := s.datasetRepo.FindByID(ctx, id)
	if err != nil {
		return false
	}
	return dataset != nil
}

func (s *DatasetService) ReloadCache(ctx context.Context) {
	datasets, _ := s.datasetRepo.FindAll(ctx)
	s.datasetCache = datasets
}

func (s *DatasetService) IsLabelValid(label string) bool {
	for _, dataset := range s.datasetCache {
		if dataset.Label == label {
			return true
		}
	}
	return false
}
