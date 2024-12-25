package model

type ArchivedUser struct {
	Archive
	User
}

func (a ArchivedUser) TableName() string {
	return "archived_user"
}
