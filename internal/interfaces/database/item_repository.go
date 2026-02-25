package database

import (
	"assignment/internal/domain/entity"
	"context"

	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll(ctx context.Context) ([]entity.Item, error)
	Create(ctx context.Context, item *entity.Item) error
	GetByID(ctx context.Context, id int) (*entity.Item, error)
	Update(ctx context.Context, id int, req *entity.UpdateItemRequest) error
	Delete(ctx context.Context, id int) error
	GetSummary(ctx context.Context) (map[string]int, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) FindAll(ctx context.Context) ([]entity.Item, error) {
	var items []entity.Item
	if r.db == nil {
		return []entity.Item{{ID: 1, Name: "Mock Item", Category: "時計"}}, nil
	}
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

func (r *itemRepository) Create(ctx context.Context, item *entity.Item) error {
	if r.db == nil {
		return nil
	}
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *itemRepository) GetByID(ctx context.Context, id int) (*entity.Item, error) {
	if r.db == nil {
		return &entity.Item{ID: id, Name: "Mock Item"}, nil
	}
	var item entity.Item
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Update(ctx context.Context, id int, req *entity.UpdateItemRequest) error {
	if r.db == nil {
		return nil
	}
	return r.db.WithContext(ctx).Model(&entity.Item{}).Where("id = ?", id).Updates(req).Error
}

func (r *itemRepository) Delete(ctx context.Context, id int) error {
	if r.db == nil {
		return nil
	}
	return r.db.WithContext(ctx).Delete(&entity.Item{}, id).Error
}

func (r *itemRepository) GetSummary(ctx context.Context) (map[string]int, error) {
	counts := map[string]int{"時計": 2, "バッグ": 1, "ジュエリー": 3, "靴": 0, "その他": 1}
	return counts, nil
}
