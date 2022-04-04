package entity

type ParentEntity struct {
	Id       int
	Username string
	Email    string
	Password string
	Gender   string
	ChildId  StudentEntity
}

type StudentEntity struct {
	Id       int
	Username string
	Email    string
	Password string
	Gender   string
}

type TeacherEntity struct {
	Id          int
	Username    string
	Email       string
	Password    string
	Gender      string
	Certificate string
}

type AdminEntity struct {
	Id       int
	Username string
	Email    string
	Password string
}
