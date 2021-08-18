package repository

import "github.com/alochym01/gin-website/models"

type AlbumRepo interface {
	Get() ([]models.Album, error)
	GetByID(id string) (models.Album, error)
	Create(title string, artist string, price float64) error
	Update(title string, artist string, price float64, id string) error
	Delete(id string) error
}
