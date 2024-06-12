package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
}

type UserMeta struct {
	Visit int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Ayomide Lawal",
		Bio:  `<script>alert("Hello, You are hacked")</script>`,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}

}
