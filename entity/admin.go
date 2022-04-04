package entity

type AdminEntity struct {
	AdminId  int    `gorm:"column:admin_id;primaryKey;autoIncrement;type:int"`
	Username string `gorm:"column:username;unique;type:varchar(50)"`
	Email    string `gorm:"column:email;type:varchar(100)"`
	Password string `gorm:"column:password;type:varchar(100)"`
	Gender   string `gorm:"column:gender;type:varchar(10)"`
	Avatar   string `gorm:"column:avatar;type:mediumtext"`
}

func (AdminEntity) TableName() string {
	return "admin"
}
