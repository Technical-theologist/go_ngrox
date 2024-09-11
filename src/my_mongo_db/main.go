// mongodb+srv://rakesh:<password>@cluster0.gzipf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0

//https://cloud.mongodb.com/v2/66b9f716b802cf387894aa1d#/metrics/replicaSet/66b9f95ff9229f6d2beb7e69/explorer/netflix/watchList/find
//The above one is mango db website url with my rakesh@gosimple.in account

package main

import (
	"fmt"
	"log"
	"my_mango/routers"
	"net/http"
)

func main() {
	fmt.Println("MangoDB API")
	r := routers.Router()
	fmt.Println("Server is getting  started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 .....")
}
