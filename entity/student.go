package entity

type StudentEntity struct {
	StudentId int    `gorm:"column:student_id;primaryKey;autoIncrement"`
	Username  string `gorm:"column:username;index:username_unique,unique"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Gender    string `gorm:"column:gender"`
}

func (StudentEntity) TableName() string {
	return "student"
}
