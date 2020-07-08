package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"net"
	"net/http"
	"net/http/fcgi"
//	"regexp"
)


var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("17923641793298746918723649781112")
	store = sessions.NewCookieStore(key)
)

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

/*var validnon = regexp.MustCompile("^/(postform|reg|regproc|enter|sign|post|posts|exit|home|profile|filesave)//(.*)$")
var vaitdn = regexp.MustCompile("^/(comform|comment|users|postview)/(.*)$")
var validsoap = regexp.MustCompile("^/(checkemail)/(.*)$")


func chknon(w http.ResponseWriter, r *http.Request)  {
	m := validnon.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.Redirect(w,r,"/static/404.htmml",301)
	}
}
func chkn(w http.ResponseWriter, r *http.Request) {
	m := vaitdn.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.Redirect(w,r,"/static/404.htmml",301)
	}
} 
func chksoap(w http.ResponseWriter, r *http.Request) {
	m := validsoap.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.Redirect(w,r,"/static/404.htmml",301)
	}
}
*/

func main() {
//	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	//	mux.HandleFunc("/user/)
	http.Handle("/reg", fs)
	http.HandleFunc("/read/",htmlhandle)
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

//	http.ListenAndServe("127.0.0.1:8000",nil)
}
