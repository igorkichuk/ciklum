package usecase

import "github.com/igorkichuk/ciklum/internal/models"

type Article interface {
	SplitResponses() (models.CiklumResponse, error)
}
