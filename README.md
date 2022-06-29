# apilearn

## Example
### Run the api server
```
$ go run main.go 
```
### Requests
```
# Show a certain cluster's info
$ curl http://localhost:8088/clusters/1

# Create a cluster
$ curl http://localhost:8088/clusters \
                --include \
                --header "Content-Type: application/json" \
                --request "POST" \
                --data '{"id": "4","cni": "awsvpc","size": 6,"ismanaged": true, "isbaremetal": false, "isoverlay": true}'

# Show all clusters
$ curl http://localhost:8088/clusters  
```
