package controllers

import (
	"fmt"
	"net/http"
	c "todo/models"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var dbinfo string = c.DbInfo

func CreateUser(context *gin.Context) {
	var user c.User
	if err := context.Bind(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error 100": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 500": "internal server error"})
		return
	}
	// create table user(
	//     id int auto_increment not null,
	//     username varchar(20) unique not null,
	//     email varchar(30) unique not null,
	//     password varchar(60) not null,
	//     created_at date not null,
	//     primary key(id)
	//     );
	query := "insert into user (username, email,password,created_at) values(?,?,?, curdate())"
	_, err = db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 501": "internal server error"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "user created succesfully"})

}
