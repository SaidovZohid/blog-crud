package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

type Blog struct {
	Id          int64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetBlogsQueryParam struct {
	Author string
	Title  string
	Page   int32
	Limit  int32
}

func (b *DBManager) Create(blog *Blog) (*Blog, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return nil, err
	}
	query := `
		INSERT INTO blogs(
			title,
			description,
			author
		) VALUES ($1,$2,$3)
		RETURNING id, title, description, author, created_at
	`
	row := tx.QueryRow(
		query,
		blog.Title,
		blog.Description,
		blog.Author,
	)
	var result Blog
	err = row.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (b *DBManager) GetAll(params *GetBlogsQueryParam) ([]*Blog, error) {
	var blogs []*Blog
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)
	filter := " WHERE true "
	if params.Author != "" {
		filter += " AND author ilike '%" + params.Author + "%' "
	}

	if params.Title != "" {
		filter += " AND title ilike '%" + params.Title + "%' "
	}
	query := `
		SELECT 
			id,
			title,
			description,
			author,
			created_at
		FROM blogs ` + filter + `
		ORDER BY created_at desc
		` + limit
	rows, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog
		err := rows.Scan(
			&blog.Id,
			&blog.Title,
			&blog.Description,
			&blog.Author,
			&blog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	return blogs, nil
}

func (d *DBManager) GetBlog(id int64) (*Blog, error) {
	query := `
		select 
			id,
			title,
			description,
			author,
			created_at
		from blogs where id = $1
	`
	row := d.db.QueryRow(query, id)
	var blog Blog
	err := row.Scan(
		&blog.Id,
		&blog.Title,
		&blog.Description,
		&blog.Author,
		&blog.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (d *DBManager) DeleteBlog(id int64) error {
	query := `
		delete from blogs where id = $1
	`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBManager) UpdateBlog(blog *Blog) (*Blog, error) {
	tx, err := d.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	query := `
		INSERT INTO blogs(
			title,
			description,
			author
		) VALUES ($1,$2,$3)
		RETURNING id, title, description, author, created_at
	`
	row := tx.QueryRow(
		query,
		blog.Title,
		blog.Description,
		blog.Author,
	)
	var result Blog
	err = row.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return blog, err
}
