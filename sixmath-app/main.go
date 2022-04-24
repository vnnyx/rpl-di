package main

import (
	"rpl-sixmath/config"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
)

func main() {
	configuration := config.New(".env")
	mysqlMaster := config.MysqlConnection(configuration)
	err := mysqlMaster.AutoMigrate(
		entity.UserEntity{},
		entity.MessageEntity{},
		entity.TransactionEntity{},
		entity.PlaylistEntity{},
		entity.VideoEntity{},
		entity.ExamEntity{},
		entity.QuestionEntity{},
		entity.AnswerEntity{},
	)
	exception.PanicIfNeeded(err)
}
