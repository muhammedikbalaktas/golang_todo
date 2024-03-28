package models

type Todo struct {
	Token       string `json:"token,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	IsFinished  bool   `json:"is_finished,omitempty"`
}

// +-------------+-------------+------+-----+---------+----------------+
// | Field       | Type        | Null | Key | Default | Extra          |
// +-------------+-------------+------+-----+---------+----------------+
// | id          | int         | NO   | PRI | NULL    | auto_increment |
// | user_id     | int         | NO   | MUL | NULL    |                |
// | title       | varchar(10) | NO   |     | NULL    |                |
// | description | varchar(20) | NO   |     | NULL    |                |
// | due_date    | date        | NO   |     | NULL    |                |
// | is_finished | tinyint(1)  | NO   |     | 0       |                |
// +-------------+-------------+------+-----+---------+----------------+
