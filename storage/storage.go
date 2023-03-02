package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}
type AdminFilter struct {
	SearchTerm string
}
type ResultFilter struct {
	SearchTerm string
}
type User struct {
	ID        int          `form:"-" db:"id"`
	FirstName string       `form:"FirstName" db:"first_name"`
	LastName  string       `form:"LastName" db:"last_name"`
	Class_id  int          `form:"Class_id" db:"class_id"`
	Class_name string      `form:"Class_name" db:"class_name"`
	Email     string       `form:"Email" db:"email"`
	Username  string       `form:"Username" db:"username"`
	Password  string       `form:"Password" db:"password"`
	Status    bool         `form:"Status" db:"status"`
	CreatedAt time.Time    `form:"Created_at" db:"created_at"`
	UpdatedAt time.Time    `form:"Updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `form:"Deleted_at" db:"deleted_at"`
}
type Admin struct {
	ID        int          `form:"-" db:"id"`
	FirstName string       `form:"FirstName" db:"first_name"`
	LastName  string       `form:"LastName" db:"last_name"`
	Email     string       `form:"Email" db:"email"`
	Username  string       `form:"Username" db:"username"`
	Password  string       `form:"Password" db:"password"`
	Status    bool         `form:"Status" db:"status"`
	CreatedAt time.Time    `form:"-" db:"created_at"`
	UpdatedAt time.Time    `form:"-" db:"updated_at"`
	DeletedAt sql.NullTime `form:"-" db:"deleted_at"`
	Total     int          `form:"-" db:"total"`
}

type Class struct {
	ID        int          `form:"-" db:"id"`
	Class_name    string   `db:"class_name" form:"Class_name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Subject struct {
	ID        int          `form:"-" db:"id"`
	Subject   string       `db:"subject" form:"Subject"`
	Class_id     string    `db:"class" form:"Class_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
type MarkEdit struct {
	ID int       `form:"ID" db:"subject_id"`
	Marks string `form:"Marks" db:"marks"`
}

type ClassChoice struct {
	Class_id string `form:"class_id" db:"class_id"`
}

type StudentSubject struct {
	ID        int          `db:"id" form:"-"`
	StudentID int          `db:"student_id" form:"student_id"`
	SubjectID int          `db:"subject_id" form:"subject_id"`
	Marks     int          `db:"marks" form:"marks"`
	Mark      map[int]int
 	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
}
type MarkIN struct {
	ID        int          `db:"id" form:"-"`
	StudentID int          `db:"student_id" form:"student_id"`
	Subjects string        `db:"subject" form:"subject"`
	Class string           `db:"class" form:"class"`
	First_name string `db:"first_name" form:"first_name"`
	Last_name string    `db:"last_name" form:"last_name"`   
	SubjectID int          `db:"subject_id" form:"subject_id"`
	Marks     int          `db:"marks" form:"marks"`
	Mark      map[int]int
 	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
}
type Mark struct {
	Class_id   int `from:"Class_id"`
	Student_id int `form:"Student_id"`
}
type SubjectFrom struct {
	Id string  	`db:"userid"`
	Subject string `form:"Subject" db:"subject"`
	Subject_id string `form:"subject_id" db:"subject_id"`
	Class 	string `form:"Class" db:"class"`
	First_name string `form:"First_name" db:"first_name"`
	Last_name string `form:"Last_name" db:"last_name"`
    ID string `db:"id"`
}
func (u User) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.FirstName,
		validation.Required.Error("fast name can not be blank"),
		validation.Length(3, 45).Error("fast name must be between 3 to 45 characters"),
	),
		validation.Field(&u.LastName,
			validation.Required.Error("last name can not be blank"),
			validation.Length(3, 45).Error("last name must be between 3 to 45 characters"),
		),
		validation.Field(&u.Username,
			validation.Required.Error("username cannot be blank"),
		),
		validation.Field(&u.Class_id,
			validation.Required.Error("class cannot be blank"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email cannot be blank"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
		),
	)
}
func (u Admin) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.FirstName,
		validation.Required.Error("first name can not be blank"),
		validation.Length(3, 45).Error("first name must be between 3 to 45 characters"),
	),
		validation.Field(&u.LastName,
			validation.Required.Error("last name can not be blank"),
			validation.Length(3,45).Error("last name must be between 3 to 45 characters"),
		),
		validation.Field(&u.Username,
			validation.Required.Error("username cannot be blank"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("email cannot be blank"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
		),
	)
}
func (u Class) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.Class_name,
		validation.Required.Error("class can not be blank"),
		validation.Length(3, 45).Error("Name must be between 3 to 45 characters"),
	),
)
}
func (s Subject) Validate() error {
	return validation.ValidateStruct(&s, validation.Field(&s.Subject,
		validation.Required.Error("Subject can not be blank"),
		validation.Length(3,45).Error("Subject must be between 3 to 45 characters"),
	),
	validation.Field(&s.Class_id,
		validation.Required.Error("class cannot be blank"),
	),
	)
}
