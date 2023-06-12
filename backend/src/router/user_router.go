package router

import (
	"net/http"
	"rakukaji/model"
)

// test data
var test_users = []model.User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func UserCreateTest(w http.ResponseWriter, r *http.Request) {
	for _, user := range test_users {
		err := user.Create()
		if err != nil {
			return
		}
	}
	users, err := model.Users()
	if err != nil {
		return
	}
	for _, user := range users{
		w.Write([]byte(user.Name))
	}
}
