package mysql 

import (
	"database/sql"

	"lavireza.com/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return nil, nil
}

func (m *SnippetModel) Get(id int) (* models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
    return nil, nil
}