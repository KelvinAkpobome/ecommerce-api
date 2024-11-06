package http

import (
	"ecommerce-api/application"
	"ecommerce-api/domain"
	"ecommerce-api/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *application.ProductService
}

func NewProductController(productService *application.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var input domain.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid input", Error: err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Validation failed", Error: err.Error()})
		return
	}

	if err := ctrl.ProductService.CreateProduct(&input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to create product", Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Product created successfully",
		Response: struct {
			ProductID uint `json:"product_id"`
		}{ProductID: input.ID},
	})
}

func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
	products, err := ctrl.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to fetch products", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Products fetched successfully",
		Response: products,
	})
}

func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid product ID", Error: err.Error()})
		return
	}

	product, err := ctrl.ProductService.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.Response{Success: false, Message: "Product not found", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Product found",
		Response: product,
	})
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid product ID", Error: err.Error()})
		return
	}

	var input domain.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid input", Error: err.Error()})
		return
	}

	input.ID = uint(productID)

	if err := ctrl.ProductService.UpdateProduct(&input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to update product", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Product updated successfully",
		Response: struct {
			ProductID int `json:"product_id"`
		}{ProductID: productID},
	})
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "Invalid product ID", Error: err.Error()})
		return
	}

	if err := ctrl.ProductService.DeleteProduct(uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Success: false, Message: "Failed to delete product", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Product deleted successfully",
		Response: struct {
			ProductID int `json:"product_id"`
		}{ProductID: productID},
	})
}
