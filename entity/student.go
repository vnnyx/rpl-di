package entity

type StudentEntity struct {
	StudentId int    `gorm:"column:student_id;primaryKey;autoIncrement;type:int"`
	Username  string `gorm:"column:username;unique;type:varchar(50)"`
	Email     string `gorm:"column:email;type:varchar(100)"`
	Password  string `gorm:"column:password;type:varchar(50)"`
	Gender    string `gorm:"column:gender;type:varchar(10)"`
	Avatar    string `gorm:"column:avatar;type:mediumtext"`
}

func (StudentEntity) TableName() string {
	return "student"
}
