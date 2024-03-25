package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
	"github.com/rifqidamarali/final-project-golang006/pkg/helper"
)

type UserService interface {
	// GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id uint64) (model.User, error)
	DeleteUserById(ctx context.Context, id uint64) (model.User, error)
	UpdateUser(ctx context.Context, userData model.User) (model.UserUpdate, error)

	// activity
	SignUp(ctx context.Context, userSignUp model.UserSignUp) (model.User, error)
	SignIn(ctx context.Context, userSignIn model.UserSignIn) (model.User, error)
	
	// misc
	GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// func (u *userServiceImpl) GetUsers(ctx context.Context) ([]model.User, error) {
// 	users, err := u.repo.GetUsers(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, err
// }

func (u *userServiceImpl) GetUserById(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (u *userServiceImpl) DeleteUserById(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	// if user doesn't exist, return
	if user.ID == 0 {
		return model.User{}, nil
	}

	// delete user by id
	err = u.repo.DeleteUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (u *userServiceImpl) UpdateUser(ctx context.Context, user model.User) (model.UserUpdate, error){
	userUpdate := model.UserUpdate{}
	user.UpdatedAt = time.Now()
	userFind, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return userUpdate, err
	}

	if userFind.ID != 0 && user.ID != userFind.ID {
		return userUpdate, errors.New("email already exist")
	}

	user, err = u.repo.UpdateUser(ctx, user)
	if err != nil {
		return userUpdate, err
	}

	userUpdate.Email = user.Email
	userUpdate.Username = user.Username
	userUpdate.DoB = user.DoB
	return userUpdate, err
}

func (u *userServiceImpl) SignUp(ctx context.Context, userSignUp model.UserSignUp) (model.User, error) {
	// assumption: semua user adalah user baru
	user := model.User{
		Username: userSignUp.Username,
		Email:    userSignUp.Email,
		DoB:      userSignUp.DoB,
		// FirstName: userSignUp.FirstName,
		// LastName:  userSignUp.LastName,
	}

	// encryption password
	// hashing
	pass, err := helper.GenerateHash(userSignUp.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = pass

	// store to db
	res, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return res, err
}

func (u *userServiceImpl) SignIn(ctx context.Context, userSignIn model.UserSignIn) (model.User, error){
	user, err := u.repo.GetUserByEmail(ctx, userSignIn.Email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user does not exist")
	}

	isValidLogin := helper.CheckHash(userSignIn.Password, user.Password)
	if !isValidLogin {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}

func (u *userServiceImpl) GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error) {
	// generate claim
	now := time.Now()

	claim := model.StandardClaim{
		Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
		Iss: "MyGram",
		Aud: user.Username,
		Sub: "access-token",
		Exp: uint64(now.Add(time.Hour).Unix()),
		Iat: uint64(now.Unix()),
		Nbf: uint64(now.Unix()),
	}

	userClaim := model.AccessClaim{
		StandardClaim: claim,
		UserID:        user.ID,
		Username:      user.Username,
		Dob:           user.DoB,
	}

	token, err = helper.GenerateToken(userClaim)
	return
}
