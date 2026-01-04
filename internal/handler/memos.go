package handler

import (
	"net/http"
	"strconv"

	"cap-api/internal/service"

	"github.com/gin-gonic/gin"
)

type MemoHandler struct {
	svc *service.MemoService
}

func NewMemoHandler(svc *service.MemoService) *MemoHandler {
	return &MemoHandler{svc: svc}
}

type memoReq struct {
	Content string `json:"content" binding:"required"`
}

func (h *MemoHandler) Register(r *gin.Engine) {
	r.POST("/memos", h.create)
	r.GET("/memos/:id", h.get)
	r.PUT("/memos/:id", h.update)
	r.DELETE("/memos/:id", h.delete)
	r.GET("/memos", h.list)
}

func (h *MemoHandler) create(c *gin.Context) {
	var req memoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := h.svc.Create(c.Request.Context(), req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, m)
}

func (h *MemoHandler) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	m, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "memo not found"})
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *MemoHandler) update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req memoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := h.svc.Update(c.Request.Context(), id, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *MemoHandler) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MemoHandler) list(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	q := c.Query("q")

	memos, err := h.svc.List(c.Request.Context(), limit, offset, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memos)
}
