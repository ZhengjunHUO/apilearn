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
	// eg. $ curl http://localhost:8088/clusters
	server.GET("/clusters", getClusters)

	// eg. curl http://localhost:8088/clusters/1
	server.GET("/clusters/:id", getClusterByID)

	/* eg.
	$ curl http://localhost:8088/clusters \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"id": "4","cni": "awsvpc","size": 6,"ismanaged": true, "isbaremetal": false, "isoverlay": true}'
	*/
	server.POST("/clusters", addCluster)

	server.Run("localhost:8088")
}

// handler func to get a list of k8s cluster
func getClusters(c *gin.Context) {
	// Context carries request details and more ...
	// marshaling the struct to json & add to response
	c.IndentedJSON(http.StatusOK, clusters)
}

// handler func to get a specific k8s cluster
func getClusterByID(c *gin.Context) {
	// need a placeholder ":id" in the path when register this handler to api server
	id := c.Param("id")

	for _, cluster := range clusters {
		if cluster.ID == id {
			c.IndentedJSON(http.StatusOK, cluster)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"INFO": "cluster not found"})
}

// handler func to add a k8s cluster to the pool
func addCluster(c *gin.Context) {
	var cluster k8sCluster

	// Bind the request body(recieved json) in to a k8sCluster struct
	if err := c.BindJSON(&cluster); err != nil {
		return
	}

	clusters = append(clusters, cluster)
	c.IndentedJSON(http.StatusCreated, cluster)
}
