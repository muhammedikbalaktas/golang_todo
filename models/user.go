package models

type User struct {
	Id         int    `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// create table user(
//     id int auto_increment not null,
//     username varchar(20) unique not null,
//     email varchar(30) unique not null,
//     password varchar(60) not null,
//     created_at date not null,
//     primary key(id)
//     );
