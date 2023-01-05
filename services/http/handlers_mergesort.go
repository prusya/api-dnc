package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/prusya/api-dnc/services"

	"github.com/gin-gonic/gin"
)

func getMergesort(c *gin.Context) {
	arg := c.Query("arr")
	args := strings.Split(arg, ",")
	arr := []int{}
	for i := 0; i < len(args); i++ {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, "w")
			return
		}
		arr = append(arr, n)
	}
	fmt.Println(arr)
	sorted, err := services.Mergesort.Sort(arr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "w")
		return
	}
	c.JSON(http.StatusOK, sorted)
}
