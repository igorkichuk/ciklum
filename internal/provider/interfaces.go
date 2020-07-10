package provider

import "github.com/igorkichuk/ciklum/internal/models"

type CiklumProvider interface {
	GetResponse(url string) (models.CiklumResponse, error)
}
