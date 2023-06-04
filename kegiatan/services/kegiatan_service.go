package services

import (
	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/models"
	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/repositories"
)

type KegiatanService struct {
	KegiatanRepository *repositories.KegiatanRepository
}

func NewKegiatanService(kegiatanRepository *repositories.KegiatanRepository) *KegiatanService {
	return &KegiatanService{
		KegiatanRepository: kegiatanRepository,
	}
}

func (ks *KegiatanService) AddKegiatan(kegiatan *models.Kegiatan) error {
	// Implement the logic to add a kegiatan to the database
	// Use KegiatanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ks.KegiatanRepository.AddKegiatan(kegiatan)
	if err != nil {
		return err
	}

	return nil
}

func (ks *KegiatanService) UpdateKegiatan(kegiatan *models.Kegiatan) error {
	// Implement the logic to update a kegiatan in the database
	// Use KegiatanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ks.KegiatanRepository.UpdateKegiatan(kegiatan)
	if err != nil {
		return err
	}

	return nil
}

func (ks *KegiatanService) DeleteKegiatan(kegiatanID int) error {
	// Implement the logic to delete a kegiatan from the database
	// Use KegiatanRepository to execute SQL queries
	// Return an error if the operation fails
	err := ks.KegiatanRepository.DeleteKegiatan(kegiatanID)
	if err != nil {
		return err
	}

	return nil
}
