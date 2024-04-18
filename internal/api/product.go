package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zaenalarifin12/test-ecommerce/internal/domain"
	"github.com/zaenalarifin12/test-ecommerce/internal/dto/product"
	"github.com/zaenalarifin12/test-ecommerce/internal/utils"
	"net/http"
	"strconv"
)

type productAPI struct {
	productService domain.ProductService
}

func NewProduct(router *gin.Engine, productService domain.ProductService, middlewares ...gin.HandlerFunc) {
	handler := productAPI{productService: productService}

	v1 := router.Group("/api/v1")
	v1.Use(middlewares...)
	{
		v1.POST("/products", handler.CreateProduct)
		v1.GET("/products/:id", handler.GetProduct)
		v1.PUT("/products/:id", handler.UpdateProduct)
		v1.DELETE("/products/:id", utils.AuthMiddleware, handler.DeleteProduct)
		v1.GET("/products", handler.ListProduct)
	}
}
func (p *productAPI) CreateProduct(ctx *gin.Context) {
	// Parse request body
	var req product.ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(req)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	// Call ProductService to create product
	productId, err := p.productService.CreateProduct(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Call ProductService to get product by ID
	//product, err := p.productService.GetProduct(ctx, productId)
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	//	return
	//}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "data": productId})
}

func (p *productAPI) GetProduct(ctx *gin.Context) {
	// Get product ID from URL parameter
	id := ctx.Param("id")

	// Convert ID to int64
	productId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	fmt.Println(productId)
	// Call ProductService to get product by ID

	product, err := p.productService.GetProduct(ctx, int32(productId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productAPI) UpdateProduct(ctx *gin.Context) {
	// Get product ID from URL parameter
	id := ctx.Param("id")

	// Convert ID to int64
	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Parse request body
	var req product.ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call ProductService to update product
	err = p.productService.UpdateProduct(ctx, &req, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (p *productAPI) DeleteProduct(ctx *gin.Context) {
	// Get product ID from URL parameter
	id := ctx.Param("id")

	// Convert ID to int64
	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	_, err = p.productService.GetProduct(ctx, int32(productId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	// Call ProductService to delete product
	err = p.productService.DeleteProduct(ctx, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (p *productAPI) ListProduct(ctx *gin.Context) {
	// Call ProductService to list all products
	products, err := p.productService.ListProduct(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
