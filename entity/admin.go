package entity

type AdminEntity struct {
	Id       int    `gorm:"column:admin_id;primaryKey;autoIncrement"`
	Username string `gorm:"column:username;index:unique"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
