package handler

import (
	"strconv"
	"strings"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service   product.ServiceProductInterface
	validator *validator.Validate
}

// GetByCustomer implements product.HandlerProductInterface.
func (h *ProductHandler) GetByCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := "SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at FROM products WHERE 1=1 AND is_available = true"
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

		query := "SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at FROM products WHERE 1=1"
		filters := make([]interface{}, 0)

		// TODO: add logic to get id
		idClean := strings.ReplaceAll(c.QueryParam("id"), "\"", "")
		if idClean != "" {
			// isExist, _ := h.service.IsProductExists(idClean)
			// if !isExist {
			// 	return response.SendStatusOkWithDataResponse(c, "Success", &[]dto.ResponseProducts{})
			// }
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
