package server

import (
	"golang/storage"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetBlog(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return 
	}

	blog, err := h.storage.GetBlog(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return 
	}
	ctx.JSON(http.StatusOK, blog)
}

func (h *handler) CreateBlog(ctx *gin.Context) {
	var b storage.Blog
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return 
	}
	blog, err := h.storage.Create(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return 
	}
	ctx.JSON(http.StatusInternalServerError, blog)
}