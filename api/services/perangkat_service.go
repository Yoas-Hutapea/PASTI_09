package services

import (
	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/repositories"
)

type PerangkatDesaService struct {
	PerangkatDesaRepository *repositories.PerangkatDesaRepository
}

func NewPerangkatDesaService(perangkatRepository *repositories.PerangkatDesaRepository) *PerangkatDesaService {
	return &PerangkatDesaService{
		PerangkatDesaRepository: perangkatRepository,
	}
}

func (ps *PerangkatDesaService) AddPerangkatDesa(perangkat *models.PerangkatDesa) error {
	// Implement the logic to add a perangkat desa to the database
	// Use PerangkatDesaRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PerangkatDesaRepository.AddPerangkatDesa(perangkat)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PerangkatDesaService) UpdatePerangkatDesa(perangkat *models.PerangkatDesa) error {
	// Implement the logic to update a perangkat desa in the database
	// Use PerangkatDesaRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PerangkatDesaRepository.UpdatePerangkatDesa(perangkat)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PerangkatDesaService) DeletePerangkatDesa(perangkatID int) error {
	// Implement the logic to delete a perangkat desa from the database
	// Use PerangkatDesaRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PerangkatDesaRepository.DeletePerangkatDesa(perangkatID)
	if err != nil {
		return err
	}

	return nil
}