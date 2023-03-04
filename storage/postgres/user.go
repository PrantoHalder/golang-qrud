package postgres

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"main.go/storage"
)

const listQuery = `

SELECT users.id,users.first_name,users.last_name,users.class_id,users.email,users.status,class.class_name
FROM users
FULL OUTER JOIN class ON class.id = users.class_id
WHERE
	users.deleted_at IS NULL
	AND 
    (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`

func (s PostGressStorage) ListUser(uf storage.UserFilter) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, listQuery, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}

const statusedit=`SELECT status,id
FROM users
WHERE
  deleted_at IS NULL
  AND
  id = $1`
func (s PostGressStorage) StatusEdit(id string) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, statusedit,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
const getStatusbyUsernameQuery=`SELECT status
FROM admin
WHERE
  deleted_at IS NULL
  AND
  username = $1`
func (s PostGressStorage) GetStatusbyUsernameQuery(username string) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, getStatusbyUsernameQuery,username); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
const getStatusbyUsernameQueryOfFaculty=`SELECT status
FROM faculty
WHERE
  deleted_at IS NULL
  AND
  username = $1`
func (s PostGressStorage) GetStatusbyUsernameQueryOFFaculty(username string) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, getStatusbyUsernameQueryOfFaculty,username); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
const getStatusbyUsernameQueryOfUsers=`SELECT status
FROM users
WHERE
  deleted_at IS NULL
  AND
  username = $1`
func (s PostGressStorage) GetStatusbyUsernameQueryOFUsers(username string) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, getStatusbyUsernameQueryOfUsers,username); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}

const Adminstatusedit=`SELECT status,id
FROM admin
WHERE
  deleted_at IS NULL
  AND
  id = $1`
func (s PostGressStorage) StatusEditAdmin(id string) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, Adminstatusedit,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
const facultystatusedit=`SELECT status,id
FROM faculty
WHERE
  deleted_at IS NULL
  AND
  id = $1`
func (s PostGressStorage) StatusEditfaculty(id string) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, facultystatusedit,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}

const UpdateStatusQuery = `
UPDATE users SET
status = $1
	WHERE id = $2 AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateStatus(status bool,id int) error {
	res, err := s.DB.Exec(UpdateStatusQuery,status,id)
	if err != nil {
			fmt.Println(err)
			return nil
	}
	
		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return err
		}
	
		if rowCount <= 0 {
			return nil
		}
	
		return nil
	}
const UpdateAdminStatusQuery = `
UPDATE admin SET
status = $1
	WHERE id = $2 AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateAdminStatus(status bool,id int) error {
	res, err := s.DB.Exec(UpdateAdminStatusQuery,status,id)
	if err != nil {
			fmt.Println(err)
			return nil
	}
	
		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return err
		}
	
		if rowCount <= 0 {
			return nil
		}
	
		return nil
	}

const UpdateFacultyStatusQuery = `
UPDATE faculty SET
status = $1
	WHERE id = $2 AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdatefacultyStatus(status bool,id int) error {
	res, err := s.DB.Exec(UpdateFacultyStatusQuery,status,id)
	if err != nil {
			fmt.Println(err)
			return nil
	}
	
		rowCount, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return err
		}
	
		if rowCount <= 0 {
			return nil
		}
	
		return nil
	}



const listQueryResult = `

SELECT id,first_name,last_name,class_id FROM users
WHERE
	deleted_at IS NULL
	AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`

func (s PostGressStorage) ListUserResult(uf storage.ResultFilter) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, listQueryResult, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}

const insertQuery = `
		INSERT INTO users(
			first_name,
			last_name,
			class_id,
			username,
			email,
			password
		) VALUES (
			:first_name,
			:last_name,
			:class_id,
			:username,
			:email,
			:password
		) RETURNING *;
	`

func (s PostGressStorage) CreateUser(u storage.User) (*storage.User, error) {

	stmt, err := s.DB.PrepareNamed(insertQuery)
	if err != nil {
		log.Fatalln(err)
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashPass)

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

const UpdateuserQuery = `
	UPDATE users SET
		first_name = :first_name,
		last_name = :last_name,
		class_id = :class_id,
		status = :status
	WHERE id = :id AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateUser(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(UpdateuserQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}

const GetUserByIDQuery = `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetUserByID(id string) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, GetUserByIDQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil
}
const GetMarkByIDQuery = `SELECT subject_class.subject, subject_class.class, users.first_name, users.last_name,student_subjects.subject_id,student_subjects.marks
FROM subject_class
FULL OUTER JOIN student_subjects ON subject_class.id = student_subjects.subject_id
FULL OUTER JOIN users ON users.id = student_subjects.student_id
WHERE student_subjects.student_id = $1`

func (s PostGressStorage) GetMarkByID(id string) ([]storage.MarkIN, error) {
	var u []storage.MarkIN
	if err := s.DB.Select(&u, GetMarkByIDQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}


const deleteUserbyID = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteUserByID(id string) error {
	res, err := s.DB.Exec(deleteUserbyID, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}
const UpdateMarksbyIDQuery = `UPDATE student_subjects
SET marks = $1
WHERE subject_id = $2 AND deleted_at IS NULL;`

func (s PostGressStorage) UpdateMarksbyID(marks string,id string) error {
	res, err := s.DB.Exec(UpdateMarksbyIDQuery, marks,id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}
const insertAdminQuery = `
		INSERT INTO admin(
			first_name,
			last_name,
			username,
			email,
			password
		) VALUES (
			:first_name,
			:last_name,
			:username,
			:email,
			:password
		) RETURNING *;
	`

func (s PostGressStorage) CreateAdmin(u storage.Admin) (*storage.Admin, error) {

	stmt, err := s.DB.PrepareNamed(insertAdminQuery)
	if err != nil {
		log.Fatalln(err)
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashPass)

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

const listAdminQuery = 
`
SELECT * FROM admin
WHERE
	deleted_at IS NULL
	AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`
func (s PostGressStorage) ListAdmin(uf storage.AdminFilter) ([]storage.Admin, error) {
	var listAdmin []storage.Admin
	if err := s.DB.Select(&listAdmin, listAdminQuery,uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listAdmin, nil
}

const GetAdminByIDQuery = `SELECT * FROM admin WHERE id = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetAdminByID(id string) (*storage.Admin, error) {
	var u storage.Admin
	if err := s.DB.Get(&u, GetAdminByIDQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil
}
const GetMarkEditQuery = `SELECT subject_id,marks FROM student_subjects WHERE subject_id = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) MarkEdit(id string) (*storage.MarkEdit, error) {
	var u storage.MarkEdit
	if err := s.DB.Get(&u, GetMarkEditQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil
}

const UpdateAdminQuery = `
	UPDATE admin SET
		first_name = :first_name,
		last_name = :last_name,
		status = :status
	WHERE id = :id AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateAdmin(u storage.Admin) (*storage.Admin, error) {
	stmt, err := s.DB.PrepareNamed(UpdateAdminQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}

const deleteAdminbyID = `UPDATE admin SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteAdminByID(username string) error {
	res, err := s.DB.Exec(deleteAdminbyID, username)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}

const GetAdminByUsernameQuery = `SELECT * FROM admin WHERE username = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetAdminByUsername(username string) (*storage.Admin, error) {
	var u storage.Admin
	if err := s.DB.Get(&u, GetAdminByUsernameQuery, username); err != nil {
		if err.Error() == NotFound {
			return nil, fmt.Errorf(NotFound)
		}
		return nil, err
	}
	return &u, nil
}

const GetFacultyByUsernameQuery = `SELECT * FROM faculty WHERE username = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetFacultyByUsername(username string) (*storage.Admin, error) {
	var u storage.Admin
	if err := s.DB.Get(&u, GetFacultyByUsernameQuery, username); err != nil {
		if err.Error() == NotFound {
			return nil, fmt.Errorf(NotFound)
		}
		return nil, err
	}
	return &u, nil
}
const GetUsersByUsernameQuery = `SELECT * FROM faculty WHERE username = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetUsersByUsername(username string) (*storage.Admin, error) {
	var u storage.Admin
	if err := s.DB.Get(&u, GetUsersByUsernameQuery, username); err != nil {
		if err.Error() == NotFound {
			return nil, fmt.Errorf(NotFound)
		}
		return nil, err
	}
	return &u, nil
}

const insertFacultyQuery = `
		INSERT INTO faculty(
			first_name,
			last_name,
			username,
			email,
			password
		) VALUES (
			:first_name,
			:last_name,
			:username,
			:email,
			:password
		) RETURNING *;
	`

func (s PostGressStorage) CreateFaculty(u storage.Admin) (*storage.Admin, error) {

	stmt, err := s.DB.PrepareNamed(insertFacultyQuery)
	if err != nil {
		log.Fatalln(err)
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashPass)

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

const listFacultyQuery = 
`
SELECT * FROM faculty
WHERE
	deleted_at IS NULL
	AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`
func (s PostGressStorage) ListFaculty(lf storage.AdminFilter) ([]storage.Admin, error) {
	var listUser []storage.Admin
	if err := s.DB.Select(&listUser, listFacultyQuery,lf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}

const UpdateFacultyQuery = `
	UPDATE faculty SET
		first_name = :first_name,
		last_name = :last_name,
		status = :status
	WHERE id = :id AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateFaculty(u storage.Admin) (*storage.Admin, error) {
	stmt, err := s.DB.PrepareNamed(UpdateFacultyQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}

const GetFacultyByIDQuery = `SELECT * FROM faculty WHERE id = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetFacultyByID(id string) (*storage.Admin, error) {
	var u storage.Admin
	if err := s.DB.Get(&u, GetFacultyByIDQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil
}

const deleteFacultybyIDQuery = `UPDATE faculty SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteFacultyByID(id string) error {
	res, err := s.DB.Exec(deleteFacultybyIDQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}

const insertClassQuery = `
		INSERT INTO class(
			class_name
		) VALUES (
			:class_name
		) RETURNING *;
	`

func (s PostGressStorage) CreateClass(u storage.Class) (*storage.Class, error) {

	stmt, err := s.DB.PrepareNamed(insertClassQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

const listClassQuery = `SELECT * FROM class WHERE deleted_at IS NULL ORDER BY id ASC ;`

func (s PostGressStorage) ListClass() ([]storage.Class, error) {
	var listClass []storage.Class
	if err := s.DB.Select(&listClass, listClassQuery); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listClass, nil
}

const GetClassByIDQuery = `SELECT * FROM class WHERE id = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetClassByID(id string) (*storage.Class, error) {
	var u storage.Class
	if err := s.DB.Get(&u, GetClassByIDQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil
}

const UpdateClassQuery = `
	UPDATE class SET
		class_name = :class_name
	WHERE id = :id AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateClass(u storage.Class) (*storage.Class, error) {
	stmt, err := s.DB.PrepareNamed(UpdateClassQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}

const deleteClassbyIDQuery = `UPDATE class SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteClassyByID(id string) error {
	res, err := s.DB.Exec(deleteClassbyIDQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}

const GetClass = `SELECT id, class_name FROM class WHERE deleted_at IS NULL;`

func (s PostGressStorage) GetClass() ([]storage.Class, error) {
	var u []storage.Class
	if err := s.DB.Select(&u, GetClass); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}

const insertSubjectQuery = `
		INSERT INTO subject_class(
			subject,
			class
		) VALUES (
			:subject,
			:class
		) RETURNING id;
	`

func (s PostGressStorage) CreateSubject(u storage.Subject) (*storage.Subject, error) {

	stmt, err := s.DB.PrepareNamed(insertSubjectQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}
const GetSubjectByClassIDQuery = `SELECT id,subject FROM subject_class WHERE class = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetSubjectByClassID(id string) ([]storage.Subject, error) {
	var u []storage.Subject
	if err := s.DB.Select(&u,GetSubjectByClassIDQuery,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}
const GetSubjectByIDQuery = `SELECT id,subject FROM subject_class WHERE id = $1 AND deleted_at IS NULL;`

func (s PostGressStorage) GetSubjectByID(id string) (*storage.Subject, error) {
	var u storage.Subject
	if err := s.DB.Get(&u, GetSubjectByIDQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
const UpdateSubjectQuery = `
	UPDATE subject_class SET
		subject = :subject
	WHERE id = :id AND deleted_at is NULL
	RETURNING *;
	`

func (s PostGressStorage) UpdateSubject(u storage.Subject) (*storage.Subject, error) {
	stmt, err := s.DB.PrepareNamed(UpdateSubjectQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
const deleteSubjectbyIDQuery = `UPDATE subject_class SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteSubjectyByID(id string) error {
	res, err := s.DB.Exec(deleteSubjectbyIDQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}
const getSubjectByClassIDQuery = `SELECT * FROM subject_class WHERE class=$1 AND deleted_at IS NULL`

func (s PostGressStorage) GetSubjectByClassID2(class int) ([]storage.Subject, error) {

	var u []storage.Subject
	if err := s.DB.Select(&u, getSubjectByClassIDQuery, class); err != nil {
		log.Println(err)
		return nil, err
	}
	return u, nil
}

const insertMarkQuery = `
	INSERT INTO student_subjects(
		student_id,
		subject_id,
        marks
		)  
	VALUES(
		:student_id,
		:subject_id,
		:marks
		)RETURNING *;
	`

func (p PostGressStorage) InsertMark(s storage.StudentSubject) (*storage.StudentSubject, error) {

	stmt, err := p.DB.PrepareNamed(insertMarkQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}
const GetStudet = `SELECT id,first_name,last_name FROM users WHERE deleted_at IS NULL;`

func (s PostGressStorage) GetStudent() ([]storage.User, error) {
	var u []storage.User
	if err := s.DB.Select(&u, GetStudet); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}
const GetSubjectQuery = `SELECT subject_class.subject, subject_class.class, users.first_name, users.last_name, users.id,student_subjects.subject_id,student_subjects.id
FROM subject_class
FULL OUTER JOIN student_subjects ON subject_class.id = student_subjects.subject_id
FULL OUTER JOIN users ON users.id = student_subjects.student_id
WHERE users.id = $1
ORDER BY subject_class.subject;`

func (s PostGressStorage) GetSubject(id string) ([]storage.SubjectFrom, error) {
	var u []storage.SubjectFrom
	if err := s.DB.Select(&u, GetSubjectQuery,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}

const createmarkQuery = `
UPDATE student_subjects
SET marks = :marks
WHERE id = :id
	returning *;`

func (p PostGressStorage) Markcreate(s storage.StudentSubject) (*storage.StudentSubject, error) {

	stmt, _ := p.DB.PrepareNamed(createmarkQuery)

	stmt.Get(&s, s)

	return &s, nil

}
const getFacultyEditQuery = `SELECT id
FROM faculty
WHERE
  deleted_at IS NULL
  AND
  id = $1`
  func (s PostGressStorage) GetFacultyEdit(id string) ([]storage.User, error) {
	var u []storage.User
	if err := s.DB.Select(&u,getFacultyEditQuery,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}
const getUserEditQuery = `SELECT id
FROM users
WHERE
  deleted_at IS NULL
  AND
  id = $1`
  func (s PostGressStorage) GetUserEditQuery(id string) ([]storage.User, error) {
	var u []storage.User
	if err := s.DB.Select(&u,getUserEditQuery,id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}