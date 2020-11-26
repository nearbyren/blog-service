package model

type Model struct {
	ID         uint32
	CreateBy   string
	ModifiedBy string
	CreatedOn  uint32
	ModifiedOn uint32
	DeletedOn  uint32
	ISDel      uint32
}
