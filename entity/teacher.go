package entity

type TeacherEntity struct {
	Id          int    `gorm:"column:teacher;primaryKey;autoIncrement"`
	Username    string `gorm:"column:username;index:unique"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	Gender      string `gorm:"column:gender"`
	Certificate string `gorm:"column:certificate"`
}

func (TeacherEntity) TableName() string {
	return "teacher"
}
