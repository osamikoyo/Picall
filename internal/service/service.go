package service

import (
	"github.com/osamikoyo/picall/internal/data"
	"github.com/osamikoyo/picall/internal/data/models"
	"github.com/osamikoyo/picall/pkg/loger"
)

type Service interface {
	Get(name string) (models.Image, error)
	GetAll() ([]models.Image, error)
	Save(image models.Image) error
	Delete(name string) error
}

type ImageService struct {
	Storage data.Storage
	Logger  loger.Logger
}
