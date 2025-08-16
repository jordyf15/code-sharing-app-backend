package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordyf15/code-sharing-app/snippet"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SnippetController interface {
	CreateSnippet(c *gin.Context)
	UpdateSnippet(c *gin.Context)
	GetSnippet(c *gin.Context)
}

type snippetController struct {
	usecase snippet.Usecase
}

func NewSnippetController(usecase snippet.Usecase) SnippetController {
	return &snippetController{usecase: usecase}
}

func (controller *snippetController) CreateSnippet(c *gin.Context) {
	snippetCode := c.PostForm("code")
	snippetLanguage := c.PostForm("language")
	snippetTheme := c.PostForm("theme")

	snippet, err := controller.usecase.Create(snippetCode, snippetLanguage, snippetTheme)
	if err != nil {
		respondBasedOnError(c, err)
		return
	}

	c.JSON(http.StatusOK, snippet)
}

func (controller *snippetController) UpdateSnippet(c *gin.Context) {
	idStr := c.Param("snippet_id")
	snippetId, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		respondBasedOnError(c, err)
		return
	}

	updates := map[string]interface{}{}

	code, isExist := c.GetPostForm("code")
	if isExist {
		updates["code"] = code
	}

	language, isExist := c.GetPostForm("language")
	if isExist {
		updates["language"] = language
	}

	theme, isExist := c.GetPostForm("theme")
	if isExist {
		updates["theme"] = theme
	}

	snippet, err := controller.usecase.UpdateSnippet(snippetId, updates)
	if err != nil {
		respondBasedOnError(c, err)
		return
	}

	c.JSON(http.StatusOK, snippet)
}

func (controller *snippetController) GetSnippet(c *gin.Context) {
	idStr := c.Param("snippet_id")
	snippetId, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		respondBasedOnError(c, err)
		return
	}

	snippet, err := controller.usecase.GetSnippet(snippetId)
	if err != nil {
		respondBasedOnError(c, err)
		return
	}

	c.JSON(http.StatusOK, snippet)
}
