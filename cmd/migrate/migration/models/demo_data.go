package models

type DemoData struct {
	Model

	Name      string `json:"name,omitempty" gorm:"not null;index;type:varchar(255)" valid:"name"`
	Comments  string `json:"comments,omitempty" gorm:"comments,not null;" valid:"comments"`
	OtherInfo string `json:"other_info,omitempty" gorm:"other_info,not null;" valid:"other_info"`

	ControlBy
	ModelTime
}

func (DemoData) TableName() string {
	return "demo_data"
}
