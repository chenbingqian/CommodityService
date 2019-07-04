package entity

type Attribute struct {
	AttributeId         int64  `xorm:"int(10) unsigned  not null   'attribute_id' " json:"attribute_id"`
	AttributeName       string `xorm:"varchar(100)  not null  'attribute_name' " json:"attribute_name"`
	AttributeTypeId     int64  `xorm:"int(10)  not null  'attribute_type_id' " json:"attribute_type_id"`
	AttributeTypeName   string `xorm:"varchar(50)  not null  'attribute_type_name' " json:"attribute_type_name"`
	AttributeValue      string `xorm:"text  'attribute_value' " json:"attribute_value"`
	AttributeValueStyle string `xorm:"text  'attribute_value_style' " json:"attribute_value_style"`
	IsCustom            string `xorm:"varchar(2) 'is_custom' " json:"is_custom"`
	AddTime             int64  `xorm:"int(10)  not null  'add_time' " json:"add_time"`
	UpdateTime          int64  `xorm:"int(10)  not null  'update_time' " json:"update_time"`
}

func (self *Attribute) TableName() string {
	return "kiki_attribute"
}
