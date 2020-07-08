package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"net"
	"net/http"
	"net/http/fcgi"
	"regexp"
)


var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("17923641793298746918723649781112")
	store = sessions.NewCookieStore(key)
)
type regdata struct {
	Invite string
	Username string
	Password string
	Resume string
	Url string
	To string
}

type regchk struct {
	Invitech int
	Usernamereg int
	Usernameex int
	Passwordrx int
	Passwordcon int
	fnrx, snrx int
    pln, fln, sln int
}

type zombie struct {
	Zho string
}

type Getout struct{
	Who string
}
type User struct {
	id int
	username string
	fname string
	sname string
	password string
	email string
	about string
	userpic string
	timereg string
}

type Signdata struct {
	Id int
	Userpic string
	Username,Fname,Sname string
}

var validStr = regexp.MustCompile("^([a-zA-Z0-9а-яА-Я]+)$")

type posttype struct
{
	Id int
	Post []byte
	Pfname string
	Psname string
	Posttime string
}
type commtype struct {
	Comment []byte
	Username, Comtime string
}

type Postdata struct {
	Id int
	Post []byte
	Fname, Sname string
	Username string
	Posttime string
	Userpic string
}

type Postformdata struct {
	Id int
	Username,Fname, Sname string
	Userpic string
}

type Comformdata struct {
	Id int
	Username,Fname,Sname string
	Userpic string
}

type Comformdata2 struct {
	Postid int
	Pusername,Fname, Sname string
	Posttime string
}
type tmp struct {
	Title, Note  string
}

func main() {
//	db := dbConnect()
//	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	//	mux.HandleFunc("/user/)
	http.handleFunc("/read/",htmlhandle)
	http.Handle("/reg", fs)
	http.HandleFunc("/regproc/", regprocHandle)
	http.Handle("/enter",fs)
	http.HandleFunc("/sign/",signHandler)
	http.HandleFunc("/postform/",postForm)
	http.HandleFunc("/post/",postHandle)
	http.HandleFunc("/posts/",postsHandle)
	http.HandleFunc("/comform/",comForm)
	http.HandleFunc("/comment/",comHandler)
	http.HandleFunc("/exit/", leaveHandler)
	http.HandleFunc("/users/",usersHandler)
	http.HandleFunc("/checkemail/",checkemail)
	http.HandleFunc("/home/",Home);
	http.HandleFunc("/profile/",fileform)
	http.HandleFunc("/filesave/",filesave)
	http.HandleFunc("/postview/",postview)
	//	mux.HandleFunc("/comment/",comments)
	//	mux.HandleFunc("/secret", secret)
	//	mux.HandleFunc("/logout", logout)
	//	r.HandleFunc("/enter", enterHandler)
	l, err := net.Listen("tcp", ":9001")
	if err != nil{
		return
	}
	fcgi.Serve(l, nil)
//	db.Close()

//	http.ListenAndServe(":8000", r)
}
