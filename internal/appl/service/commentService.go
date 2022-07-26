package service

import (
	"github.com/google/uuid"

	"github.com/johnHPX/blog-hard-backend/internal/domain/models"
	"github.com/johnHPX/blog-hard-backend/internal/infra/repository"
	"github.com/johnHPX/validator-hard/pkg/validator"
)

type commentServiceInterface interface {
	CreateComment(postID, title, content string) error
	ListCommentsPost(postID string, offset, limit, page int) ([]models.Comment, int, error)
	ListCommentsUser(offset, limit, page int) ([]models.Comment, int, error)
	ListCommentsPostUser(postID string, offset, limit, page int) ([]models.Comment, int, error)
	FindComment(commentID string) (*models.Comment, error)
	UpdateComment(commentID, title, content string) error
	RemoveComment(commentID string) error
}

type commentServiceImpl struct {
	userID string
	kind   string
}

func (s *commentServiceImpl) CreateComment(postID, title, content string) error {
	val := validator.NewValidator()

	postIDval, err := val.CheckAnyData("post id", 36, postID, true)
	if err != nil {
		return err
	}
	titleVal, err := val.CheckAnyData("titulo", 255, title, true)
	if err != nil {
		return err
	}
	contentVal, err := val.CheckAnyData("conteudo", 2024, content, true)
	if err != nil {
		return err
	}

	commentID := uuid.New()
	commentEntity := new(models.Comment)
	commentEntity.CommentID = commentID.String()
	commentEntity.Title = titleVal.(string)
	commentEntity.Content = contentVal.(string)
	commentEntity.UserID = s.userID
	commentEntity.PostID = postIDval.(string)

	repComment := repository.NewCommentRepository()
	err = repComment.Store(commentEntity)
	if err != nil {
		return err
	}

	systemService := NewSystemService()
	err = systemService.SendEmailComment(commentID.String())
	if err != nil {
		return err
	}

	return nil
}

func (s *commentServiceImpl) ListCommentsPost(postID string, offset, limit, page int) ([]models.Comment, int, error) {
	val := validator.NewValidator()
	postIDval, err := val.CheckAnyData("id da postagem", 36, postID, true)
	if err != nil {
		return nil, 0, err
	}

	repComment := repository.NewCommentRepository()
	commentsEntities, err := repComment.List(postIDval.(string), offset, limit, page)
	if err != nil {
		return nil, 0, err
	}
	count, err := repComment.Count(postIDval.(string))
	if err != nil {
		return nil, 0, err
	}

	return commentsEntities, count, nil
}

func (s *commentServiceImpl) ListCommentsUser(offset, limit, page int) ([]models.Comment, int, error) {

	val := validator.NewValidator()
	userIDval, err := val.CheckAnyData("post id", 36, s.userID, true)
	if err != nil {
		return nil, 0, err
	}

	repComment := repository.NewCommentRepository()
	commentsEntities, err := repComment.ListUser(userIDval.(string), offset, limit, page)
	if err != nil {
		return nil, 0, err
	}
	count, err := repComment.CountUser(s.userID)
	if err != nil {
		return nil, 0, err
	}

	return commentsEntities, count, nil
}

func (s *commentServiceImpl) ListCommentsPostUser(postID string, offset, limit, page int) ([]models.Comment, int, error) {

	val := validator.NewValidator()
	postIDval, err := val.CheckAnyData("post id", 36, postID, true)
	if err != nil {
		return nil, 0, err
	}
	userIDval, err := val.CheckAnyData("post id", 36, s.userID, true)
	if err != nil {
		return nil, 0, err
	}

	repComment := repository.NewCommentRepository()
	commentsEntities, err := repComment.ListUserPost(postIDval.(string), userIDval.(string), offset, limit, page)
	if err != nil {
		return nil, 0, err
	}
	count, err := repComment.CountUserPost(postIDval.(string), s.userID)
	if err != nil {
		return nil, 0, err
	}

	return commentsEntities, count, nil
}

func (s *commentServiceImpl) FindComment(commentID string) (*models.Comment, error) {
	val := validator.NewValidator()
	commentIDval, err := val.CheckAnyData("post id", 36, commentID, true)
	if err != nil {
		return nil, err
	}

	repComment := repository.NewCommentRepository()
	comment, err := repComment.Find(commentIDval.(string))
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentServiceImpl) UpdateComment(commentID, title, content string) error {
	val := validator.NewValidator()
	commentIDVal, err := val.CheckAnyData("id do comentario", 36, commentID, true)
	if err != nil {
		return err
	}
	titleVal, err := val.CheckAnyData("titulo", 255, title, true)
	if err != nil {
		return err
	}
	contentVal, err := val.CheckAnyData("conteudo", 2024, content, true)
	if err != nil {
		return err
	}

	commentEntity := new(models.Comment)
	commentEntity.CommentID = commentIDVal.(string)
	commentEntity.Title = titleVal.(string)
	commentEntity.Content = contentVal.(string)

	repComment := repository.NewCommentRepository()
	err = repComment.Update(commentEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *commentServiceImpl) RemoveComment(commentID string) error {
	val := validator.NewValidator()
	commentIDVal, err := val.CheckAnyData("id do comentario", 36, commentID, true)
	if err != nil {
		return err
	}

	repComment := repository.NewCommentRepository()
	err = repComment.Remove(commentIDVal.(string))
	if err != nil {
		return err
	}

	return nil
}

func NewCommentService(userID, kind string) commentServiceInterface {
	return &commentServiceImpl{
		userID: userID,
		kind:   kind,
	}
}
