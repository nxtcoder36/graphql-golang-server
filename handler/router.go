package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/nxtCoder36/graphql-golang-server/Impl"
	"net/http"
)

type RequestParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func TodoGraphRouter(c *gin.Context) {
	var reqObj RequestParams
	if err := c.ShouldBindJSON(&reqObj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         Impl.TodoSchema,
		RequestString:  reqObj.Query,
		VariableValues: reqObj.Variables,
		OperationName:  reqObj.Operation,
	})
	if len(result.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Errors})
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
}
