package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"todo/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddTodo(context *gin.Context) {
	var todo models.Todo
	if err := context.Bind(&todo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", models.DbInfo)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 400": "database error"})
		return
	}
	query := "insert into todos (user_id,title,description,due_date, is_finished) values(?,?,?,?,?)"
	userId, err := parseToken(todo.Token)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 401": "token error"})
		return
	}

	_, err = db.Exec(query, userId, todo.Title, todo.Description, todo.DueDate, 0)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error 400": "failed to add todo"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "todo added succesfully"})
	db.Close()
}
func parseToken(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userID := int(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
func ListTodos(context *gin.Context) {
	type UserToken struct {
		Token string `json:"token,omitempty"`
	}
	var userToken UserToken
	if err := context.Bind(&userToken); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", models.DbInfo)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error opening database"})
		return
	}
	userId, err := parseToken(userToken.Token)
	if err != nil {

		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "token error"})
		return
	}

	query := "select title,description, due_date,is_finished from todos where user_id=?"
	rows, err := db.Query(query, userId)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error on database rows"})
		return
	}
	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Title, &todo.Description, &todo.DueDate, &todo.IsFinished)
		if err != nil {
			fmt.Println(err)
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error on database rows"})
			return
		}
		todos = append(todos, todo)
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func ChangeCondition(context *gin.Context) {
	type TempTodo struct {
		Id        int    `json:"id"`
		Token     string `json:"token"`
		Condition bool   `json:"condition"`
	}
	var tempTodo TempTodo
	if err := context.Bind(&tempTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": "error opening database"})
		return
	}
	userId, err := parseToken(tempTodo.Token)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "token error"})
		return
	}

	query := "update todos set is_finished=? where id=? and user_id=?"
	result, err := db.Exec(query, tempTodo.Condition, tempTodo.Id, userId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error on database while changing condition"})
		return
	}
	fmt.Println(result.RowsAffected())
	context.IndentedJSON(http.StatusOK, gin.H{"succes": "changement is made succesfully"})
	db.Close()
}
func RemoveTodo(context *gin.Context) {
	type TempTodo struct {
		Id    int    `json:"id"`
		Token string `json:"token"`
	}
	var tempTodo TempTodo
	if err := context.Bind(&tempTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error on parameters"})
		return
	}
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": "error opening database"})
		return
	}
	userId, err := parseToken(tempTodo.Token)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "token error"})
		return
	}
	query := "delete from todos where id=? and user_id=?"
	result, err := db.Exec(query, tempTodo.Id, userId)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error on database"})
		return
	}
	fmt.Println(result.RowsAffected())
	context.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted succesfully"})
	db.Close()
}
