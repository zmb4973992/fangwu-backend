package model

type ArchivedSeekHouse struct {
	Archive
	SeekHouse
}

func (a ArchivedSeekHouse) TableName() string {
	return "archived_seek_house"
}
