package routes

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"study_go/models"
)

// UserController 定义用户控制器
type UserController struct {
	db *gorm.DB
}

// NewUserController 创建用户控制器实例
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

// homeHandler 处理主页请求
func (c *UserController) homeHandler(ctx *gin.Context) {
	// 渲染HTML模板
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// CreateUser 创建用户
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ShouldBindJSON error": err.Error()})
		return
	}

	if err := c.db.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// GetUserList 获取用户列表
func (c *UserController) GetUserList(ctx *gin.Context) {
	var users []models.User
	if err := c.db.Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUser 获取用户详情
func (c *UserController) GetUser(ctx *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := c.db.First(&user, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// UpdateUser 更新用户
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := c.db.First(&user, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser 删除用户
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := c.db.Delete(&models.User{}, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Setup 设置路由
func Setup(router *gin.Engine, db *gorm.DB) {
	userController := NewUserController(db)

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", userController.homeHandler)            // 主页路由
	router.POST("/users", userController.CreateUser)       // 创建用户路由
	router.GET("/users", userController.GetUserList)       // 获取用户列表路由
	router.GET("/users/:id", userController.GetUser)       // 获取用户详情路由
	router.PUT("/users/:id", userController.UpdateUser)    // 更新用户路由
	router.DELETE("/users/:id", userController.DeleteUser) // 删除用户路由
}
