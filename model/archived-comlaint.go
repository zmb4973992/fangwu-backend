package model

type ArchivedComplaint struct {
	Archive
	Complaint
}

func (a ArchivedComplaint) TableName() string {
	return "archived_complaint"
}
