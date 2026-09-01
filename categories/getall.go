package categories

import (
	"context"

	"gorm.io/gorm"
)

type GetAllRequest struct {
	repository *gorm.DB
	create
}

func (g getAll) GetAllCategories(ctx context.Context) ([]Category, error) {

}
