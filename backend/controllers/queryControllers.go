package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tanishqahluwalia25/portfolio-backend/database"
	"github.com/tanishqahluwalia25/portfolio-backend/models"
)

var queryCollection = database.OpenCollection(database.Client, "queries")

func MakeQuery() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		var query models.Query

		if err := c.BindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "There was an error finding query"})
			fmt.Println(err)
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		query.Id = primitive.NewObjectID()
		query.QueryId = query.Id.Hex()

		_, err := queryCollection.InsertOne(ctx, query)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "there was some problem adding your query"})
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusOK, bson.M{"success": "your query was added"})

	}

}
