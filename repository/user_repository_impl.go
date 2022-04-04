package repository

import (
	"context"
	"database/sql"
	"rpl-sixmath/model/entity"
)

type UserRepositoryImpl struct {
}

func (repo *UserRepositoryImpl) InsertTeacher(ctx context.Context, tx *sql.Tx, teacher entity.TeacherEntity) entity.TeacherEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) InsertStudent(ctx context.Context, tx *sql.Tx, student entity.StudentEntity) entity.StudentEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) InsertParent(ctx context.Context, tx *sql.Tx, parent entity.ParentEntity) entity.ParentEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) UpdateTeacher(ctx context.Context, tx *sql.Tx, teacher entity.TeacherEntity) entity.TeacherEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) UpdateStudent(ctx context.Context, tx *sql.Tx, student entity.StudentEntity) entity.StudentEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) UpdateParent(ctx context.Context, tx *sql.Tx, parent entity.ParentEntity) entity.ParentEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) DeleteTeacher(ctx context.Context, tx *sql.Tx, teacherId int) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) DeleteStudent(ctx context.Context, tx *sql.Tx, studentId int) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) DeleteParent(ctx context.Context, tx *sql.Tx, parentId int) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) FindTeacherById(ctx context.Context, tx *sql.Tx, teacherId int) (entity.TeacherEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) FindStudentById(ctx context.Context, tx *sql.Tx, studentId int) (entity.StudentEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) FindParentById(ctx context.Context, tx *sql.Tx, parentId int) (entity.ParentEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) FindTeacherAll(ctx context.Context, tx *sql.Tx) []entity.TeacherEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) FindStudentAll(ctx context.Context, tx *sql.Tx) []entity.StudentEntity {
	//TODO implement me
	panic("implement me")
}

func (repo *UserRepositoryImpl) FindParentAll(ctx context.Context, tx *sql.Tx) []entity.ParentEntity {
	//TODO implement me
	panic("implement me")
}
