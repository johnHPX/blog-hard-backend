package repository

import (
	"database/sql"
	"errors"

	"github.com/johnHPX/blog-hard-backend/internal/domain/models"
	"github.com/johnHPX/blog-hard-backend/internal/infra/utils/databaseConn"
	"github.com/johnHPX/blog-hard-backend/internal/infra/utils/messages"
)

type postCategoryRepositoryInterface interface {
	Store(entity *models.PostCategory) error
	Find(postID string, categoryID string) (*models.PostCategory, error)
	Update(entity *models.PostCategory) error
	Remove(postID, categoryID string) error
}

type postCategoryRepositoryImpl struct{}

func (r *postCategoryRepositoryImpl) scanIterator(rows *sql.Rows) (*models.PostCategory, error) {
	id := sql.NullString{}
	postID := sql.NullString{}
	categoryID := sql.NullString{}

	err := rows.Scan(
		&id,
		&postID,
		&categoryID,
	)

	if err != nil {
		return nil, err
	}

	postCategoryEntity := new(models.PostCategory)

	if id.Valid {
		postCategoryEntity.PostCategoryId = id.String
	}

	if postID.Valid {
		postCategoryEntity.PostId = postID.String
	}

	if categoryID.Valid {
		postCategoryEntity.CategoryId = categoryID.String
	}

	return postCategoryEntity, nil
}

func (r *postCategoryRepositoryImpl) Store(entity *models.PostCategory) error {
	db, err := databaseConn.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlText := `INSERT INTO tb_post_category 
		(id, post_pid, category_cid)
		VALUES
		($1, $2, $3)
	 `
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}
	result, err := statement.Exec(entity.PostCategoryId, entity.PostId, entity.CategoryId)
	if err != nil {
		return err
	}
	defer statement.Close()

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		return errors.New(messages.StoreError)
	}

	return nil
}

func (r *postCategoryRepositoryImpl) Find(postID string, categoryID string) (*models.PostCategory, error) {
	db, err := databaseConn.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlText := `
		SELECT 
			id,
			post_pid,
			category_cid 
		FROM tb_post_category
		WHERE deleted_at is null and post_pid = $1 and category_cid = $2
	`

	rows, err := db.Query(sqlText, postID, categoryID)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		postCategory, err := r.scanIterator(rows)
		if err != nil {
			return nil, err
		}

		return postCategory, nil
	}

	return nil, errors.New(messages.FindError)
}

func (r *postCategoryRepositoryImpl) Update(entity *models.PostCategory) error {
	db, err := databaseConn.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlText := `
	update tb_post_category set
		post_pid = $2,
		category_cid = $3,
		updated_at = now()
	where deleted_at is null and id = $1
	`

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}
	result, err := statement.Exec(entity.PostCategoryId, entity.PostId, entity.CategoryId)
	if err != nil {
		return err
	}
	defer statement.Close()

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		return errors.New(messages.UpdateError)
	}

	return nil
}

func (r *postCategoryRepositoryImpl) Remove(postID, categoryID string) error {
	db, err := databaseConn.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlText := `
	delete from tb_post_category
	where post_pid = $1 and category_cid = $2
	`
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}
	result, err := statement.Exec(postID, categoryID)
	if err != nil {
		return err
	}
	defer statement.Close()

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		return errors.New(messages.RemoveError)
	}

	return nil
}

func NewPostCategoryRepository() postCategoryRepositoryInterface {
	return &postCategoryRepositoryImpl{}
}
