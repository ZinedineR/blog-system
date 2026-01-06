package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"blog-system/pkg/constant"
	"blog-system/pkg/signature"
	"context"
	//"blog-system/pkg/exception"
	"blog-system/pkg/exception"
	"blog-system/pkg/xvalidator"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db         *gorm.DB
	userRepo   repository.UserRepository
	signaturer signature.Signaturer
	validate   *xvalidator.Validator
}

func NewUserService(
	db *gorm.DB, repo repository.UserRepository,
	signaturer signature.Signaturer,
	validate *xvalidator.Validator,
) UserService {
	return &UserServiceImpl{
		db:         db,
		userRepo:   repo,
		signaturer: signaturer,
		validate:   validate,
	}
}

func (s *UserServiceImpl) Register(
	ctx context.Context, req *model.CreateUserReq,
) (*model.CreateUserRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	duplicateCheck, err := s.userRepo.FindByMap(ctx, s.db, map[string]interface{}{
		"username": req.Username,
	})
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if duplicateCheck != nil {
		return nil, exception.PermissionDenied("username already exists")
	}

	password, err := s.signaturer.HashBscryptPassword(req.Password)
	if err != nil {
		return nil, exception.Internal("can't create password", err)
	}
	body := req.ToEntity(password)
	if err := s.userRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreateUserRes{
		Users: *body,
	}, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req *model.CreateUserReq) (
	*model.LoginUserRes, *exception.Exception,
) {
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	result, err := s.userRepo.FindByMap(ctx, s.db, map[string]interface{}{
		"username": req.Username,
	})
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if result == nil {
		return nil, exception.NotFound("username not found")
	}
	if ok := s.signaturer.CheckBscryptPasswordHash(req.Password, result.Password); !ok {
		return nil, exception.PermissionDenied("username/password unmatched")
	}
	jwtToken, err := s.signaturer.GenerateJWT(result.ReferencesId, result.Username)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	return &model.LoginUserRes{
		Username: result.Username,
		Token:    jwtToken,
	}, nil
}

func (s *UserServiceImpl) Find(ctx context.Context, req *model.GetAllUsersReq) (
	*model.GetAllUsersRes, *exception.Exception,
) {
	result, err := s.userRepo.FindByPagination(ctx, s.db, req.Page, req.Order, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get user", err)
	}
	return &model.GetAllUsersRes{
		PaginationData: *result,
	}, nil
}

func (s *UserServiceImpl) Detail(ctx context.Context, req *model.GetUserByIDReq) (
	*model.GetUserByIDRes, *exception.Exception,
) {
	req.ReferencesId = constant.GetUserRefId(ctx)
	result, err := s.userRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if result == nil {
		return nil, exception.NotFound("user not found, id: " + req.ReferencesId)
	}
	return &model.GetUserByIDRes{
		Users: *result,
	}, nil
}
