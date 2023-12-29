package services

import (
	"gin/dto"
	"gin/entities"
	"gin/repositories"
	"gin/utils"
	"github.com/mashingan/smapping"
	"log"
)

type UserService interface {
	CreateUser(user dto.RegisterDTO) entities.User
	Update(user dto.UserUpdateDTO) entities.User
	Delete(userID string) error
	Profile(userID string) entities.User
	FindAll(id string) []entities.User
	IsDuplicateEmail(email string) bool
	SendActivationEmail(email string, hash string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) CreateUser(user dto.RegisterDTO) entities.User {
	userToCreate := entities.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	userToCreate.Hash = utils.RandomToken()
	return service.userRepository.InsertUser(userToCreate)
}

func (service *userService) Update(user dto.UserUpdateDTO) entities.User {
	userToUpdate := entities.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	return service.userRepository.UpdateUser(userToUpdate)
}

func (service *userService) Delete(userID string) error {
	return service.userRepository.Delete(userID)
}

func (service *userService) Profile(userID string) entities.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) FindAll(id string) []entities.User {
	return service.userRepository.FindAll(id)
}

func (service *userService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *userService) SendActivationEmail(email string, hash string) error {
	link := "http://localhost:8080/#/activate?hash=" + hash
	body := "<strong>Here is your activation <a href='" + link + "'>link</a></strong>"
	subject := "Account activation"
	return utils.SendMail(email, body, subject)
}
