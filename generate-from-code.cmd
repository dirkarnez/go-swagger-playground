docker run --rm -it --env GOPATH=/go -v "%CD%:/go/src/dirkarnez/go-swagger-playground" -w /go/src/dirkarnez/go-swagger-playground ^
 quay.io/goswagger/swagger generate spec --scan-models -o ./swagger-ui/swagger.json

pause


