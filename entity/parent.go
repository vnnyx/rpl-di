package entity

type ParentEntity struct {
	ParentId int            `gorm:"column:parent_id;autoIncrement;primaryKey"`
	Username string         `gorm:"column:username;index:username_unique,unique"`
	Email    string         `gorm:"column:email"`
	Password string         `gorm:"column:password"`
	Gender   string         `gorm:"column:gender"`
	Child    *StudentEntity `gorm:"association_foreignKey:StudentId"`
	ChildId  int            `gorm:"column:child_id"`
}

func (ParentEntity) TableName() string {
	return "parent"
}
