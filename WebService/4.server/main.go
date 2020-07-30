package main

import (
	"time"
	"log"
	"fmt"
	"net"
	// "io"
	"bufio"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(100 * time.Second))

	request(conn)

	// scan := bufio.NewScanner(conn)
	// for scan.Scan() {
	// 	ln := scan.Text()
	// 	fmt.Println(ln)
	// 	fmt.Fprintf(conn, "I heard you say '%s\n'", ln)
	// }

	// io.WriteString(conn, "Hello from Server!\n")
	// fmt.Fprintln(conn, "How are you?")
	// fmt.Fprintf(conn, "%v", "I hope you are doing well!")

	// defer conn.Close()
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			fields := strings.Fields(ln)
			verb := fields[0]
			resource := fields[1]
			fmt.Println("METHOD** ", verb)
			fmt.Println("Resource* ", resource)
			response(conn, verb, resource)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn, verb, resource string) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title>
		</head><body><strong>Hello World</strong>
		` + verb + ` ` + resource + `</body></html>` 

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}