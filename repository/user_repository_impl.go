package repository

import (
	"context"
	"database/sql"
	"errors"
	"rpl-sixmath/entity"
	"rpl-sixmath/helper"
)

type UserRepositoryImpl struct {
}

func (repo *UserRepositoryImpl) InsertTeacher(ctx context.Context, tx *sql.Tx, teacher entity.TeacherEntity) entity.TeacherEntity {
	SQL := "insert into teacher(username, email, password, gender, certificate) values(?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, teacher.Username, teacher.Email, teacher.Password, teacher.Gender, teacher.Certificate)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	teacher.Id = int(id)
	return teacher
}

func (repo *UserRepositoryImpl) InsertStudent(ctx context.Context, tx *sql.Tx, student entity.StudentEntity) entity.StudentEntity {
	SQL := "insert into student(username, email, password, gender, certificate) values(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, student.Username, student.Email, student.Password, student.Gender)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	student.Id = int(id)
	return student
}

func (repo *UserRepositoryImpl) InsertParent(ctx context.Context, tx *sql.Tx, parent entity.ParentEntity) entity.ParentEntity {
	SQL := "insert into parent(username, email, password, gender, certificate) values(?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, parent.Username, parent.Email, parent.Password, parent.Gender, parent.ChildId.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	parent.Id = int(id)
	return parent
}

func (repo *UserRepositoryImpl) UpdateTeacher(ctx context.Context, tx *sql.Tx, teacher entity.TeacherEntity) entity.TeacherEntity {
	SQL := "update teacher set username=?, email=?, password=?, gender=?, certificate=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, teacher.Username, teacher.Email, teacher.Password, teacher.Gender, teacher.Certificate, teacher.Id)
	helper.PanicIfError(err)

	return teacher
}

func (repo *UserRepositoryImpl) UpdateStudent(ctx context.Context, tx *sql.Tx, student entity.StudentEntity) entity.StudentEntity {
	SQL := "update student set username=?, email=?, password=?, gender=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, student.Username, student.Email, student.Password, student.Gender, student.Id)
	helper.PanicIfError(err)

	return student
}

func (repo *UserRepositoryImpl) UpdateParent(ctx context.Context, tx *sql.Tx, parent entity.ParentEntity) entity.ParentEntity {
	SQL := "update parent set username=?, email=?, password=?, gender=?, child_id=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, parent.Username, parent.Email, parent.Password, parent.Gender, parent.ChildId.Id, parent.Id)
	helper.PanicIfError(err)

	return parent
}

func (repo *UserRepositoryImpl) DeleteTeacher(ctx context.Context, tx *sql.Tx, teacherId int) {
	SQL := "delete from teacher where id=?"
	_, err := tx.ExecContext(ctx, SQL, teacherId)
	helper.PanicIfError(err)
}

func (repo *UserRepositoryImpl) DeleteStudent(ctx context.Context, tx *sql.Tx, studentId int) {
	SQL := "delete from student where id=?"
	_, err := tx.ExecContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
}

func (repo *UserRepositoryImpl) DeleteParent(ctx context.Context, tx *sql.Tx, parentId int) {
	SQL := "delete from parent where id=?"
	_, err := tx.ExecContext(ctx, SQL, parentId)
	helper.PanicIfError(err)
}

func (repo *UserRepositoryImpl) FindTeacherById(ctx context.Context, tx *sql.Tx, teacherId int) (entity.TeacherEntity, error) {
	SQL := "select * from teacher where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, teacherId)
	helper.PanicIfError(err)
	defer rows.Close()

	teacher := entity.TeacherEntity{}
	if rows.Next() {
		err := rows.Scan(&teacher.Id, &teacher.Username, &teacher.Email, &teacher.Password, &teacher.Gender, &teacher.Certificate)
		helper.PanicIfError(err)
		return teacher, nil
	} else {
		return teacher, errors.New("category not found")
	}
}

func (repo *UserRepositoryImpl) FindStudentById(ctx context.Context, tx *sql.Tx, studentId int) (entity.StudentEntity, error) {
	SQL := "select * from student where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := entity.StudentEntity{}
	if rows.Next() {
		err := rows.Scan(&student.Id, &student.Username, &student.Email, &student.Password, &student.Gender)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("category not found")
	}
}

func (repo *UserRepositoryImpl) FindParentById(ctx context.Context, tx *sql.Tx, parentId int) (entity.ParentEntity, error) {
	SQL := "select * from parent where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, parentId)
	helper.PanicIfError(err)
	defer rows.Close()

	parent := entity.ParentEntity{}
	if rows.Next() {
		err := rows.Scan(&parent.Id, &parent.Username, &parent.Email, &parent.Password, &parent.Gender, &parent.ChildId.Id)
		helper.PanicIfError(err)
		return parent, nil
	} else {
		return parent, errors.New("category not found")
	}
}

func (repo *UserRepositoryImpl) FindTeacherAll(ctx context.Context, tx *sql.Tx) []entity.TeacherEntity {
	SQL := "select * from teacher"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var teachers []entity.TeacherEntity
	for rows.Next() {
		teacher := entity.TeacherEntity{}
		err := rows.Scan(&teacher.Id, &teacher.Username, &teacher.Email, &teacher.Password, &teacher.Gender, &teacher.Certificate)
		helper.PanicIfError(err)
		teachers = append(teachers, teacher)
	}
	return teachers
}

func (repo *UserRepositoryImpl) FindStudentAll(ctx context.Context, tx *sql.Tx) []entity.StudentEntity {
	SQL := "select * from student"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var students []entity.StudentEntity
	for rows.Next() {
		student := entity.StudentEntity{}
		err := rows.Scan(&student.Id, &student.Username, &student.Email, &student.Password, &student.Gender)
		helper.PanicIfError(err)
		students = append(students, student)
	}
	return students
}

func (repo *UserRepositoryImpl) FindParentAll(ctx context.Context, tx *sql.Tx) []entity.ParentEntity {
	SQL := "select * from parent"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var parents []entity.ParentEntity
	for rows.Next() {
		parent := entity.ParentEntity{}
		err := rows.Scan(&parent.Id, &parent.Username, &parent.Email, &parent.Password, &parent.Gender, &parent.ChildId.Id)
		helper.PanicIfError(err)
		parents = append(parents, parent)
	}
	return parents
}
