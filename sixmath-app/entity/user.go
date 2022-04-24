package entity

type UserEntity struct {
	UserId          int         `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username        string      `gorm:"column:username;unique;type:varchar(50)"`
	Email           string      `gorm:"column:email;type:varchar(100)"`
	Handphone       string      `gorm:"column:handphone;type:varchar(20)"`
	Password        string      `gorm:"column:password;type:varchar(50)"`
	Role            string      `gorm:"column:role;type:varchar(10)"`
	Gender          string      `gorm:"column:gender;type:varchar(10)"`
	Certificate     string      `gorm:"column:certificate;type:mediumtext"`
	Avatar          string      `gorm:"column:avatar;type:mediumtext"`
	StudentUsername string      `gorm:"column:student_username;type:varchar(50)"`
	Student         *UserEntity `gorm:"references:Username;association_foreignkey:StudentUsername"`
}

func (UserEntity) TableName() string {
	return "user"
}
