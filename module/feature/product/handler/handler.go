package handler

import (
	"strconv"
	"strings"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/helper"
	"github.com/agusheryanto182/go-inventory-management/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service         product.ServiceProductInterface
	customerService customer.ServiceCustomerInterface
	validator       *validator.Validate
}

// CheckoutProduct implements product.HandlerProductInterface.
func (h *ProductHandler) CheckoutProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			c.Logger().Error("unauthorized: missing token or invalid token")
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to get payload
		payload := new(dto.CheckoutProductRequest)
		if err := c.Bind(&payload); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to validate payload
		if err := h.validator.Struct(payload); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		if payload.Change == nil {
			c.Logger().Error("change cannot be nil")
			return response.SendBadRequestResponse(c, "change cannot be nil")
		}

		for i := 0; i < len(payload.ProductDetails); i++ {
			if err := h.validator.Struct(payload.ProductDetails[i]); err != nil {
				c.Logger().Error(err.Error())
				return response.SendBadRequestResponse(c, err.Error())
			}
		}

		// TODO: add logic to checkout product
		if err := h.service.CheckoutProduct(payload); err != nil {
			switch err.Error() {
			case "customerId is not found":
				c.Logger().Error(err.Error())
				return response.SendStatusNotFoundResponse(c, err.Error())
			case "stock not enough":
				c.Logger().Error(err.Error())
				return response.SendBadRequestResponse(c, err.Error())
			case "product not found":
				c.Logger().Error(err.Error())
				return response.SendStatusNotFoundResponse(c, err.Error())
			default:
				c.Logger().Error(err.Error())
				return response.SendBadRequestResponse(c, err.Error())
			}
		}

		// TODO: add logic to return response
		return response.SendStatusOkResponse(c, "success checkout product")
	}
}

// GetHistoryCheckout implements product.HandlerProductInterface.
func (h *ProductHandler) GetHistoryCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			c.Logger().Error("unauthorized: missing token or invalid token")
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to define query and filters
		query := "SELECT * FROM checkouts WHERE 1=1"
		filters := make([]interface{}, 0)

		// TODO: add logic to get params customer id
		customerIdCleaned := strings.ReplaceAll(c.QueryParam("customerId"), "\"", "")
		if customerIdCleaned != "" {
			query += " AND customer_id = $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, customerIdCleaned)
		}

		// TODO: add logic to get params createdAt
		createdAtCleaned := strings.ReplaceAll(c.QueryParam("createdAt"), "\"", "")
		if createdAtCleaned != "" {
			if createdAtCleaned == "asc" || createdAtCleaned == "desc" {
				query += " ORDER BY created_at " + createdAtCleaned
			}
		} else {
			query += " ORDER BY created_at DESC"
		}

		// TODO: add logic to limit and offset
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		if limit != 0 {
			query += " LIMIT $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, limit)
		} else {
			query += " LIMIT 5"
		}

		if offset != 0 {
			query += " OFFSET $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, offset)
		} else {
			query += " OFFSET 0"
		}

		// TODO: add logic to get history checkout
		histories, err := h.service.GetHistoryCheckout(query, filters)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		if len(histories) == 0 {
			return response.SendStatusOkWithDataResponse(c, "Success get history checkout", []entities.Checkout{})
		}

		return response.SendStatusOkWithDataResponse(c, "Success get history checkout", histories)
	}
}

// GetByCustomer implements product.HandlerProductInterface.
func (h *ProductHandler) GetByCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := "SELECT * FROM products WHERE 1=1 AND is_available = true"
		filters := make([]interface{}, 0)

		// TODO: add logic to limit and offset
		limit := c.QueryParam("limit")
		limitInt, _ := strconv.Atoi(limit)

		offset := c.QueryParam("offset")
		offsetInt, _ := strconv.Atoi(offset)

		// TODO: add logic to get name
		nameClean := strings.ReplaceAll(c.QueryParam("name"), "\"", "")
		if nameClean != "" {
			query += " AND name ILIKE(CONCAT('%', $" + strconv.Itoa(len(filters)+1) + "::text, '%'))"
			filters = append(filters, nameClean)
		}

		// TODO: add logic to get category param
		categoryClean := strings.ReplaceAll(c.QueryParam("category"), "\"", "")
		if categoryClean == "Clothing" || categoryClean == "Accessories" || categoryClean == "Footwear" || categoryClean == "Beverages" {
			query += " AND category = $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, categoryClean)
		}

		// TODO: add logic to get sku param
		skuClean := strings.ReplaceAll(c.QueryParam("sku"), "\"", "")
		if skuClean != "" {
			// isExist, _ := h.service.IsSkuExists(sku)
			// if !isExist {
			// 	return response.SendStatusOkWithDataResponse(c, "Success", &[]dto.ResponseProducts{})
			// }
			query += " AND sku = $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, skuClean)
		}

		// TODO: add logic to get inStock param
		inStock := c.QueryParam("inStock")
		if inStock == "true" {
			query += " AND stock > 0"
		} else if inStock == "false" {
			query += " AND stock = 0"
		}

		// TODO: add logic to get price param
		priceClean := strings.ReplaceAll(c.QueryParam("price"), "\"", "")
		if priceClean == "asc" || priceClean == "desc" {
			query += " ORDER BY price " + priceClean
		}

		// TODO: add logic to get createdAt param
		if priceClean == "" {
			query += " ORDER BY created_at DESC"
		} else if priceClean != "" {
			query += ", created_at DESC"
		}

		if limit == "" {
			limitInt = 5
			query += " LIMIT $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, limitInt)
		} else {
			query += " LIMIT $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, limitInt)
		}

		if offset == "" {
			offsetInt = 0
			query += " OFFSET $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, offsetInt)
		} else {
			query += " OFFSET $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, offsetInt)
		}

		// TODO: add logic to get product
		product, err := h.service.GetByCustomer(query, filters)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendStatusInternalServerResponse(c, err.Error())
		}

		if len(product) == 0 {
			return response.SendStatusOkWithDataResponse(c, "success", &[]dto.ResponseProducts{})
		}

		return response.SendStatusOkWithDataResponse(c, "success", product)
	}
}

// Create implements product.HandlerProductInterface.
func (h *ProductHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to bind request
		product := new(dto.RequestCreateAndUpdateProduct)
		if err := c.Bind(&product); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(product); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		if product.Stock == nil {
			c.Logger().Error("stock cannot be empty")
			return response.SendBadRequestResponse(c, "stock cannot be empty")
		}

		if product.IsAvailable == nil {
			c.Logger().Error("isAvailable cannot be empty")
			return response.SendBadRequestResponse(c, "isAvailable cannot be empty")
		}

		if !helper.IsValidURL(product.ImageURL) {
			c.Logger().Error("invalid image url")
			return response.SendBadRequestResponse(c, "invalid image url")
		}

		// TODO: add logic to create product
		createdProduct, err := h.service.Create(product)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response
		return response.SendStatusCreatedResponse(c, "Product successfully created", createdProduct)
	}
}

// Delete implements product.HandlerProductInterface.
func (h *ProductHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to get id
		id := c.Param("id")

		// TODO: add logic to check product
		isExist, _ := h.service.IsProductExists(id)
		if !isExist {
			c.Logger().Error("Product not found")
			return response.SendStatusNotFoundResponse(c, "Product not found")
		}

		// TODO: add logic to delete product
		if err := h.service.Delete(id); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response
		return response.SendStatusOkResponse(c, "Product successfully deleted")
	}
}

// GetByParams implements product.HandlerProductInterface.
func (h *ProductHandler) GetProductByFilters() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		query := "SELECT * FROM products WHERE 1=1"
		filters := make([]interface{}, 0)

		// TODO: add logic to get id
		idClean := strings.ReplaceAll(c.QueryParam("id"), "\"", "")
		if idClean != "" {
			query += " AND id = $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, idClean)
		}

		// TODO: add logic to limit and offset
		limit := c.QueryParam("limit")
		limitInt, _ := strconv.Atoi(limit)

		offset := c.QueryParam("offset")
		offsetInt, _ := strconv.Atoi(offset)

		// TODO: add logic to get name
		nameClean := strings.ReplaceAll(c.QueryParam("name"), "\"", "")
		if nameClean != "" {
			query += " AND name ILIKE (CONCAT('%', $" + strconv.Itoa(len(filters)+1) + "::text, '%'))"
			filters = append(filters, nameClean)
		}

		// TODO: add logic to get isAvailable param
		isAvailable := c.QueryParam("isAvailable")
		if isAvailable == "true" || isAvailable == "false" {
			query += " AND is_available = $" + strconv.Itoa(len(filters)+1)
			isAvailableBool, _ := strconv.ParseBool(isAvailable)
			filters = append(filters, isAvailableBool)
		}

		// TODO: add logic to get category param
		categoryClean := strings.ReplaceAll(c.QueryParam("category"), "\"", "")
		if categoryClean == "Clothing" || categoryClean == "Accessories" || categoryClean == "Footwear" || categoryClean == "Beverages" {
			query += " AND category = $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, categoryClean)
		}

		// TODO: add logic to get sku param
		skuClean := strings.ReplaceAll(c.QueryParam("sku"), "\"", "")
		if skuClean != "" {
			query += " AND sku = $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, skuClean)
		}

		// TODO: add logic to get inStock param
		inStock := c.QueryParam("inStock")
		if inStock == "true" {
			query += " AND stock > 0"
		} else if inStock == "false" {
			query += " AND stock = 0"
		}

		// TODO: add logic to get price param
		priceClean := strings.ReplaceAll(c.QueryParam("price"), "\"", "")
		if priceClean == "asc" || priceClean == "desc" {
			query += " ORDER BY price " + priceClean
		}

		// TODO: add logic to get createdAt param
		createdAtClean := strings.ReplaceAll(c.QueryParam("createdAt"), "\"", "")
		if priceClean == "" {
			if createdAtClean == "asc" || createdAtClean == "desc" {
				query += " ORDER BY created_at " + createdAtClean
			} else {
				query += " ORDER BY created_at DESC"
			}
		} else if priceClean != "" {
			if createdAtClean == "asc" || createdAtClean == "desc" {
				query += ", created_at " + createdAtClean
			} else {
				query += ", created_at DESC"
			}
		}

		if limit == "" {
			limitInt = 5
			query += " LIMIT $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, limitInt)
		} else {
			query += " LIMIT $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, limitInt)
		}

		if offset == "" {
			offsetInt = 0
			query += " OFFSET $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, offsetInt)
		} else {
			query += " OFFSET $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, offsetInt)
		}

		// TODO: add logic to get product
		product, err := h.service.GetProductByFilters(query, filters)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendStatusInternalServerResponse(c, err.Error())
		}

		if len(product) == 0 {
			return response.SendStatusOkWithDataResponse(c, "success", &[]dto.ResponseProducts{})
		}

		return response.SendStatusOkWithDataResponse(c, "success", product)
	}
}

// Update implements product.HandlerProductInterface.
func (h *ProductHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to get id
		id := c.Param("id")

		// TODO: add logic to check product
		isExist, _ := h.service.IsProductExists(id)
		if !isExist {
			c.Logger().Error("Product not found")
			return response.SendStatusNotFoundResponse(c, "Product not found")
		}

		// TODO: add logic to bind request
		updateRequest := new(dto.RequestCreateAndUpdateProduct)
		if err := c.Bind(&updateRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(updateRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		if updateRequest.IsAvailable == nil {
			c.Logger().Error("IsAvailable is required")
			return response.SendBadRequestResponse(c, "IsAvailable is required")
		}

		if updateRequest.Stock == nil {
			c.Logger().Error("Stock is required")
			return response.SendBadRequestResponse(c, "Stock is required")
		}

		if !helper.IsValidURL(updateRequest.ImageURL) {
			c.Logger().Error("invalid image url")
			return response.SendBadRequestResponse(c, "invalid image url")
		}

		updateRequest.ID = id

		// TODO: add logic to update product
		if err := h.service.Update(updateRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response
		return response.SendStatusOkResponse(c, "Product successfully updated")
	}
}

func NewProductHandler(service product.ServiceProductInterface, validator *validator.Validate) product.HandlerProductInterface {
	return &ProductHandler{
		service:   service,
		validator: validator,
	}
}
