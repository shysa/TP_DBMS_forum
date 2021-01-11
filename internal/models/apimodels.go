package models

import "time"

//easyjson:json
type Forum struct {
	Posts   int    `json:"posts,omitempty"`
	Slug    string `json:"slug"`
	Threads int    `json:"threads,omitempty"`
	Title   string `json:"title"`
	User    string `json:"user"`
}

//easyjson:json
type Post struct {
	Id       int       `json:"id"`
	Author   string    `json:"author"`
	Created  time.Time `json:"created,omitempty"`
	Forum    string    `json:"forum,omitempty"`
	IsEdited bool      `json:"isEdited,omitempty"`
	Message  string    `json:"message"`
	Parent   int       `json:"parent,omitempty"`
	Thread   int       `json:"thread,omitempty"`
	Tree     []int     `json:"tree,omitempty"`
}
//easyjson:json
type PostDetails struct {
	Author *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   *Post   `json:"post,omitempty"`
	Thread *Thread `json:"thread,omitempty"`
}
//easyjson:json
type PostUpdate struct {
	Message string `json:"message,omitempty"`
}
//easyjson:json
type Posts []Post

type Parents struct {
	ParentId int
	Tree     []int
}

//easyjson:json
type Thread struct {
	Id      int64     `json:"id"`
	Author  string    `json:"author"`
	Created time.Time `json:"created,omitempty"`
	Forum   string    `json:"forum,omitempty"`
	Message string    `json:"message"`
	Slug    string    `json:"slug,omitempty"`
	Title   string    `json:"title"`
	Votes   int       `json:"votes,omitempty"`
}
//easyjson:json
type Threads []Thread

//easyjson:json
type ThreadUpdate struct {
	Message string `json:"message,omitempty"`
	Title   string `json:"title,omitempty"`
}

//easyjson:json
type User struct {
	Nickname string `json:"nickname"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}
//easyjson:json
type Users []User

//easyjson:json
type UserUpdate struct {
	About    string `json:"about,omitempty"`
	Email    string `json:"email,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}

//easyjson:json
type Status struct {
	Forum  int `json:"forum"`
	Post   int `json:"post"`
	Thread int `json:"thread"`
	User   int `json:"user"`
}

//easyjson:json
type Vote struct {
	Nickname string `json:"nickname"`
	Voice    int32  `json:"voice"`
}

type Params struct {
	Limit   int    `form:"limit"`
	Since   string `form:"since"`
	Desc    bool   `form:"desc"`
	Sort    string `form:"sort"`
	Related string `form:"related"`
}
