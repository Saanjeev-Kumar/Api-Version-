package essentials

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Controller struct {
	data map[string]Employee
}

func New() *Controller {
	return &Controller{
		data: make(map[string]Employee),
	}
}

func (c *Controller) Add(ctx *gin.Context) {
	var e Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		fmt.Println("falied to decode body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := c.data[e.Id]; ok {
		fmt.Println("id exists")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"id": e.Id, "message": "already exist"})
		return
	}
	c.data[e.Id] = e
	ctx.JSON(http.StatusCreated, gin.H{"message": "Inserted successfully", "id": e.Id, "name": e.Name})
}

func (c *Controller) Get(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, c.data)
}
