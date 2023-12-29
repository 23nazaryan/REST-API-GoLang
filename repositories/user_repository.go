package repositories

import (
	"gin/dto"
	"gin/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	InsertUser(user entities.User) entities.User
	UpdateUser(user entities.User) entities.User
	Delete(userID string) error
	VerifyCredential(email string) interface{}
	VerifyHash(hashDTO dto.HashDTO) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) interface{}
	ProfileUser(userID string) entities.User
	FindAll(id string) []entities.User
	SetPassword(pwdDTO dto.PwdDTO) entities.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entities.User) entities.User {
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entities.User) entities.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entities.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string) interface{} {
	var user entities.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}

	return nil
}

func (db *userConnection) VerifyHash(hashDTO dto.HashDTO) interface{} {
	var user entities.User
	var res *gorm.DB

	if hashDTO.Action == "reset" {
		res = db.connection.Where("hash = ? AND updated_at >= DATE_SUB(NOW(), INTERVAL 24 HOUR)", hashDTO.Hash).Take(&user)
	} else {
		res = db.connection.Where("hash = ? AND created_at >= DATE_SUB(NOW(), INTERVAL 24 HOUR)", hashDTO.Hash).Take(&user)
	}

	if res.Error != nil {
		return nil
	}

	return user
}

func (db *userConnection) SetPassword(pwdDTO dto.PwdDTO) entities.User {
	var user entities.User
	res := db.connection.Where("hash = ?", pwdDTO.Hash).Take(&user)
	if res.Error != nil {
		panic("Failed to get user by hash")
	}

	user.Password = hashAndSalt([]byte(pwdDTO.Password))
	user.Hash = ""
	db.connection.Save(&user)
	return user
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) Delete(userID string) error {
	var user entities.User
	db.connection.Find(&user, userID)
	res := db.connection.Delete(&user)
	return res.Error
}

func (db *userConnection) FindByEmail(email string) interface{} {
	var user entities.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return res.Error
	}
	return user
}

func (db *userConnection) ProfileUser(userID string) entities.User {
	var user entities.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) FindAll(id string) []entities.User {
	var users []entities.User
	db.connection.Where("id != ?", id).Find(&users)
	return users
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}

	return string(hash)
}
