package main

import (
	"time"
	"net/http"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var temp *template.Template

type user struct {
	Email string
	Password string
	First string
	Last string
	Role string
}

// Session ...2 fields with user email and timestamp
type Session struct {
	UN string
	Timestamp time.Time
}

var sessions map[string]*Session = make(map[string]*Session)
var info map[string]user = make(map[string]user)

func foo(w http.ResponseWriter, req *http.Request) {
	// if session exists then redirect to bar
	_, err := req.Cookie("session")
	if err == nil {
		// cookie exists
		fmt.Println("has 'session' cookie")
		http.Redirect(w, req, "/bar", http.StatusTemporaryRedirect)
		return
	}

	if req.Method == http.MethodPost {
		fmt.Println("validating user POST")
		// fmt.Println(sessions, info)
		// fmt.Println(req.FormValue("user"), req.FormValue("user") == "test@test.com")
		bufUser, ok := info[req.FormValue("user")]
		if !ok {
			fmt.Fprintln(w, `<p>no email found, please register</p><a href="/register">Register</a>`)
			return
		}
		// fmt.Println(bufUser.Password, req.FormValue("pass"))
		err = bcrypt.CompareHashAndPassword([]byte(bufUser.Password), []byte(req.FormValue("pass")))
		fmt.Println(err)

		if err != nil {
			fmt.Fprintln(w, `<p>invalid password</p><a href="/">Sign In</a>`)
			return
		}

		sessionID := uuid.NewV4()
		bufSession := Session {
			UN: req.FormValue("user"),
			Timestamp: time.Now(),
		}
		fmt.Println(time.Now())
		sessions[sessionID.String()] = &bufSession

		http.SetCookie(w, &http.Cookie {
			Name: "session",
			Value: sessionID.String(),
		})

		http.Redirect(w, req, "/bar", http.StatusSeeOther)
		fmt.Println("done validating user POST")
		return
	}

	temp.ExecuteTemplate(w, "signin.html", nil)
}

func register(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fmt.Println("registering user")
		cipher, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)
		if err != nil {
			log.Println("could not create cipher")
			fmt.Fprintln(w, "could not create cipher")
			return
		}
		bufUser := user {
			Email: req.FormValue("email"),
			Password: string(cipher),
			First: req.FormValue("firstname"),
			Last: req.FormValue("lastname"),
			Role: req.FormValue("role"),
		}
		info[req.FormValue("email")] = bufUser

		// fmt.Println(cipher, bufUser, sessions, info, req.FormValue("password"))
		// check if user already exists with same email

		http.Redirect(w, req, "/", http.StatusSeeOther)
		fmt.Println("done registering")
		return
	}

	temp.ExecuteTemplate(w, "register.html", nil)
}

func expire(w http.ResponseWriter, req *http.Request) {
	sess, _ := req.Cookie("session")
	// check err

	delete(sessions, sess.Value)

	fmt.Println(sessions)

	var cookies []*http.Cookie = req.Cookies()

	for _, cookie := range cookies {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}

	// or just remove the cookie name "session"

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func bar(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/deny", http.StatusTemporaryRedirect)
		return
	}
	userSession, ok := sessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/deny", http.StatusTemporaryRedirect)
		return
	}

	// if session if over
	if time.Now().Sub(userSession.Timestamp).Seconds() >= 20 {
		// delete(sessions, cookie.Value)
		// cookie.MaxAge = -1
		// http.SetCookie(w, cookie)
		// http.Redirect(w, req, "/deny", http.StatusTemporaryRedirect)
		// return
		http.Redirect(w, req, "/expire", http.StatusSeeOther)
	} else { // refresh the session
		fmt.Println(userSession.Timestamp)
		userSession.Timestamp = time.Now()
		fmt.Println(userSession.Timestamp)
	}

	user, ok := info[userSession.UN]
	if !ok {
		fmt.Fprintln(w, "user not found")
		return
	}
	temp.ExecuteTemplate(w, "info.html", user)
}

func deny(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<p>access denied</p><a href="/">Sign In</a>`)
}

func admin(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/deny", http.StatusSeeOther)
		return
	}

	userSession, ok := sessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/deny", http.StatusSeeOther)
		return
	}

	// if session if over
	if time.Now().Sub(userSession.Timestamp).Seconds() >= 20 {
		// delete(sessions, cookie.Value)
		// cookie.MaxAge = -1
		// http.SetCookie(w, cookie)
		// http.Redirect(w, req, "/deny", http.StatusTemporaryRedirect)
		// return
		http.Redirect(w, req, "/expire", http.StatusSeeOther)
	} else { // refresh the session
		// fmt.Println(userSession.Timestamp)
		userSession.Timestamp = time.Now()
		fmt.Println(userSession.Timestamp)
	}

	bufUser, ok := info[userSession.UN]
	if !ok {
		fmt.Fprintln(w, `<p>user not found</p><a href="/">Sign In</a>`)
		return
	}
	if bufUser.Role != "admin" {
		http.Redirect(w, req, "/deny", http.StatusSeeOther)
		return
	}

	temp.ExecuteTemplate(w, "admin.html", bufUser)
}

func init() {
	temp = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/register", register)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/expire", expire)
	http.HandleFunc("/deny", deny)
	http.HandleFunc("/admin", admin)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}