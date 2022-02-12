package router

import (
	"net/http"

	"github.com/dhruvbehl/address-book/api"
	"github.com/dhruvbehl/address-book/pkg/contact"
	"github.com/gin-gonic/gin"
)

// GetAllContact godoc
// @id GetAllContact
// @Summary Show all the contacts stored on the server
// @Description Show all the contacts stored on the server
// @Accept */*
// @Produce json
// @Success 200 {array} map[string]api.Contact{}
// Failure 404 {object} map[string]interface{} "Not Found"
// @Router /contacts [get]
func GetAllContactHandler(ctx *gin.Context) {
	response, err := contact.GetAllContact()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// GetContactById godoc
// @id GetContactById
// @Summary Show contact by id
// @Description Show contact by id
// @Param id path string true "Contact ID"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]api.Contact{}
// Failure 404 {object} map[string]interface{} "Not Found"
// @Router /contact/{id} [get]
func GetContactByIdHandler(ctx *gin.Context) {
	id := ctx.Param(("id"))
	response, err := contact.GetContactById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// GetContactById godoc
// @id DeleteContactById
// @Summary Delete contact by id
// @Description Delete contact by id "Contact ID"
// @Param id path string true "Contact ID"
// @Accept */*
// @Produce json
// @Success 200
// Failure 404 {object} map[string]interface{} "Not Found"
// @Router /contact/{id} [delete]
func DeleteContactByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := contact.DeleteContactById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// CreateContact godoc
// @id CreateContact
// @Summary Create contact
// @Description Create contact
// @Param contactRequest body api.Contact true "Contact Request Body"
// @Accept json
// @Produce json
// @Success 201 {object} map[string]api.Contact{}
// Failure 409 {object} map[string]interface{} Conflict
// Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /contact [post]
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

// UpdateContact godoc
// @id UpdateContact
// @Summary Update contact
// @Description Update contact
// @Param id path string true "Contact ID"
// @Param contactRequest body api.Contact true "Contact Request Body"
// @Accept json
// @Produce json
// @Success 201 {object} map[string]api.Contact{}
// Failure 500 {object} map[string]interface{} "Internal Server Error"
// Failure 404 {object} map[string]interface{} "Not Found"
// @Router /contact/{id} [put]
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
