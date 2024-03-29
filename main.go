package main

import (
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
    app := iris.New()
    
    /*
        copy all files from https://github.com/swagger-api/swagger-ui/dist
        to ./swagger-ui, in this folder there will also be a swaager.json file generated by
        running generate-from-code.cmd
    */
    app.HandleDir("/swagger-ui", iris.Dir("./swagger-ui"))
    
	err := app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":9000"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if err != nil {
		log.Println(err.Error())
	}
}