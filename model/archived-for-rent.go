package model

type ArchivedForRent struct {
	Archive
	ForRent
}

func (a ArchivedForRent) TableName() string {
	return "archived_for_rent"
}
