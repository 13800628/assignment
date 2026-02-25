package usecase

import (
	"assignment/internal/domain/entity"
	"assignment/internal/domain/errors"
	"assignment/internal/interfaces/database"
	"context"
)

type ItemUsecase interface {
	FindAll(ctx context.Context) ([]entity.Item, error)
	Create(ctx context.Context, item *entity.Item) error
	FindByID(ctx context.Context, id int) (*entity.Item, error)
	Delete(ctx context.Context, id int) error
	GetSummary(ctx context.Context) (map[string]interface{}, error)
	UpdateItem(ctx context.Context, id int, req *entity.UpdateItemRequest) error
}

type itemUsecase struct {
	repo database.ItemRepository
}

func NewItemUsecase(repo database.ItemRepository) ItemUsecase {
	return &itemUsecase{repo: repo}
}

func (u *itemUsecase) FindAll(ctx context.Context) ([]entity.Item, error) {
	return u.repo.FindAll(ctx)
}

func (u *itemUsecase) Create(ctx context.Context, item *entity.Item) error {
	return u.repo.Create(ctx, item)
}

func (u *itemUsecase) FindByID(ctx context.Context, id int) (*entity.Item, error) {
	item, err := u.repo.GetByID(ctx, id)
	if err != nil || item == nil {
		return nil, errors.ErrItemNotFound
	}
	return item, nil
}

func (u *itemUsecase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}

func (u *itemUsecase) GetSummary(ctx context.Context) (map[string]interface{}, error) {
	counts, err := u.repo.GetSummary(ctx)
	if err != nil {
		return nil, err
	}
	total := 0
	for _, c := range counts {
		total += c
	}
	return map[string]interface{}{"categories": counts, "total": total}, nil
}

func (u *itemUsecase) UpdateItem(ctx context.Context, id int, req *entity.UpdateItemRequest) error {
	item, err := u.repo.GetByID(ctx, id)
	if err != nil || item == nil {
		return errors.ErrItemNotFound
	}
	return u.repo.Update(ctx, id, req)
}
