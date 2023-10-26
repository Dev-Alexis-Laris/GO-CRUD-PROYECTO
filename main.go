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

	r.Use(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":        "Diego Laris",
			"matricula":   "200607",
			"grupo":       "10B IDGS",
			"universidad": "Universidad Tecnologica de Aguascalientes",
		})
	})

	dsn := "root:@tcp(127.0.0.1:3306)/crud_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)

		c.JSON(200, gin.H{
			"data": users,
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		var user User

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Create(&user)

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var user User

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Save(&user)

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		var user User

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&user)

		c.JSON(200, gin.H{
			"message": "User deleted successfully!",
		})
	})

	r.Run(":8080")
}
