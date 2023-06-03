package services

import (
	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/repositories"
)

type PendudukService struct {
	UserRepository *repositories.UserRepository
}

func (ps *PendudukService) AddUser(user *models.User) error {
	// Implement the logic to add a user to the database
	// Use UserRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.UserRepository.AddUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PendudukService) UpdateUser(user *models.User) error {
	// Implement the logic to update a user in the database
	// Use UserRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.UserRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PendudukService) DeleteUser(userID int) error {
	// Implement the logic to delete a user from the database
	// Use UserRepository to execute SQL queries
	// Return an error if the operation fails
	err := ps.UserRepository.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}