package models

type Property struct {
	TenantId            int
	ManagementGroupId   int
	CondominiumId       int
	PropertyId          int
	PropertyAddressName string
	Residents           []Resident
}
