package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Endpoint GET /tasks qui retourne une liste vide de tâches
    r.GET("/tasks", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"tasks": []string{}})
    })

    // Lancer le serveur sur le port 8080
    r.Run(":8080")
}
