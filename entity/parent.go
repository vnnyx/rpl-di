package entity

type ParentEntity struct {
	ParentId int            `gorm:"column:parent_id;autoIncrement;primaryKey;type:int"`
	Username string         `gorm:"column:username;index:username_unique,unique;type:varchar(50)"`
	Email    string         `gorm:"column:email;type:varchar(100)"`
	Password string         `gorm:"column:password;type:varchar(50)"`
	Gender   string         `gorm:"column:gender;type:varchar(10)"`
	Child    *StudentEntity `gorm:"association_foreignKey:StudentId"`
	ChildId  int            `gorm:"column:child_id;type:int"`
	Avatar   string         `gorm:"column:avatar;type:mediumtext"`
}

func (ParentEntity) TableName() string {
	return "parent"
}
