package models

type User struct {
	Id         int    `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// +------------+-------------+------+-----+---------+----------------+
// | Field      | Type        | Null | Key | Default | Extra          |
// +------------+-------------+------+-----+---------+----------------+
// | id         | int         | NO   | PRI | NULL    | auto_increment |
// | username   | varchar(20) | NO   | UNI | NULL    |                |
// | email      | varchar(30) | NO   | UNI | NULL    |                |
// | password   | varchar(60) | NO   |     | NULL    |                |
// | created_at | date        | NO   |     | NULL    |                |
// +------------+-------------+------+-----+---------+----------------+
