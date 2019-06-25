package entity

type Codetype struct {
	Codetype    string `xorm:"varchar(50) not null 'code_type'" json:"code_type"`
	Description string `xorm:"varchar(255) null 'description'" json:"description"` //  "说明"
}

func (self *Codetype) TableName() string {
	return "kiki_code_type"
}
