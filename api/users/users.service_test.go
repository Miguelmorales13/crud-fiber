package users

import (
	"crud/api/users/dto"
	"crud/api/users/entities"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Create(t *testing.T) {
	type args struct {
		user *dto.UserCreateDto
	}
	tests := []struct {
		name            string
		args            args
		wantUserCreated entities.User
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			gotUserCreated, err := s.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUserCreated, tt.wantUserCreated) {
				t.Errorf("Create() gotUserCreated = %v, want %v", gotUserCreated, tt.wantUserCreated)
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAll(t *testing.T) {
	tests := []struct {
		name      string
		wantUsers []entities.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if gotUsers := s.GetAll(); !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("GetAll() = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
}

func TestService_GetOne(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name     string
		args     args
		wantUser entities.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if gotUser := s.GetOne(tt.args.id); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetOne() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestService_GetOneByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		args     args
		wantUser entities.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if gotUser := s.GetOneByEmail(tt.args.email); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetOneByEmail() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	type args struct {
		userUpdated dto.UserUpdateDto
		id          int
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.Update(tt.args.userUpdated, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
