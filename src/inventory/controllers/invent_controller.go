package controllers

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"inventory/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var db *sql.DB
var secretKey = []byte("secret-key")

func init() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "rakesh12",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "employee",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}
	fmt.Println("connected")
}
func createProduct(prod models.Product) {
	query := "insert into products (productname,productdescription,price,quantity,category,createddate) values(?,?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Exec(prod.ProductName, prod.ProductDescription, prod.Price, prod.Quantity, prod.Category, prod.CreatedDate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.RowsAffected())
}
func deleteProduct(id string) {
	res, err := db.Exec("delete from products where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())
}
func getAllProduct() []models.Product {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := db.QueryContext(ctx, "SELECT * FROM products")
	defer cancelfunc()
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	var products []models.Product
	for res.Next() {

		var sProduct models.Product
		var unusedColumn interface{}
		if err := res.Scan(&sProduct.Id, &sProduct.ProductName,
			&sProduct.ProductDescription, &sProduct.Price, &sProduct.Quantity, &sProduct.Category, &unusedColumn); err != nil {
			fmt.Println(err)
			return products
		}
		products = append(products, sProduct)
	}
	return products
}
func searchByTitle(title string) []models.Product {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	rows, err := db.QueryContext(ctx, "SELECT * FROM products WHERE productname LIKE ?", "%"+title+"%")
	defer cancelfunc()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var sProduct models.Product
		var unusedColumn interface{}
		if err := rows.Scan(&sProduct.Id, &sProduct.ProductName, &sProduct.ProductDescription, &sProduct.Price, &sProduct.Quantity, &sProduct.Category, &unusedColumn); err != nil {
			fmt.Println(err)
			return products
		}
		fmt.Println("my s prod", sProduct)
		products = append(products, sProduct)
	}
	return products
}
func getOneProduct(id string) []models.Product {
	res, err := db.Query("select * from products where id= ?", id)
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	var products []models.Product
	for res.Next() {
		var unusedColumn interface{}
		var sProduct models.Product
		if err := res.Scan(&sProduct.Id, &sProduct.ProductName,
			&sProduct.ProductDescription, &sProduct.Price, &sProduct.Quantity, &sProduct.Category, &unusedColumn); err != nil {
			fmt.Println(err)
			return products
		}
		products = append(products, sProduct)
	}
	return products
}
func updateOneProduct(id string, prod models.Product) string {
	query := "UPDATE products SET productname = ?, productdescription = ?, price = ?, quantity = ?, category = ?,createddate = ? WHERE id = ?"
	result, err := db.Exec(query, prod.ProductName, prod.ProductDescription, prod.Price, prod.Quantity, prod.Category, prod.CreatedDate, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
	return id
}
func filterBy(filter models.FilterModel) []models.Product {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	x := 100000000
	if filter.InStock {
		x = 0
	}

	query := "SELECT * FROM products WHERE price BETWEEN ? AND ? AND quantity > ?"
	rows, err := db.QueryContext(ctx, query, filter.StartingPrice, filter.EndingPrice, x)
	defer cancelfunc()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var sProduct models.Product
		var unusedColumn interface{}
		if err := rows.Scan(&sProduct.Id, &sProduct.ProductName, &sProduct.ProductDescription, &sProduct.Price, &sProduct.Quantity, &sProduct.Category, &unusedColumn); err != nil {
			fmt.Println(err)
			return products
		}
		fmt.Println("my s prod", sProduct)
		products = append(products, sProduct)
	}
	return products
}

func bulkDataUpload() {
	file, err := os.Open("C:\\Users\\accou\\Downloads\\test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(records)

	for _, record := range records {
		query := "INSERT INTO countries (id, country, population, capital) values(?,?,?,?)"
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancelfunc()
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Exec(record[0], record[1], record[2], record[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rows.RowsAffected())
	}

}

func createUser(user models.Users) {
	query := "insert into users (lastname,firstname,address,city,email,pass,role) values(?,?,?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Exec(user.LastName, user.FirstName, user.Adresss, user.City, user.Email, user.Password, user.Role)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.RowsAffected())

}

func userLogin(user models.Users) []models.Users {
	res, err := db.Query("select * from users where email= ? and pass =?", user.Email, user.Password)
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	var users []models.Users
	var tokenString string = ""
	for res.Next() {
		var unusedColumn interface{}
		var sUser models.Users
		if err := res.Scan(&sUser.ID, &sUser.LastName,
			&sUser.FirstName, &sUser.Adresss, &sUser.City, &sUser.Email, &unusedColumn, &sUser.Role); err != nil {
			fmt.Println(err)
			return users
		}
		users = append(users, sUser)
		fmt.Println("my users data", users)
	}

	if len(users) > 0 {
		fmt.Println("login role", users[0].Role)
		tokenString, err = CreateToken(users[0].Email, users[0].Role)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenString)
	return users
}

func CreateToken(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func UpdateOneProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var prod models.Product
	json.NewDecoder(r.Body).Decode(&prod)
	var prodId = updateOneProduct(params["id"], prod)
	json.NewEncoder(w).Encode("product updated " + prodId)
}
func SearchByTitle(w http.ResponseWriter, r *http.Request) {
	urlstr := r.URL
	var title = urlstr.Query().Get("title")
	myproduct := searchByTitle(title)
	json.NewEncoder(w).Encode(myproduct)
}
func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	myproduct := getOneProduct(params["id"])
	json.NewEncoder(w).Encode(myproduct)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	myProducts := getAllProduct()
	json.NewEncoder(w).Encode(myProducts)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	urlstr := r.URL
	var id = urlstr.Query().Get("id")
	deleteProduct(id)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	var myProduct models.Product
	_ = json.NewDecoder(r.Body).Decode(&myProduct)
	createProduct(myProduct)
	json.NewEncoder(w).Encode(myProduct)
}
func FilterBy(w http.ResponseWriter, r *http.Request) {
	var myfilter models.FilterModel
	_ = json.NewDecoder(r.Body).Decode(&myfilter)
	var myfilteredList = filterBy(myfilter)
	json.NewEncoder(w).Encode(&myfilteredList)
}
func BulkDataUpload(w http.ResponseWriter, r *http.Request) {
	bulkDataUpload()
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var myUser models.Users
	_ = json.NewDecoder(r.Body).Decode(&myUser)
	createUser(myUser)
	json.NewEncoder(w).Encode(myUser)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	userLogin(user)
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hitting file upload")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	os.MkdirAll("uploads", os.ModePerm)
	filePath := filepath.Join("uploads", handler.Filename)
	tempFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	_, err = io.Copy(tempFile, file)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current working directory:", dir)
	absPath, _ := filepath.Abs(filePath)
	log.Println("File saved as:", filePath)
	log.Println("Absolute file path:", absPath)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}
