package usecase

import (
	"assignment/internal/domain/entity"
	"assignment/internal/domain/errors"
	"context"
	"testing"
)

// --- Mock Repository ---
// すべてのインターフェースメソッドを定義する必要があります
type MockItemRepository struct {
	GetByIDFunc func(ctx context.Context, id int) (*entity.Item, error)
	UpdateFunc  func(ctx context.Context, id int, req *entity.UpdateItemRequest) error
}

// 既存のメソッド
func (m *MockItemRepository) GetByID(ctx context.Context, id int) (*entity.Item, error) {
	return m.GetByIDFunc(ctx, id)
}
func (m *MockItemRepository) Update(ctx context.Context, id int, req *entity.UpdateItemRequest) error {
	return m.UpdateFunc(ctx, id, req)
}

// ★ 足りないメソッドをダミーで追加 (これでインターフェースを満たします)
func (m *MockItemRepository) FindAll(ctx context.Context) ([]entity.Item, error) {
	return nil, nil
}
func (m *MockItemRepository) Create(ctx context.Context, item *entity.Item) error {
	return nil
}
func (m *MockItemRepository) Delete(ctx context.Context, id int) error {
	return nil
}
func (m *MockItemRepository) GetSummary(ctx context.Context) (map[string]int, error) {
	return nil, nil
}

// --- Test Suite ---
func TestUpdateItem(t *testing.T) {
	ctx := context.Background()

	t.Run("【正常系】存在するアイテムの更新に成功すること", func(t *testing.T) {
		mockRepo := &MockItemRepository{
			GetByIDFunc: func(ctx context.Context, id int) (*entity.Item, error) {
				return &entity.Item{ID: id, Name: "Before Name"}, nil
			},
			UpdateFunc: func(ctx context.Context, id int, req *entity.UpdateItemRequest) error {
				return nil
			},
		}
		u := NewItemUsecase(mockRepo) // これでエラーが消えます

		newName := "After Name"
		req := &entity.UpdateItemRequest{Name: &newName}

		err := u.UpdateItem(ctx, 1, req)
		if err != nil {
			t.Errorf("想定外のエラーが発生しました: %v", err)
		}
	})

	t.Run("【異常系】アイテムが存在しない場合、ErrItemNotFoundを返すこと", func(t *testing.T) {
		mockRepo := &MockItemRepository{
			GetByIDFunc: func(ctx context.Context, id int) (*entity.Item, error) {
				return nil, nil
			},
		}
		u := NewItemUsecase(mockRepo)

		err := u.UpdateItem(ctx, 999, &entity.UpdateItemRequest{})
		if err != errors.ErrItemNotFound {
			t.Errorf("期待値: %v, 実際の結果: %v", errors.ErrItemNotFound, err)
		}
	})
}
