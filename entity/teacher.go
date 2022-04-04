package entity

type TeacherEntity struct {
	TeacherId   int    `gorm:"column:teacher;primaryKey;autoIncrement"`
	Username    string `gorm:"column:username;index:username_unique,unique"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	Gender      string `gorm:"column:gender"`
	Certificate string `gorm:"column:certificate"`
}

func (TeacherEntity) TableName() string {
	return "teacher"
}
