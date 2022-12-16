package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Post struct {
	Id          string
	TypeId      uint64
	Title       string
	Description string
	Image       string
	IsReleased  bool
	ReleasedAt  *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type PostService struct {
	db *sql.DB
}

// TODO: query type, pagination, limit, order by column?
func (ps *PostService) Get() ([]Post, error) {
	query := `
SELECT
	id,
	type_id,
	title,
	description,
	updated_at
FROM
	posts
WHERE
	is_released = TRUE
	AND released_at < NOW()
ORDER BY
	updated_at DESC
`

	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.TypeId, &post.Title, &post.Description, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (ps *PostService) Paginate(limit int, offset int) ([]Post, int, error) {
	var total int
	countQuery := `
	SELECT count(*)
	FROM posts
	WHERE
		is_released = TRUE
		AND released_at < NOW()
	`
	row := ps.db.QueryRow(countQuery)
	if err := row.Scan(&total); err != nil {
		return nil, -1, err
	}

	query := fmt.Sprintf(`
SELECT
	id,
	type_id,
	title,
	description,
	updated_at
FROM
	posts
WHERE
	is_released = TRUE
	AND released_at < NOW()
ORDER BY
	updated_at DESC
LIMIT %d
OFFSET %d
`, limit, offset)

	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.TypeId, &post.Title, &post.Description, &post.UpdatedAt)
		if err != nil {
			return nil, -1, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, -1, err
	}

	return posts, total, nil
}