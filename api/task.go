package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/ithaquaKr/taskManager/db/sqlc"
)

type createTaskRequest struct {
	ListID      uuid.UUID      `json:"listID" biding:"required"`
	Title       string         `json:"title" binding:"required"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	Tag         sql.NullString `json:"tag"`
	Priority    string         `json:"priority"`
	DueDate     sql.NullTime   `json:"dueDate"`
}

func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}
	task, err := server.store.CreateTask(ctx, db.CreateTaskParams{
		ListID:      req.ListID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Tag:         req.Tag,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}

type getTaskRequest struct {
	TaskID uuid.UUID `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTask(ctx *gin.Context) {
	var req getTaskRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	task, err := server.store.GetTask(ctx, req.TaskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}

type listTaskRequest struct {
	ListID   uuid.UUID `json:"listID" binding:"required"`
	PageID   int32     `json:"page_id" binding:"required,min=1"`
	PageSize int32     `json:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTask(ctx *gin.Context) {
	var req listTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	tasks, err := server.store.ListTask(ctx, db.ListTaskParams{
		ListID: req.ListID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
