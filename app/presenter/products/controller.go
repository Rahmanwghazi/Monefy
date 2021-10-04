package products

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/products/responses"
	"github.com/Rahmanwghazi/Monefy/business/products"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductUsecase products.Usecase
}

func NewInvestPlanController(productUsecase products.Usecase) *ProductController {
	return &ProductController{
		ProductUsecase: productUsecase,
	}
}

func (productController ProductController) GetProducts(echoContext echo.Context) error {
	result, err := productController.ProductUsecase.SearchProduct(products.ProductDomain{})
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}
