package data

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/osamikoyo/picall/internal/data/models"
)

type Storage interface {
	Get(name string) (models.Image, error)
	GetAll() ([]models.Image, error)
	Save(image models.Image) error
	Delete(name string) error
}

type ImageStorage struct {
	Bucket string
	Minio  *minio.Client
}

func New(addr string, secure bool) (*minio.Client, error) {
	return minio.New(addr, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: secure,
	})
}
