package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"http-type-one/models"
	"net/http"
	"strconv"
)

/* Accounts */

func (h *Handler) CreateAccounts(c *gin.Context) {
	var input models.Accounts

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
	}

	id, err := h.services.Accounts.CreateAccounts(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Id":     id,
		"Status": "Created",
	})
}

type getAllAccountsResponse struct {
	Data []models.Accounts `json:"data"`
}

func (h *Handler) GetAllAccounts(c *gin.Context) {
	all := h.services.Accounts.GetAll()

	if len(all) == 0 {
		c.String(http.StatusNotFound, "Don't have any accounts")
	}

	c.JSON(http.StatusOK, getAllAccountsResponse{
		Data: all,
	})
}

func (h *Handler) GetAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	account := h.services.Accounts.GetAccount(id)

	c.JSON(http.StatusOK, account)
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	var input models.Accounts

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
	}

	id, err := h.services.Accounts.UpdateAccount(input)
	if err != nil {
		c.String(http.StatusNotFound, "Don't have such account")
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"Id":     id,
			"Status": "Updated",
		})
	}
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	account, err := h.services.Accounts.DeleteAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, account)
}

/* Account Integration */

func (h *Handler) CreateAccountIntegration(c *gin.Context) {
	c.String(404, "not ready")
}

func (h *Handler) GetAccountIntegration(c *gin.Context) {
	c.String(404, "not ready")
}

func (h *Handler) GetAllAccountIntegration(c *gin.Context) {
	c.String(404, "not ready")
}

func (h *Handler) UpdateAccountIntegration(c *gin.Context) {
	c.String(404, "not ready")
}

func (h *Handler) DeleteAccountIntegration(c *gin.Context) {
	c.String(404, "not ready")
}
