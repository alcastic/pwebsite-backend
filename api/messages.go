package api

import (
	"net/http"

	"github.com/alcastic/pwebsite-backend/internal/generated/sqlc"
	"github.com/gin-gonic/gin"
)

type getMessagesQuery struct {
	PageNumber int32 `form:"pageNumber" binding:"required"`
	PageSize   int32 `form:"pageSize" binding:"required"`
}

func (s *Server) getMessages(ctx *gin.Context) {
	var query getMessagesQuery
	ctx.ShouldBindQuery(&query)

	params := sqlc.ListMessagesParams{
		PageOffset: query.PageNumber * query.PageSize,
		PageSize:   query.PageSize,
	}

	messages, err := s.store.ListMessages(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, &messages)
}
