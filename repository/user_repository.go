package repository

import (
	model "github.com/aruncs31s/esdcmodels"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserRepositoryReader
	UserRepositoryWriter
}

type UserRepositoryReader interface {
	FindByUsername(username string) (*model.User, error)
	FindUsersByUsernames(usernames []string) (*[]model.User, error)
	FindByID(id uint) (*model.User, error)
	FindUserIDByUsername(username string) (uint, error)
	FindByEmail(email string) (*model.User, error)
	GetAllUsers(limit int, offset int) (*[]model.User, error)
	GetAllUsersEssentials() (*[]model.User, error)
	GetUsersCount() (int, error)
	SearchUsers(query string) (*[]model.User, error)
}
type UserRepositoryWriter interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUserByID(userID uint) error
}
type userRepositoryReader struct {
	db *gorm.DB
}
type userRepositoryWriter struct {
	db *gorm.DB
}

type userRepository struct {
	reader UserRepositoryReader
	writer UserRepositoryWriter
}

func NewUserRepository(db *gorm.DB) UserRepository {
	reader := newUserRepositoryReader(db)
	writer := newUserRepositoryWriter(db)
	return &userRepository{
		reader: reader,
		writer: writer,
	}
}
func newUserRepositoryReader(db *gorm.DB) UserRepositoryReader {
	return &userRepositoryReader{db: db}
}
func newUserRepositoryWriter(db *gorm.DB) UserRepositoryWriter {
	return &userRepositoryWriter{db: db}
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	return r.reader.FindByUsername(username)
}
func (r *userRepository) FindUsersByUsernames(usernames []string) (*[]model.User, error) {
	return r.reader.FindUsersByUsernames(usernames)
}
func (r *userRepository) FindByID(id uint) (*model.User, error) {
	return r.reader.FindByID(id)
}
func (r *userRepository) FindUserIDByUsername(username string) (uint, error) {
	return r.reader.FindUserIDByUsername(username)
}
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	return r.reader.FindByEmail(email)
}
func (r *userRepository) GetAllUsers(limit int, offset int) (*[]model.User, error) {
	return r.reader.GetAllUsers(limit, offset)
}
func (r *userRepository) GetUsersCount() (int, error) {
	return r.reader.GetUsersCount()
}
func (r *userRepository) CreateUser(user *model.User) error {
	return r.writer.CreateUser(user)
}
func (r *userRepository) UpdateUser(user *model.User) error {
	return r.writer.UpdateUser(user)
}
func (r *userRepository) DeleteUserByID(userID uint) error {
	return r.writer.DeleteUserByID(userID)
}
func (r *userRepository) SearchUsers(query string) (*[]model.User, error) {
	return r.reader.SearchUsers(query)
}

func (r *userRepositoryReader) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
func (r *userRepositoryReader) FindByID(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *userRepositoryReader) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}
func (r *userRepositoryReader) FindUsersByUsernames(usernames []string) (*[]model.User, error) {
	var users []model.User
	result := r.db.Where("username IN ?", usernames).Find(&users)
	return &users, result.Error
}

func (r *userRepositoryReader) FindUserIDByUsername(username string) (uint, error) {
	var user model.User
	result := r.db.Select("id").Where("username = ?", username).First(&user)
	return user.ID, result.Error
}

func (r *userRepositoryWriter) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *userRepositoryWriter) UpdateUser(user *model.User) error {
	result := r.db.Save(user)
	return result.Error
}
func (r *userRepositoryReader) GetAllUsers(limit int, offset int) (*[]model.User, error) {
	var users []model.User
	result := r.db.Preload("Github").Limit(limit).Offset(offset).Find(&users)
	return &users, result.Error
}

func (r *userRepositoryReader) GetUsersCount() (int, error) {
	var count int64
	result := r.db.Model(&model.User{}).Count(&count)
	return int(count), result.Error
}
func (r *userRepositoryWriter) DeleteUserByID(userID uint) error {
	result := r.db.Delete(&model.User{}, userID)
	return result.Error
}

func (r *userRepositoryReader) SearchUsers(query string) (*[]model.User, error) {
	var users []model.User
	// Search by name or email with case-insensitive LIKE query
	err := r.db.Where("LOWER(name) LIKE LOWER(?) OR LOWER(email) LIKE LOWER(?)",
		"%"+query+"%", "%"+query+"%").
		Limit(1000).
		Find(&users).Error
	return &users, err
}
