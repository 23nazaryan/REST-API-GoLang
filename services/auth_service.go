package services

import (
	"errors"
	"gin/dto"
	"gin/entities"
	"gin/repositories"
	"gin/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	VerifyHash(hashDTO dto.HashDTO) interface{}
	FindByID(id string) entities.User
	IsDuplicateEmail(email string) bool
	SetPassword(pwdDTO dto.PwdDTO) entities.User
	SendForgotEmail(email string) error
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRep repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) FindByID(id string) entities.User {
	return service.userRepository.ProfileUser(id)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email)
	if v, ok := res.(entities.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}

		return false
	}

	return false
}

func (service *authService) VerifyHash(hashDTO dto.HashDTO) interface{} {
	return service.userRepository.VerifyHash(hashDTO)
}

func (service *authService) SetPassword(pwdDTO dto.PwdDTO) entities.User {
	return service.userRepository.SetPassword(pwdDTO)
}

func (service *authService) SendForgotEmail(email string) error {
	res := service.userRepository.FindByEmail(email)

	if user, ok := res.(entities.User); ok {
		user.Hash = utils.RandomToken()
		updatedUser := service.userRepository.UpdateUser(user)
		link := "http://localhost:8080/#/forgot?hash=" + updatedUser.Hash
		body := "<strong>Here is your forgot password <a href='" + link + "'>link</a></strong>"
		subject := "Forgot Password"
		err := utils.SendMail(email, body, subject)
		return err
	}

	return errors.New("user is not found")
}

func comparedPassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
