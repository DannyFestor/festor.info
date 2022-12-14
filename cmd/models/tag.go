package models

import (
	"database/sql"
	"strings"
	"time"
)

type Tag struct {
	Id              int64
	Title           string
	FontColor       *string `omitempty`
	BackgroundColor *string `omitempty`
	BorderColor     *string `omitempty`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type TagService struct {
	db *sql.DB
}

func (ts *TagService) Get() ([]Tag, error) {
	query := `
SELECT id, title, font_color, background_color, border_color
FROM tags;
	`
	rows, err := ts.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var t Tag = Tag{}
		err = rows.Scan(&t.Id, &t.Title, &t.FontColor, &t.BackgroundColor, &t.BorderColor)
		if err != nil {
			return nil, err
		}

		tags = append(tags, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (ts *TagService) FindByTitle(title string) (*int, error) {
	query := `
SELECT id
FROM tags
WHERE LOWER(title) = $1
LIMIT 1;
	`
	row := ts.db.QueryRow(query, strings.ToLower(title))
	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
