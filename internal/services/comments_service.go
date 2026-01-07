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

type CommentServiceImpl struct {
	db          *gorm.DB
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
	signaturer  signature.Signaturer
	validate    *xvalidator.Validator
}

func (s *CommentServiceImpl) Create(ctx context.Context, req *model.CreateCommentReq) (*model.CreateCommentRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity(ctx)
	postCheck, err := s.postRepo.FindByID(ctx, s.db, body.PostReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if postCheck == nil {
		return nil, exception.PermissionDenied("post does not exist")
	}

	if err := s.commentRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreateCommentRes{
		Comments: *body,
	}, nil
}

func (s *CommentServiceImpl) Update(ctx context.Context, req *model.UpdateCommentReq) (*model.UpdateCommentRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity(ctx)
	commentCheck, err := s.commentRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if commentCheck == nil {
		return nil, exception.NotFound("comment not found")
	}

	postCheck, err := s.postRepo.FindByID(ctx, s.db, commentCheck.PostReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if postCheck == nil {
		return nil, exception.PermissionDenied("post does not exist")
	}
	//if the comment is anonymous, it cannot be updated
	if commentCheck.UserReferencesId == nil {
		return nil, exception.PermissionDenied("this comment was made anonymously, cannot be updated")
	}

	if commentCheck.UserReferencesId != nil {
		if body.UserReferencesId == nil {
			return nil, exception.InvalidArgument("this comment belongs to someone, cannot be updated anonymously")
		}
		isCommentAuthor := *commentCheck.UserReferencesId == *body.UserReferencesId
		//if the comment is not anonymous, it can only be updated by the comment's author
		if !isCommentAuthor {
			return nil, exception.PermissionDenied("you do not have permission to delete this comment")
		}
	} else {
		return nil, exception.PermissionDenied("no authorization")
	}
	body.Id = commentCheck.Id
	if err := s.commentRepo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdateCommentRes{
		Comments: *body,
	}, nil
}

func (s *CommentServiceImpl) Delete(ctx context.Context, req *model.DeleteCommentReq) (*model.DeleteCommentRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	userReferences := constant.GetUserRefId(ctx)
	commentCheck, err := s.commentRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if commentCheck == nil {
		return nil, exception.NotFound("comment not found")
	}
	postCheck, err := s.postRepo.FindByID(ctx, s.db, commentCheck.PostReferencesId)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	if postCheck == nil {
		return nil, exception.PermissionDenied("post does not exist")
	}
	if commentCheck.UserReferencesId != nil {
		isPostAuthor := postCheck.UserReferencesId == userReferences
		isCommentAuthor := *commentCheck.UserReferencesId == userReferences

		if !isPostAuthor && !isCommentAuthor {
			return nil, exception.PermissionDenied("you do not have permission to delete this comment")
		}
	}
	if err := s.commentRepo.DeleteByIDTx(ctx, tx, req.ReferencesId); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.DeleteCommentRes{
		ReferencesId: req.ReferencesId,
	}, nil
}

func (s *CommentServiceImpl) Find(ctx context.Context, req *model.GetAllCommentsReq) (
	*model.GetAllCommentsRes, *exception.Exception,
) {
	result, err := s.commentRepo.FindByPagination(ctx, s.db, req.Page, req.Order, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get user", err)
	}
	return &model.GetAllCommentsRes{
		PaginationData: *result,
	}, nil
}

func (s *CommentServiceImpl) Detail(ctx context.Context, req *model.GetCommentByIDReq) (
	*model.GetCommentByIDRes, *exception.Exception,
) {
	result, err := s.commentRepo.FindByID(ctx, s.db, req.ReferencesId)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if result == nil {
		return nil, exception.NotFound("post not found, id: " + req.ReferencesId)
	}
	return &model.GetCommentByIDRes{
		Comments: *result,
	}, nil
}
func NewCommentService(
	db *gorm.DB, repo repository.CommentRepository, postRepo repository.PostRepository,
	signaturer signature.Signaturer,
	validate *xvalidator.Validator,
) CommentService {
	return &CommentServiceImpl{
		db:          db,
		commentRepo: repo,
		postRepo:    postRepo,
		signaturer:  signaturer,
		validate:    validate,
	}
}
