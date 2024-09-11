package main

import (
	"fmt"
	"inventory/routers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("inventory API")
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":4002", r))
}

// data := [][]string{}
// data = append(data, []string{"secret", "address"})
// csvExport(data)
// func csvExport(data [][]string) error {
// 	file, err := os.Create("result.csv")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	for _, value := range data {
// 		if err := writer.Write(value); err != nil {
// 			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
// 		}
// 	}
// 	return nil
// }
