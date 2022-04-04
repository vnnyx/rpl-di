package repository

import "rpl-sixmath/entity"

type UserRepository interface {
	InsertTeacher(teacher entity.TeacherEntity) (entity.TeacherEntity, error)
	InsertStudent(student entity.StudentEntity) (entity.StudentEntity, error)
	InsertParent(parent entity.ParentEntity) (entity.ParentEntity, error)
	UpdateTeacher(teacher entity.TeacherEntity) (entity.TeacherEntity, error)
	UpdateStudent(student entity.StudentEntity) (entity.StudentEntity, error)
	UpdateParent(parent entity.ParentEntity) (entity.ParentEntity, error)
	DeleteTeacher(teacherId int) error
	DeleteStudent(studentId int) error
	DeleteParent(parentId int) error
	FindTeacherById(teacherId int) (response entity.TeacherEntity, err error)
	FindStudentById(studentId int) (responer entity.StudentEntity, err error)
	FindParentById(parentId int) (responer entity.ParentEntity, err error)
	FindTeacherAll() (response []entity.TeacherEntity, err error)
	FindStudentAll() (response []entity.StudentEntity, err error)
	FindParentAll() (response []entity.ParentEntity, err error)
}
