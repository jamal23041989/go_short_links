package link

import (
	"github.com/jamal23041989/go_short_links/pkg/db"
)

type LinkRepository struct {
	Database *db.Db
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}
