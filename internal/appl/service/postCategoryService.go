package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/johnHPX/blog-hard-backend/internal/domain/models"
	"github.com/johnHPX/blog-hard-backend/internal/infra/repository"
	"github.com/johnHPX/blog-hard-backend/internal/infra/utils/messages"
	"github.com/johnHPX/validator-hard/pkg/validator"
)

type postCategoryServiceInterface interface {
	StorePostCategory(postID, categoryID string) error
	RemovePostCategory(postID, categoryID string) error
}

type postCategoryServiceImpl struct {
	userID string
	kind   string
}

func (s *postCategoryServiceImpl) StorePostCategory(postID, categoryID string) error {

	if s.kind != "adm" {
		return errors.New(messages.AdmMessage)
	}

	val := validator.NewValidator()
	postIDval, err := val.CheckAnyData("id da postagem", 36, postID, true)
	if err != nil {
		return err
	}
	categoryIDval, err := val.CheckAnyData("id da categoria", 36, categoryID, true)
	if err != nil {
		return err
	}

	postCatergoryID := uuid.New()

	postCategoryEntity := new(models.PostCategory)
	postCategoryEntity.PostCategoryId = postCatergoryID.String()
	postCategoryEntity.PostId = postIDval.(string)
	postCategoryEntity.CategoryId = categoryIDval.(string)

	repPostCategory := repository.NewPostCategoryRepository()

	_, err = repPostCategory.Find(postIDval.(string), categoryIDval.(string))
	if err == nil {
		return errors.New("esse cadastro já foi realizado!")
	}

	err = repPostCategory.Store(postCategoryEntity)
	if err != nil {
		return err
	}

	return nil

}

func (s *postCategoryServiceImpl) RemovePostCategory(postID, categoryID string) error {

	if s.kind != "adm" {
		return errors.New(messages.AdmMessage)
	}

	val := validator.NewValidator()
	postIDval, err := val.CheckAnyData("id da postagem", 36, postID, true)
	if err != nil {
		return err
	}
	categoryIDval, err := val.CheckAnyData("id da categoria", 36, categoryID, true)
	if err != nil {
		return err
	}

	repPostCategory := repository.NewPostCategoryRepository()

	_, err = repPostCategory.Find(postIDval.(string), categoryIDval.(string))
	if err != nil {
		return errors.New("já foi deletetado ou não existe")
	}

	err = repPostCategory.Remove(postIDval.(string), categoryIDval.(string))
	if err != nil {
		return err
	}

	return nil
}

func NewPostCategoryService(userID, kind string) postCategoryServiceInterface {
	return &postCategoryServiceImpl{
		userID: userID,
		kind:   kind,
	}
}
