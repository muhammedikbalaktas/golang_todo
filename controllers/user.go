package controllers

import (
	"fmt"
	"net/http"
	models "todo/models"

	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbinfo    string = models.DbInfo
	secretKey        = []byte("your_password")
)

func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.Bind(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 100": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	hashedPassword, err := encrpytPassword(user.Password)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 502": "internal server error"})
		return
	}
	query := "insert into user (username, email,password,created_at) values(?,?,?, curdate())"
	_, err = db.Exec(query, user.Username, user.Email, hashedPassword)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 501": "internal server error"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "user created succesfully"})
	db.Close()
}
func encrpytPassword(password string) (*string, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(hashedPasswordByte)
	return &hashedPassword, nil
}
func GetUser(context *gin.Context) {
	var user models.User
	if err := context.Bind(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 100": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)

	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	query := "select id,username, password from user where username=?"
	rows, err := db.Query(query, user.Username)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	var username, password string
	var userId int
	for rows.Next() {

		rows.Scan(&userId, &username, &password)
		fmt.Println(username)
		fmt.Println(password)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 101": "invalid password or username"})
		return
	}
	token, err := generateToken(userId)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"succes": token})
	db.Close()
}
func generateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})
	tokeString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokeString, nil
}
