package service

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
	"github.com/zetsux/gin-gorm-clean-starter/common/constant"
	"github.com/zetsux/gin-gorm-clean-starter/common/util"
	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	errs "github.com/zetsux/gin-gorm-clean-starter/core/helper/errors"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type userService struct {
	userRepository       repository.UserRepository
	animalTypeRepository repository.AnimalTypeRepository
	questRepository      repository.QuestRepository
}

type UserService interface {
	VerifyLogin(ctx context.Context, username string, password string) bool
	CreateNewUser(ctx context.Context, ud dto.UserAuthenticationRequest) (dto.UserResponse, error)
	GetAllUsers(ctx context.Context, req base.GetsRequest) ([]dto.UserResponse, base.PaginationResponse, error)
	GetUserByPrimaryKey(ctx context.Context, key string, value string) (dto.UserResponse, error)
	UpdateUserByID(ctx context.Context, ud dto.UserUpdateRequest, id string) (dto.UserResponse, error)
	DeleteUserByID(ctx context.Context, id string) error
	ChangePicture(ctx context.Context, req dto.UserChangePictureRequest, userID string) (dto.UserResponse, error)
	DeletePicture(ctx context.Context, userID string) error
}

func NewUserService(userR repository.UserRepository, atr repository.AnimalTypeRepository, qr repository.QuestRepository) UserService {
	return &userService{userRepository: userR, animalTypeRepository: atr, questRepository: qr}
}

func (us *userService) VerifyLogin(ctx context.Context, username string, password string) bool {
	userCheck, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, constant.DBAttrUsername, username)
	if err != nil {
		return false
	}
	passwordCheck, err := util.PasswordCompare(userCheck.Password, []byte(password))
	if err != nil {
		return false
	}

	if userCheck.Username == username && passwordCheck {
		return true
	}
	return false
}

func (us *userService) CreateNewUser(ctx context.Context, ud dto.UserAuthenticationRequest) (dto.UserResponse, error) {
	db, err := us.userRepository.TxRepository().BeginTx(ctx)
	if err != nil {
		return dto.UserResponse{}, err
	}
	defer us.userRepository.TxRepository().CommitOrRollbackTx(ctx, db, nil)

	userCheck, err := us.userRepository.GetUserByPrimaryKey(ctx, db, constant.DBAttrUsername, ud.Username)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if !(reflect.DeepEqual(userCheck, entity.User{})) {
		return dto.UserResponse{}, errs.ErrUsernameAlreadyExists
	}

	user := entity.User{
		Username:    ud.Username,
		Password:    ud.Password,
		Role:        constant.EnumRoleUser,
		LastAttempt: time.Now(),
	}

	// create new user
	newUser, err := us.userRepository.CreateNewUser(ctx, db, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	animalTypes, err := us.animalTypeRepository.GetAllAnimalTypes(ctx, db)
	if err != nil {
		return dto.UserResponse{}, err
	}
	for _, animalType := range animalTypes {
		us.questRepository.CreateNewQuest(ctx, db, entity.Quest{
			Step:         1,
			UserID:       newUser.ID.String(),
			AnimalTypeID: animalType.ID.String(),
		})
	}

	return dto.UserResponse{
		ID:       newUser.ID.String(),
		Username: newUser.Username,
		Role:     newUser.Role,
	}, nil
}

func (us *userService) GetAllUsers(ctx context.Context, req base.GetsRequest) (
	userResp []dto.UserResponse, pageResp base.PaginationResponse, err error) {
	if req.Limit < 0 {
		req.Limit = 0
	}

	if req.Page < 0 {
		req.Page = 0
	}

	if req.Sort != "" && req.Sort[0] == '-' {
		req.Sort = req.Sort[1:] + " DESC"
	}

	users, lastPage, total, err := us.userRepository.GetAllUsers(ctx, nil, req)
	if err != nil {
		return []dto.UserResponse{}, base.PaginationResponse{}, err
	}

	for _, user := range users {
		userResp = append(userResp, dto.UserResponse{
			ID:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
			Picture:  user.Picture,
		})
	}

	if req.Limit == 0 {
		return userResp, base.PaginationResponse{}, nil
	}

	pageResp = base.PaginationResponse{
		Page:     int64(req.Page),
		Limit:    int64(req.Limit),
		LastPage: lastPage,
		Total:    total,
	}
	return userResp, pageResp, nil
}

func (us *userService) GetUserByPrimaryKey(ctx context.Context, key string, val string) (dto.UserResponse, error) {
	user, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, key, val)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
		Picture:  user.Picture,
	}, nil
}

func (us *userService) UpdateUserByID(ctx context.Context,
	ud dto.UserUpdateRequest, id string) (dto.UserResponse, error) {
	user, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, constant.DBAttrID, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if reflect.DeepEqual(user, entity.User{}) {
		return dto.UserResponse{}, errs.ErrUserNotFound
	}

	if ud.Username != "" && ud.Username != user.Username {
		us, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, constant.DBAttrUsername, ud.Username)
		if err != nil {
			return dto.UserResponse{}, err
		}

		if !(reflect.DeepEqual(us, entity.User{})) {
			return dto.UserResponse{}, errs.ErrUsernameAlreadyExists
		}
	}

	userEdit := entity.User{
		ID:       user.ID,
		Username: ud.Username,
		Role:     ud.Role,
		Password: ud.Password,
	}

	edited, err := us.userRepository.UpdateUser(ctx, nil, userEdit)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if edited.Role == "" {
		edited.Role = user.Role
	}

	return dto.UserResponse{
		ID:       edited.ID.String(),
		Username: edited.Username,
		Role:     edited.Role,
		Picture:  user.Picture,
	}, nil
}

func (us *userService) DeleteUserByID(ctx context.Context, id string) error {
	userCheck, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, constant.DBAttrID, id)
	if err != nil {
		return err
	}

	if reflect.DeepEqual(userCheck, entity.User{}) {
		return errs.ErrUserNotFound
	}

	err = us.userRepository.DeleteUserByID(ctx, nil, id)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) ChangePicture(ctx context.Context,
	req dto.UserChangePictureRequest, userID string) (dto.UserResponse, error) {
	user, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, constant.DBAttrID, userID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if reflect.DeepEqual(user, entity.User{}) {
		return dto.UserResponse{}, errs.ErrUserNotFound
	}

	if user.Picture != "" {
		if err := util.DeleteFile(user.Picture); err != nil {
			return dto.UserResponse{}, err
		}
	}

	picID := uuid.New()
	picPath := fmt.Sprintf("user_picture/%v", picID)

	userEdit := entity.User{
		ID:      user.ID,
		Picture: picPath,
	}

	if err := util.UploadFile(req.Picture, picPath); err != nil {
		return dto.UserResponse{}, err
	}

	userUpdate, err := us.userRepository.UpdateUser(ctx, nil, userEdit)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:      userUpdate.ID.String(),
		Picture: userUpdate.Picture,
	}, nil
}

func (us *userService) DeletePicture(ctx context.Context, userID string) error {
	user, err := us.userRepository.GetUserByPrimaryKey(ctx, nil, constant.DBAttrID, userID)
	if err != nil {
		return err
	}

	if reflect.DeepEqual(user, entity.User{}) {
		return errs.ErrUserNotFound
	}

	if user.Picture == "" {
		return errs.ErrUserNoPicture
	}

	if err := util.DeleteFile(user.Picture); err != nil {
		return err
	}

	userEdit := entity.User{
		ID:      user.ID,
		Picture: "",
	}

	_, err = us.userRepository.UpdateUser(ctx, nil, userEdit)
	if err != nil {
		return err
	}

	return nil
}
