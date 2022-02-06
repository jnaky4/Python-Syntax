package main

import (
	"awesomeProject/httpserver/WebService/controllers"
	"log"
	"net/http"
)

/*To server up the webservice, navigate to packages in terminal:
	run: go build .
	will create <packagename>.exe
	ls will show for this WebService.exe
	then type: .\WebService.exe
*/

func main(){
	controllers.RegisterControllers()
	//2 params, IP address, SerMux(serveMultiplexor)
	//handle all requests coming in, (front controller) -> decided which back controller
	//handles that request defined in reqisterControllers
	err := http.ListenAndServe(
		":3000", nil)
	if err != nil {
		log.Fatal("Failed to Serve on Port 3000, Reason:", err )
		return
	}
}
