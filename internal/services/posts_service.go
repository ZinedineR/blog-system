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

type PostServiceImpl struct {
	db         *gorm.DB
	postRepo   repository.PostRepository
	userRepo   repository.UserRepository
	signaturer signature.Signaturer
	validate   *xvalidator.Validator
}

func (s *PostServiceImpl) Create(ctx context.Context, req *model.CreatePostReq) (*model.CreatePostRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity(ctx)
	userCheck, err := s.userRepo.FindByID(ctx, s.db, body.UserReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if userCheck == nil {
		return nil, exception.PermissionDenied("user/author does not exist")
	}
	if err := s.postRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreatePostRes{
		Posts: *body,
	}, nil
}

func (s *PostServiceImpl) Update(ctx context.Context, req *model.UpdatePostReq) (*model.UpdatePostRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity(ctx)
	postCheck, err := s.postRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if postCheck == nil {
		return nil, exception.NotFound("post not found")
	}
	if postCheck.UserReferencesId != body.UserReferencesId {
		return nil, exception.PermissionDenied("user/author does not match")
	}
	body.Id = postCheck.Id
	if err := s.postRepo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdatePostRes{
		Posts: *body,
	}, nil
}

func (s *PostServiceImpl) Delete(ctx context.Context, req *model.DeletePostReq) (*model.DeletePostRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	userReferences := constant.GetUserRefId(ctx)
	postCheck, err := s.postRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if postCheck == nil {
		return nil, exception.NotFound("post not found")
	}
	if postCheck.UserReferencesId != userReferences {
		return nil, exception.PermissionDenied("user/author does not match")
	}
	if err := s.postRepo.DeleteByIDTx(ctx, tx, req.ReferencesId); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.DeletePostRes{
		ReferencesId: req.ReferencesId,
	}, nil
}

func (s *PostServiceImpl) Find(ctx context.Context, req *model.GetAllPostsReq) (
	*model.GetAllPostsRes, *exception.Exception,
) {
	result, err := s.postRepo.FindByPagination(ctx, s.db, req.Page, req.Order, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get user", err)
	}
	return &model.GetAllPostsRes{
		PaginationData: *result,
	}, nil
}

func (s *PostServiceImpl) Detail(ctx context.Context, req *model.GetPostByIDReq) (
	*model.GetPostByIDRes, *exception.Exception,
) {
	result, err := s.postRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if result == nil {
		return nil, exception.NotFound("post not found, id: " + req.ReferencesId)
	}
	return &model.GetPostByIDRes{
		Posts: *result,
	}, nil
}
func NewPostService(
	db *gorm.DB, repo repository.PostRepository, userRepo repository.UserRepository,
	signaturer signature.Signaturer,
	validate *xvalidator.Validator,
) PostService {
	return &PostServiceImpl{
		db:         db,
		postRepo:   repo,
		userRepo:   userRepo,
		signaturer: signaturer,
		validate:   validate,
	}
}
