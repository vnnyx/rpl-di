package repository

import (
	"rpl-sixmath/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repo UserRepositoryImpl) InsertTeacher(teacher entity.TeacherEntity) (entity.TeacherEntity, error) {
	//TODO implement me
	err := repo.DB.Create(&teacher).Error
	return teacher, err
}

func (repo UserRepositoryImpl) InsertStudent(student entity.StudentEntity) (entity.StudentEntity, error) {
	//TODO implement me
	err := repo.DB.Create(&student).Error
	return student, err
}

func (repo UserRepositoryImpl) InsertParent(parent entity.ParentEntity) (entity.ParentEntity, error) {
	//TODO implement me
	err := repo.DB.Create(&parent).Error
	return parent, err
}

func (repo UserRepositoryImpl) UpdateTeacher(teacher entity.TeacherEntity) (entity.TeacherEntity, error) {
	//TODO implement me
	err := repo.DB.Where("teacher_id", teacher.Id).Updates(&teacher).Error
	return teacher, err
}

func (repo UserRepositoryImpl) UpdateStudent(student entity.StudentEntity) (entity.StudentEntity, error) {
	//TODO implement me
	err := repo.DB.Where("teacher_id", student.Id).Updates(&student).Error
	return student, err
}

func (repo UserRepositoryImpl) UpdateParent(parent entity.ParentEntity) (entity.ParentEntity, error) {
	//TODO implement me
	err := repo.DB.Where("teacher_id", parent.Id).Updates(&parent).Error
	return parent, err
}

func (repo UserRepositoryImpl) DeleteTeacher(teacherId int) error {
	//TODO implement me
	err := repo.DB.Where("teacher_id", teacherId).Delete(&entity.TeacherEntity{}).Error
	return err
}

func (repo UserRepositoryImpl) DeleteStudent(studentId int) error {
	//TODO implement me
	err := repo.DB.Where("student_id", studentId).Delete(&entity.StudentEntity{}).Error
	return err
}

func (repo UserRepositoryImpl) DeleteParent(parentId int) error {
	//TODO implement me
	err := repo.DB.Where("parent_id", parentId).Delete(&entity.ParentEntity{}).Error
	return err
}

func (repo UserRepositoryImpl) FindTeacherById(teacherId int) (response entity.TeacherEntity, err error) {
	//TODO implement me
	err = repo.DB.Where("teacher_id", teacherId).First(&response).Error
	return response, err
}

func (repo UserRepositoryImpl) FindStudentById(studentId int) (response entity.StudentEntity, err error) {
	//TODO implement me
	err = repo.DB.Where("student_id", studentId).First(&response).Error
	return response, err
}

func (repo UserRepositoryImpl) FindParentById(parentId int) (response entity.ParentEntity, err error) {
	//TODO implement me
	err = repo.DB.Where("parent_id", parentId).First(&response).Error
	return response, err
}

func (repo UserRepositoryImpl) FindTeacherAll() (response []entity.TeacherEntity, err error) {
	//TODO implement me
	err = repo.DB.Find(&response).Error
	return response, err
}

func (repo UserRepositoryImpl) FindStudentAll() (response []entity.StudentEntity, err error) {
	//TODO implement me
	err = repo.DB.Find(&response).Error
	return response, err
}

func (repo UserRepositoryImpl) FindParentAll() (response []entity.ParentEntity, err error) {
	//TODO implement me
	err = repo.DB.Find(&response).Error
	return response, err
}
