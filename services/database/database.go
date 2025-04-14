package services

import (
	"log/slog"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string
	Password string
	Role     models.Role
}

type TokenModel struct {
	gorm.Model
	Token  string
	Bucket string
	UserID uint
	User   UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TestServiceAdapter struct {
	logger slog.Logger
	db     *gorm.DB
}

func NewServiceAdapter(logger slog.Logger) *TestServiceAdapter {
	db, err := gorm.Open(sqlite.Open("bsb.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserModel{}, &TokenModel{})

	return &TestServiceAdapter{db: db}
}

func (tsa *TestServiceAdapter) SetLogger(logger slog.Logger) {
	tsa.logger = logger
}

func (tsa *TestServiceAdapter) GetUserById(
	id uint,
) (*models.UserModel, error) {
	var user UserModel
	result := tsa.db.First(&user, id)
	if result.RowsAffected < 1 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserModel{
		UserId: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}, nil
}

func (tsa *TestServiceAdapter) GetUserByEmail(
	email string,
) (*models.UserModel, error) {
	var user UserModel
	result := tsa.db.First(&user, "email = ?", email)
	if result.RowsAffected < 1 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserModel{
		UserId: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}, nil
}

func (tsa *TestServiceAdapter) GetUserByCredentials(
	email, password string,
) (*models.UserModel, error) {
	var user UserModel
	result := tsa.db.First(&user, "email = ? AND password = ?", email, password)
	if result.RowsAffected < 1 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserModel{
		UserId: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}, nil
}

func (tsa *TestServiceAdapter) UpdateUser(
	id uint, email, password string, role models.Role,
) error {
	var user UserModel
	result := tsa.db.First(&user, id)
	if result.RowsAffected < 1 {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	if email != "" {
		user.Email = email
	}
	if password != "" {
		user.Password = password
	}
	if role != "" {
		user.Role = role
	}
	result = tsa.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tsa *TestServiceAdapter) DeleteUser(id uint) error {
	var user UserModel
	result := tsa.db.First(user, id)
	if result.RowsAffected < 1 {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	result = tsa.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tsa *TestServiceAdapter) InsertUser(
	email, pass string, role models.Role,
) (uint, error) {
	user := &UserModel{Email: email, Password: pass, Role: role}
	result := tsa.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (tsa *TestServiceAdapter) SendEmail(id uint, token, bucket string) error {
	user := &TokenModel{Token: token, Bucket: bucket, UserID: id}
	result := tsa.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tsa *TestServiceAdapter) GetIdByToken(
	token, bucket string,
) (uint, bool, error) {
	var tokenEntry TokenModel
	result := tsa.db.First(
		&tokenEntry,
		"token = ? AND bucket = ?",
		token,
		bucket,
	)
	if result.RowsAffected < 1 {
		return 0, false, nil
	}
	if result.Error != nil {
		return 0, false, result.Error
	}
	return tokenEntry.UserID, true, nil
}

func (tsa *TestServiceAdapter) DeleteToken(token, bucket string) error {
	var tokenEntry TokenModel
	result := tsa.db.First(
		&tokenEntry,
		"token = ? AND bucket = ?",
		token,
		bucket,
	)
	if result.RowsAffected < 1 {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	result = tsa.db.Delete(&tokenEntry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
