package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	r := gin.Default()

	// Configuración de la base de datos
	dsn := "root:@tcp(127.0.0.1:3306)/crud_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	// Migrar el modelo User a la base de datos
	db.AutoMigrate(&User{})

	// Rutas para las operaciones CRUD
	r.GET("/users", GetUsers(db))
	r.GET("/users/:id", GetUser(db))
	r.POST("/users", CreateUser(db))
	r.PUT("/users/:id", UpdateUser(db))
	r.DELETE("/users/:id", DeleteUser(db))

	r.Run(":8080")
}

// Middleware para responder con información personal
func InfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":        "Diego Laris",
			"matricula":   "200607",
			"grupo":       "10B IDGS",
			"universidad": "Universidad Tecnologica de Aguascalientes",
		})
	}
}

// Middleware para manejar errores comunes
func ErrorMiddleware(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": err.Error()})
}

// Manejadores para las rutas CRUD
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, gin.H{
			"data": users,
		})
	}
}

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			ErrorMiddleware(c, err)
			return
		}
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			ErrorMiddleware(c, err)
			return
		}
		db.Create(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			ErrorMiddleware(c, err)
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			ErrorMiddleware(c, err)
			return
		}
		db.Save(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			ErrorMiddleware(c, err)
			return
		}
		db.Delete(&user)
		c.JSON(200, gin.H{
			"message": "User deleted successfully!",
		})
	}
}
