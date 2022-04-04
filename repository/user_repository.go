package repository

import (
	"context"
	"database/sql"
	"rpl-sixmath/entity"
)

type UserRepository interface {
	InsertTeacher(ctx context.Context, tx *sql.Tx, teacher entity.TeacherEntity) entity.TeacherEntity
	InsertStudent(ctx context.Context, tx *sql.Tx, student entity.StudentEntity) entity.StudentEntity
	InsertParent(ctx context.Context, tx *sql.Tx, parent entity.ParentEntity) entity.ParentEntity
	UpdateTeacher(ctx context.Context, tx *sql.Tx, teacher entity.TeacherEntity) entity.TeacherEntity
	UpdateStudent(ctx context.Context, tx *sql.Tx, student entity.StudentEntity) entity.StudentEntity
	UpdateParent(ctx context.Context, tx *sql.Tx, parent entity.ParentEntity) entity.ParentEntity
	DeleteTeacher(ctx context.Context, tx *sql.Tx, teacherId int)
	DeleteStudent(ctx context.Context, tx *sql.Tx, studentId int)
	DeleteParent(ctx context.Context, tx *sql.Tx, parentId int)
	FindTeacherById(ctx context.Context, tx *sql.Tx, teacherId int) (entity.TeacherEntity, error)
	FindStudentById(ctx context.Context, tx *sql.Tx, studentId int) (entity.StudentEntity, error)
	FindParentById(ctx context.Context, tx *sql.Tx, parentId int) (entity.ParentEntity, error)
	FindTeacherAll(ctx context.Context, tx *sql.Tx) []entity.TeacherEntity
	FindStudentAll(ctx context.Context, tx *sql.Tx) []entity.StudentEntity
	FindParentAll(ctx context.Context, tx *sql.Tx) []entity.ParentEntity
}
