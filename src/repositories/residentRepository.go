package repositories

import "condoapp/src/models"

func GetResidents() []models.Resident {
	residents := []models.Resident{
		{ResidentId: 1, PropertyId: 1, ResidentFirstName: "LEO RESIDENT", ResidentType: 1},
	}
	return residents
}
