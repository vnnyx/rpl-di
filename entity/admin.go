package entity

type AdminEntity struct {
	AdminId  int    `gorm:"column:admin_id;primaryKey;autoIncrement"`
	Username string `gorm:"column:username;index:username_unique,unique"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (AdminEntity) TableName() string {
	return "admin"
}
