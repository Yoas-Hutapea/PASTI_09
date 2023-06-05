package services

import (
	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/models"
	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/repositories"
)

type PengumumanService struct {
	PengumumanRepository *repositories.PengumumanRepository
}

func NewPengumumanService(pengumumanRepository *repositories.PengumumanRepository) *PengumumanService {
	return &PengumumanService{
		PengumumanRepository: pengumumanRepository,
	}
}

func (ps *PengumumanService) AddPengumuman(pengumuman *models.Pengumuman) error {
	// Implement the logic to add a pengumuman desa to the database
	// Use PengumumanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PengumumanRepository.AddPengumuman(pengumuman)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PengumumanService) UpdatePengumuman(pengumuman *models.Pengumuman) error {
	// Implement the logic to update a pengumuman desa in the database
	// Use PengumumanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PengumumanRepository.UpdatePengumuman(pengumuman)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PengumumanService) DeletePengumuman(perangkatID int) error {
	// Implement the logic to delete a pengumuman desa from the database
	// Use PengumumanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.PengumumanRepository.DeletePengumuman(perangkatID)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PengumumanService) GetPengumumanByID(kegiatanID int) (*models.Pengumuman, error) {
	return ps.PengumumanRepository.GetPengumumanByID(kegiatanID)
}

func (ps *PengumumanService) GetAllPengumuman() ([]*models.Pengumuman, error) {
	return ps.PengumumanRepository.GetAllPengumuman()
}