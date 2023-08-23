package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alcastic/pwebsite-backend/internal/generated/sqlc"
	"github.com/gin-gonic/gin"
)

//TODO: business logic must be moved to a service layer

type getMessagesQuery struct {
	PageNumber int32 `form:"pageNumber" binding:"required,min=1"`
	PageSize   int32 `form:"pageSize" binding:"required,min=1,max=20"`
}

func (s *Server) getMessages(ctx *gin.Context) {
	var query getMessagesQuery
	err := ctx.ShouldBind(&query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	params := sqlc.ListMessagesParams{
		PageOffset: (query.PageNumber - 1) * query.PageSize,
		PageSize:   query.PageSize,
	}

	messages, err := s.store.ListMessages(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, &messages)
}

type messageBody struct {
	Content     string `json:"content"`
	AuthorName  string `json:"authorName"`
	AuthorEmail string `json:"authorEmail"`
}

const minTimeBetweenMessages = 30 * time.Minute

func previousMessageRestriction(msg *sqlc.Message) bool {
	ts := time.Since(msg.CreatedAt)
	log.Printf("time since last message: %v", ts)
	return ts < minTimeBetweenMessages
}

func (s *Server) addMessage(ctx *gin.Context) {
	var body messageBody
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	clientIP := ctx.ClientIP()
	lastMsg, err := s.store.GetLastMessageFromRemoteAddr(ctx, clientIP)
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err == nil && previousMessageRestriction(lastMsg) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("cooldown time needed to post a new message")))
		return
	}

	_, err = s.store.CreateMessage(ctx, sqlc.CreateMessageParams{
		RemoteAddr:  clientIP,
		Content:     body.Content,
		AuthorName:  body.AuthorName,
		AuthorEmail: body.AuthorEmail,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, "ok")
}
