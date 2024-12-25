package model

type ArchivedComment struct {
	Archive
	Comment
}

func (a ArchivedComment) TableName() string {
	return "archived_comment"
}
