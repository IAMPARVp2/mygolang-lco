package main

import "fmt"

func main() {
	fmt.Println("welcome to the structs")
	//no inheritance in go lang
	//no super or parent or child

	hitesh := User{"Hitesh", "jainv1071@gmail.com", true, 10, 2002}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("email is %v.\n", hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("email is %v.\n", hitesh.Email)
}

type User struct { //the first letter are capital to make them public so that we can use it at other places also
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int //this property is not exportable because first letter is not capital

}

func (u User) GetStatus() {
	fmt.Println("is user active", u.Status)
}

func (u User) NewMail() {
	u.Email = "tes@go.devt"
	fmt.Println("email of this user is", u.Email)
}

//output
/*
welcome to the structs
{Hitesh jainv1071@gmail.com true 10 2002}
hitesh details are: {Name:Hitesh Email:jainv1071@gmail.com Status:true Age:10 oneAge:2002}
email is jainv1071@gmail.com.
is user active true
email of this user is tes@go.devt
email is jainv1071@gmail.com.
*/

/*
here in thr output we can see that
before he function call the out put is original email passed in struct
but after that in function call the email is changed and after that agin got hte same email
so this scenerio is going on because we have passed
the copy of the email and in function the changes are done in copy
*/
