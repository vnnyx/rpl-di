package main

import (
	"rpl-sixmath/config"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
)

func main() {
	configuration := config.New(".env")
	mysqlMaster := config.MysqlConnection(configuration)
	err := mysqlMaster.AutoMigrate(entity.ExamEntity{}, entity.VideoEntity{}, entity.ParentEntity{},
		entity.CodeMessageEntity{}, entity.StudentEntity{}, entity.MessageEntity{}, entity.TeacherEntity{},
		entity.PlaylistEntity{}, entity.AdminEntity{}, entity.AnswerEntity{}, entity.Transaction{})
	exception.PanicIfNeeded(err)
}
