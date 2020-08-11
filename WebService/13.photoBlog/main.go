package main

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"net/http"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"strconv"
)

var temp *template.Template
var db *sql.DB
var err error

type userData struct {
	id int
	Email string
	encryptedPassword string
	FirstName string
	LastName string
	Age int
	Drink string
}

type session struct {
	id int
	Email string
	sessionID string
	lastUpdate time.Time
}

func init() {
	temp = template.Must(template.ParseGlob("templates/*"))

	// set up db
	db, err = sql.Open("mysql", "root:PenguKnight1994!@tcp(localhost:3306)/testdb?charset=utf8&parseTime=true")
	checkError(err)

	err = db.Ping()
	checkError(err)

	// create session - email table
	stmt, err := db.Prepare(`
	create table if not exists user_session (
		user_session_id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
		email varchar(100) NOT NULL UNIQUE,
		sessionId varchar(100) NOT NULL,
		lastUpdate timestamp
	);`)
	checkError(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	checkError(err)

	numR, err := result.RowsAffected()
	checkError(err)

	fmt.Print(numR, " ")

	// create user data table
	stmt, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS user_data (
		user_data_id INT PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(100) UNIQUE NOT NULL,
		encrypted_password VARCHAR(100),
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		age INT,
		fav_drink VARCHAR(25)
	);`)
	checkError(err)

	result, err = stmt.Exec()
	checkError(err)

	numR, err = result.RowsAffected()
	checkError(err)

	fmt.Print(numR, " ")

	// create user images table
	stmt, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS user_images (
		user_images_id INT PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(100) NOT NULL,
		image_path VARCHAR(255) NOT NULL
	);`)
	checkError(err)

	result, err = stmt.Exec()
	checkError(err)

	numR, err = result.RowsAffected()
	checkError(err)

	fmt.Print(numR, "\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/userInfo", userInfo)
	http.HandleFunc("/error", errorPage)
	http.HandleFunc("/signout", signOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	// can't defer close db in init() else db conn will close at end of iniy()
	db.Close()
}

func index(w http.ResponseWriter, req *http.Request) {
	// if already session, then redirect
	_, err := req.Cookie("session")
	if err == nil {
		http.Redirect(w, req, "/userInfo", http.StatusTemporaryRedirect)
		return
	}

	// check if method is POST
	if req.Method == http.MethodPost {
		var encryptedPassword string = ""
		var row *sql.Row = db.QueryRow(`SELECT encrypted_password FROM user_data WHERE email=?;`, req.FormValue("Email"))
		err = row.Scan(&encryptedPassword)
		if err != nil {
			// no user found for the given email
			fmt.Println("no user found for the given email: ", req.FormValue("Email"), err.Error())
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(req.FormValue("Password")))
		if err != nil {
			// passwords do not match
			fmt.Println("passwords do not match", err.Error())
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}

		// get UUID for session id
		sessionID := uuid.NewV4()
		
		// set db record in user_session table
		stmt, err := db.Prepare(`INSERT INTO user_session VALUES(null, ?, ?, ?);`)
		if err != nil {
			// can't insert into database
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}
		defer stmt.Close()

		result, err := stmt.Exec(
			req.FormValue("Email"),
			sessionID.String(),
			time.Now(),
		)
		if err != nil {
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}

		numR, err := result.RowsAffected()
		checkError(err)

		fmt.Println(numR)

		// set session cookie
		http.SetCookie(w, &http.Cookie{
			Name: "session",
			Value: sessionID.String(),
			HttpOnly: true,
		})

		// redirect to userInfo
		http.Redirect(w, req, "/userInfo", http.StatusSeeOther)
		return
	}

	temp.ExecuteTemplate(w, "index.html", nil)
}

func register(w http.ResponseWriter, req *http.Request) {
	// if already session, then redirect
	_, err := req.Cookie("session")
	if err == nil {
		http.Redirect(w, req, "/userInfo", http.StatusTemporaryRedirect)
		return
	}

	// check if method is POST
	if req.Method == http.MethodPost {
		stmt, err := db.Prepare(`INSERT INTO user_data VALUES(null, ?, ?, ?, ?, ?, ?);`)
		if err != nil {
			fmt.Println("can't prepare stmt for insertion into user_data for register post", err.Error())
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}
		defer stmt.Close()

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("Password")), bcrypt.MinCost)
		checkError(err)

		result, err := stmt.Exec(
			req.FormValue("Email"),
			encryptedPassword,
			req.FormValue("FirstName"),
			req.FormValue("LastName"),
			req.FormValue("Age"),
			req.FormValue("Drink"),
		)

		// if insertion to database is not ok
		if err != nil {
			fmt.Println("could not insert into user_data register post", err.Error())
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}

		numR, err := result.RowsAffected()
		checkError(err)
		fmt.Println(numR)

		// need to create a dir in user_images dir to store this user's files
		var userIDBuffer int
		err = db.QueryRow(`SELECT user_data_id FROM user_data WHERE email=?`, req.FormValue("Email")).Scan(&userIDBuffer)
		if err != nil {
			fmt.Println("error trying to retrieve user with email: ", req.FormValue("Email"))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err = os.Stat("./user_images/" + "user_Id_" + strconv.Itoa(userIDBuffer)); os.IsNotExist(err) {
			os.Mkdir("./user_images/" + "user_Id_" + strconv.Itoa(userIDBuffer), os.ModeDir)
		}

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	temp.ExecuteTemplate(w, "register.html", nil)
}

func userInfo(w http.ResponseWriter, req *http.Request) {
	// check if there is valid session
	sessionCookie, err := req.Cookie("session")
	if err != nil {
		// cookie not found
		fmt.Println("cookie not found", err.Error())
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
		return
	}

	// get the session record
	var sessionBuffer session
	row := db.QueryRow(`SELECT * FROM user_session WHERE sessionId=?;`, sessionCookie.Value)
	err = row.Scan(&sessionBuffer.id, &sessionBuffer.Email, &sessionBuffer.sessionID, &sessionBuffer.lastUpdate)
	if err != nil {
		fmt.Println("queryrow unsuccessful for sessionId: ", sessionCookie.Value, err.Error())
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}

	// check if session time has expired
	if time.Now().Sub(sessionBuffer.lastUpdate).Seconds() > 3600 {
		fmt.Println("session expired")
		http.Redirect(w, req, "/signout", http.StatusTemporaryRedirect)
		return
	}
	// refresh session time
	stmt, err := db.Prepare(`UPDATE user_session SET lastUpdate=? where sessionId=?`)
	if err != nil {
		// can't update database
		fmt.Println("prepare stmt for user_session failed for userInfo refresh", err.Error())
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(time.Now(), sessionCookie.Value)
	if err != nil {
		fmt.Println("could not update user_session to refresh timer", err.Error())
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}

	numR, err := result.RowsAffected()
	checkError(err)

	fmt.Println("session refreshed ", numR)

	// get the userData record
	var userBuffer userData
	row = db.QueryRow(`SELECT * FROM user_data WHERE email=?;`, sessionBuffer.Email)
	err = row.Scan(&userBuffer.id, &userBuffer.Email, &userBuffer.encryptedPassword, &userBuffer.FirstName,
		&userBuffer.LastName, &userBuffer.Age, &userBuffer.Drink)
	if err != nil {
		fmt.Println("queryrow unsuccessful for email: ", userBuffer.Email)
	}
	// fmt.Println(userBuffer)

	// if post, then save image on server file system and store that path in database
	if req.Method == http.MethodPost {
		file, fileHeader, err := req.FormFile("image")
		if err != nil {
			fmt.Println("error uploading image ", err.Error())
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		byteStream, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("could not read uploaded image file ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// debug
		// fmt.Println(userBuffer, userBuffer.id, string(userBuffer.id), strconv.Itoa(userBuffer.id))
		pathStringBuffer := string(filepath.Join("./user_images",  "user_Id_" + strconv.Itoa(userBuffer.id), uuid.NewV4().String() + "_" + fileHeader.Filename))
		dst, err := os.Create(pathStringBuffer)
		if err != nil {
			fmt.Println("could not create file on server ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(byteStream)
		if err != nil {
			fmt.Println("could not write image to file system file ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// save the image path to db
		stmt, err = db.Prepare(`INSERT INTO user_images VALUES(null, ?, ?);`)
		if err != nil {
			fmt.Println("could not prepare stmt to insert image to db ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err = stmt.Exec(userBuffer.Email, pathStringBuffer)
		if err != nil {
			fmt.Println("could not insert image to db ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		numR, err = result.RowsAffected()
		checkError(err)

		fmt.Println(numR)
	}

	// need to get the images for a user
	var userImages []string

	rows, err := db.Query(`SELECT image_path FROM user_images WHERE email=?`, userBuffer.Email)
	if err != nil {
		fmt.Println("error pulling images for user with email: ", userBuffer.Email, err.Error())
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}

	var pathBuffer string
	for rows.Next() {
		err = rows.Scan(&pathBuffer)
		if err != nil {
			fmt.Println("error scanning path from db to string variable", err.Error())
			http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
			return
		}
		userImages = append(userImages, pathBuffer)
	}

	fmt.Println(userImages)

	dataPayload := struct {
		Email string
		FirstName string
		LastName string
		Age int
		Drink string
		Images []string
	} {
		userBuffer.Email,
		userBuffer.FirstName,
		userBuffer.LastName,
		userBuffer.Age,
		userBuffer.Drink,
		userImages,
	}

	temp.ExecuteTemplate(w, "userInfo.html", dataPayload)
}

func signOut(w http.ResponseWriter, req *http.Request) {
	// remove session record from database
	sessionCookie, err := req.Cookie("session")
	if err != nil {
		fmt.Println("error getting session cookie", err.Error())
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}

	stmt, err := db.Prepare(`DELETE FROM user_session WHERE sessionId=?;`)
	if err != nil {
		fmt.Println("could not delete from user_session")
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(sessionCookie.Value)
	if err != nil {
		http.Redirect(w, req, "/error", http.StatusTemporaryRedirect)
		return
	}

	numR, err := result.RowsAffected()
	checkError(err)

	fmt.Println(numR)

	// clear all cookies
	for _, cookie := range req.Cookies() {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func errorPage(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "error.html", err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}