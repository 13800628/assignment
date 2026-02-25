package controller

import (
	"assignment/internal/domain/entity"
	domainErrors "assignment/internal/domain/errors"
	"assignment/internal/usecase" // これを使うように定義を直します
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ItemController の定義を復活（undefined の解消）
type ItemController struct {
	usecase usecase.ItemUsecase
}

// NewItemController の定義（usecase パッケージを使用）
func NewItemController(u usecase.ItemUsecase) *ItemController {
	return &ItemController{usecase: u}
}

func (c *ItemController) PatchItem(ec echo.Context) error {
	id, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusBadRequest, domainErrors.AppError{Error: "invalid ID"})
	}

	var req entity.UpdateItemRequest
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, domainErrors.AppError{Error: "invalid body"})
	}

	// 1. バリデーション実行
	if details := req.Validate(); len(details) > 0 {
		return ec.JSON(http.StatusBadRequest, domainErrors.AppError{
			Error:   domainErrors.ErrInvalidInput.Error(),
			Details: details,
		})
	}

	// 2. ユースケース実行
	if err := c.usecase.UpdateItem(ec.Request().Context(), id, &req); err != nil {
		// 3. エラーハンドリング
		if err == domainErrors.ErrItemNotFound {
			return ec.JSON(http.StatusNotFound, domainErrors.AppError{Error: "item not found"})
		}
		return ec.JSON(http.StatusInternalServerError, domainErrors.AppError{Error: "internal server error"})
	}

	return ec.NoContent(http.StatusOK)
}

// 全取得
func (c *ItemController) GetItems(ec echo.Context) error {
	items, _ := c.usecase.FindAll(ec.Request().Context())
	return ec.JSON(http.StatusOK, items)
}

// 登録
func (c *ItemController) CreateItem(ec echo.Context) error {
	var item entity.Item
	if err := ec.Bind(&item); err != nil {
		return ec.JSON(http.StatusBadRequest, domainErrors.AppError{Error: "invalid body"})
	}
	// TODO: 登録バリデーション
	if err := c.usecase.Create(ec.Request().Context(), &item); err != nil {
		return ec.JSON(http.StatusBadRequest, domainErrors.AppError{Error: "create failed"})
	}
	return ec.JSON(http.StatusCreated, item)
}

// 特定取得
func (c *ItemController) GetItemByID(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	item, err := c.usecase.FindByID(ec.Request().Context(), id)
	if err != nil {
		return ec.JSON(http.StatusNotFound, domainErrors.AppError{Error: "not found"})
	}
	return ec.JSON(http.StatusOK, item)
}

// 削除
func (c *ItemController) DeleteItem(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	if err := c.usecase.Delete(ec.Request().Context(), id); err != nil {
		return ec.JSON(http.StatusNotFound, domainErrors.AppError{Error: "not found"})
	}
	return ec.NoContent(http.StatusNoContent)
}

// サマリー (集計)
func (c *ItemController) GetSummary(ec echo.Context) error {
	summary, _ := c.usecase.GetSummary(ec.Request().Context())
	return ec.JSON(http.StatusOK, summary)
}
