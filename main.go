package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3" // Import the SQLite driver
    "log"
    "net/http"
)

var db *gorm.DB
var err error

// User represents the user model
type User struct {
    ID    uint   `json:"id" gorm:"primary_key"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Initialize and connect to the SQLite database
func initDB() {
    db, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatalf("Failed to connect to SQLite: %v", err)
    }
    log.Println("Connected to SQLite successfully")

    // Migrate the schema
    db.AutoMigrate(&User{})
}

func main() {
    // Initialize the database connection
    initDB()
    defer db.Close()

    // Set up the Gin router
    router := gin.Default()

    // Define API routes
    router.GET("/users", getUsers)
    router.GET("/users/:id", getUserByID)
    router.POST("/users", createUser)
    router.PUT("/users/:id", updateUser)
    router.DELETE("/users/:id", deleteUser)

    // Start the server
    router.Run(":9090")
}

// Get all users
func getUsers(c *gin.Context) {
    var users []User
    db.Find(&users)
    c.JSON(http.StatusOK, users)
}

// Get a user by ID
func getUserByID(c *gin.Context) {
    id := c.Param("id")
    var user User
    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// Create a new user
func createUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&user)
    c.JSON(http.StatusCreated, user)
}

// Update a user by ID
func updateUser(c *gin.Context) {
    id := c.Param("id")
    var user User
    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Save(&user)
    c.JSON(http.StatusOK, user)
}

// Delete a user by ID
func deleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := db.Delete(&User{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.Status(http.StatusNoContent)
}
