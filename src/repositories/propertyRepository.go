package repositories

import "condoapp/src/models"

func GetProperties() []models.Property {
	properties := []models.Property{
		{PropertyId: 1, PropertyAddressName: "LEO ADDRESS"},
	}
	return properties
}