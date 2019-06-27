package entity

type Code struct {
	Codetype    string `xorm:"varchar(50) not null  'code_type' " json:"code_type"`
	Name        string `xorm:"varchar(100) not null 'name' " json:"name"`
	Value       string `xorm:"varchar(100) not null 'value'" json:"value"`
	Description string `xorm:"varchar(255) 'description' " json:"description"`
}

func (self *Code) TableName() string {
	return "kiki_code"
}
