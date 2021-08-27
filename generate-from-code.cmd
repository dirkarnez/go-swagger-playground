docker run --rm -it --env GOPATH=/go -v "%CD%:/go/src/dirkarnez/go-swagger-playground" -w /go/src/dirkarnez/go-swagger-playground ^
 quay.io/goswagger/swagger generate spec --scan-models -o ./swagger-ui/swagger.json

REM remember to add  base <href="/swagger-ui/"> to swagger-ui/index.html
pause
