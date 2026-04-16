package controller

import (
	"bubble/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) { //包外要使用，要用一个首字母大写的形式
	c.HTML(http.StatusOK, "index.html", nil)
}

//这里只调用不用实现
//controller 调用logic层的函数
//url-->controller-->logic-->model
//请求-->控制器-->业务逻辑-->模型层的增删改查

func CreateATodo(c *gin.Context) {
	//如果在前端发一个添加的需求
	//1.得到请求中的数据
	//2.存入数据库
	//3.放回响应
	var todo models.Todo
	err := c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = models.CreatATodo(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "created"})
	}
}

func GetTodoList(c *gin.Context) {
	var todoList []models.Todo
	err := models.GetTodoList(&todoList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetATodo(c *gin.Context) {
	var todo models.Todo
	id1 := c.Param("id")
	id, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title":  todo.Title,
			"status": todo.Status,
		})
	}
}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id1, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "updated"})
}

func DeleteATodo(c *gin.Context) {
	id1, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "deleted"})
}
