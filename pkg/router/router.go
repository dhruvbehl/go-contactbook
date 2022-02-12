package router

import (
	"net/http"

	"github.com/dhruvbehl/address-book/api"
	"github.com/dhruvbehl/address-book/pkg/contact"
	"github.com/gin-gonic/gin"
)

func GetAllContactHandler(ctx *gin.Context) {
	response, err := contact.GetAllContact()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func GetContactByIdHandler(ctx *gin.Context) {
	id := ctx.Param(("id"))
	response, err := contact.GetContactById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func DeleteContactByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := contact.DeleteContactById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func CreateContactHandler(ctx *gin.Context) {
	var c api.Contact
	if err := ctx.Bind(&c); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := contact.UpsertContact(c, "")
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": response})
}

func UpdateContactHandler(ctx *gin.Context) {
	var c api.Contact
	id := ctx.Param("id")
	if err := ctx.Bind(&c); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := contact.UpsertContact(c, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": response})
}
