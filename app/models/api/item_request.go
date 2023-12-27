package models

type ItemCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type ItemUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
