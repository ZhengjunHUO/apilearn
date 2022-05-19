package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type k8sCluster struct {
	ID		string	`json:"id"`
	CNI		string	`json:"cni"`
	Size		int	`json:"size"`
	IsManaged	bool	`json:"ismanaged"`
	IsBaremetal	bool	`json:"isbaremetal"`
	IsOverlay	bool	`json:"isoverlay"`
}

var clusters = []k8sCluster{
	{ID: "1", CNI: "flannel", Size: 3, IsManaged: true, IsBaremetal: false, IsOverlay: true},
	{ID: "2", CNI: "cilium", Size: 10, IsManaged: false, IsBaremetal: true, IsOverlay: false},
	{ID: "3", CNI: "calico", Size: 7, IsManaged: false, IsBaremetal: true, IsOverlay: true},
}

func main() {
	server := gin.Default()
	// associate GET methods and path with a handler func
	server.GET("/clusters", getClusters)

	server.Run("localhost:8088")
}

// hander func
func getClusters(c *gin.Context) {
	// Context carries request details and more ...
	// marshaling the struct to json & add to response
	c.IndentedJSON(http.StatusOK, clusters)
}
