package services

import (
	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/repositories"
)

type PengumumanService struct {
	PengumumanRepository *repositories.PengumumanRepository
}

func NewPengumumanService(pengumumanRepository *repositories.PengumumanRepository) *PengumumanService {
	return &PengumumanService{
		PengumumanRepository: pengumumanRepository,
	}
}

func (ps *PengumumanService) AddPengumuman(perangkat *models.Pengumuman) error {
	// Implement the logic to add a perangkat desa to the database
	// Use PengumumanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PengumumanRepository.AddPengumuman(perangkat)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PengumumanService) UpdatePengumuman(perangkat *models.Pengumuman) error {
	// Implement the logic to update a perangkat desa in the database
	// Use PengumumanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PengumumanRepository.UpdatePengumuman(perangkat)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PengumumanService) DeletePengumuman(perangkatID int) error {
	// Implement the logic to delete a perangkat desa from the database
	// Use PengumumanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PengumumanRepository.DeletePengumuman(perangkatID)
	if err != nil {
		return err
	}

	return nil
}