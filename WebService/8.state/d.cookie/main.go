package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	// "strings"
)

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie {
		Name: "CookieName",
		Value: "CookieValue",
		Domain: "localhost",
	})

	fmt.Fprintln(res, "SET COOKIE")
	fmt.Fprintln(res, "check it out in developer tools!")
}

func read(res http.ResponseWriter, req *http.Request) {
	cook, err := req.Cookie("CookieName")
	if err != nil {
		// http.Error(res, err.Error(), http.StatusNoContent)
		log.Println("Cookie not found")
	} else {
		fmt.Fprintln(res, "Your cookie: ", cook)
	}
	// fmt.Fprintln(res, "Cookie Name is: " + cook.Name + "; Cookie Value is: " + cook.Value)

	genCook1, err := req.Cookie("genCookie1")
	if err != nil {
		log.Println("Cookie not found")
	} else {
		fmt.Fprintln(res, "Your cookie: ", genCook1)
	}

	genCook2, err := req.Cookie("genCookie2")
	if err != nil {
		log.Println("Cookie not found")
	} else {
		fmt.Fprintln(res, "Your cookie: ", genCook2)
	}
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name: "genCookie1",
		Value: "someval1",
	})
	http.SetCookie(res, &http.Cookie{
		Name: "genCookie2",
		Value: "someval2",
	})

	fmt.Fprintln(res, "COOKIES SET")
	fmt.Fprintln(res, "check storage cookies in dev tools!")
}

func visitCounter(res http.ResponseWriter, req *http.Request) {
	counter, err := req.Cookie("visitsCounter")
	if err == http.ErrNoCookie {
		http.SetCookie(res, &http.Cookie {
			Name: "visitsCounter",
			Value: "1",
		})
		fmt.Fprintln(res, "This is your first visit!")
	} else {
		// fmt.Println(counter.Value)
		// newValue, err := strconv.ParseInt(counter.Value, 10, 0)
		newValue, err := strconv.Atoi(counter.Value)
		if err != nil {
			log.Println("could not convert value")
			fmt.Fprintln(res, "error occurred: ", err.Error())
			return
		}
		newValue++
		counter.Value = strconv.Itoa(newValue)
		http.SetCookie(res, counter)
		fmt.Fprintln(res, "your current visit to this page is: ", newValue)
	}
}

func expire(res http.ResponseWriter, req *http.Request) {
	cook, err := req.Cookie("CookieName")
	if err != nil {
		log.Println("Cookie not found")
		return
	}
	cook.MaxAge = -1
	http.SetCookie(res, cook)

	fmt.Fprintln(res, "your cookie 'CookieName' has been expired")
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.HandleFunc("/visits", visitCounter)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}