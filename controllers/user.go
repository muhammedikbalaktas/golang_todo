package controllers

import (
	"fmt"
	"net/http"
	c "todo/models"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var dbinfo string = c.DbInfo

func CreateUser(context *gin.Context) {
	var user c.User
	if err := context.Bind(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 100": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)
	defer db.Close()
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
	var user c.User
	if err := context.Bind(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 100": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	query := "select username, password from user where username=?"
	rows, err := db.Query(query, user.Username)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	var username, password string
	for rows.Next() {

		rows.Scan(&username, &password)
		fmt.Println(username)
		fmt.Println(password)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 101": "invalid password or username"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "username and password is correct"})
}
