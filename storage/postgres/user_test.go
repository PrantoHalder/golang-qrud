package postgres

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/storage"
)

func TestCreateUser(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "CREATE_USER_SUCCESS",
			in: storage.User{
				FirstName: "first",
				LastName:  "last",
				Class_id:  2,
				Email:     "first@example.com",
				Username:  "user",
				Password:  "123456",
			},
			want: &storage.User{
				FirstName: "first",
				LastName:  "last",
				Class_id:  2,
				Email:     "first@example.com",
				Username:  "user",
				Password:  "123456",
				Status:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.CreateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateUser() error = got %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestCreateAdmin(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Admin
		want    *storage.Admin
		wantErr bool
	}{
		{
			name: "ADMIN_CREATE_SUCCESS",
			in: storage.Admin{
				FirstName: "first",
				LastName:  "last",
				Email:     "first@gmail.com",
				Username:  "first",
				Password:  "123456",
			},
			want: &storage.Admin{
				FirstName: "first",
				LastName:  "last",
				Email:     "first@gmail.com",
				Username:  "first",
				Password:  "123456",
				Status:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateAdmin(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Admin{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestCreateFaculty(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Admin
		want    *storage.Admin
		wantErr bool
	}{
		{
			name: "FACULTY_CREATE_SUCCESS",
			in: storage.Admin{
				FirstName: "first",
				LastName:  "last",
				Email:     "first@gmail.com",
				Username:  "first",
				Password:  "123456",
			},
			want: &storage.Admin{
				FirstName: "first",
				LastName:  "last",
				Email:     "first@gmail.com",
				Username:  "first",
				Password:  "123456",
				Status:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateFaculty(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Admin{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestCreateClass(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Class
		want    *storage.Class
		wantErr bool
	}{
		{
			name: "CREATE_CLASS_SUCCESS",
			in: storage.Class{
				Class_name: "FAST",
			},
			want: &storage.Class{
				Class_name: "FAST"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateClass(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Class{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestCreateSubject(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Subject
		want    *storage.Subject
		wantErr bool
	}{
		{
			name: "CREATE_SUBJECT_SUCCESS",
			in: storage.Subject{
				Subject:  "Class-1",
				Class_id: "1",
			},
			want: &storage.Subject{
				Subject:  "Class-1",
				Class_id: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.CreateSubject(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Subject{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
func TestMarkcreate(t *testing.T) {
	p := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.StudentSubject
		want    *storage.StudentSubject
		wantErr bool
	}{
	   {
	   	name:    "CREATE_MARK_SUCCESS",
	   	in:      storage.StudentSubject{
	   		StudentID: 1,
	   		SubjectID: 1,
	   		Marks:     45,
	   	},
	   	want:    &storage.StudentSubject{
			StudentID: 1,
	   		SubjectID: 1,
	   		Marks:     45,
		},
	   },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			got, err := p.Markcreate(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Markcreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

            opts := cmp.Options{
				cmpopts.IgnoreFields(storage.StudentSubject{}, "ID","Mark","CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want,opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
