package main

import (
	"bufio"
	"fmt"
	"log"
	"lorgus"
	"net"
	"os"
	"strings"
)

//server represents a chat server
type server struct {
     
	welcomemessage string
	listner net.Listner
	logger lorgus.StdLogger
	clients chan[*net.Conn]string
	broadcast chan string
	register chan *connection
	unregister chan *net.Conn
}


//represents a user connection 
type connection struct {

	conn *net.Conn 
	username string
}

func (s *server) run() {

}

func (s *server) listen() {

	for {

		conn, err = s.listner.Accept() // waits for somemone's connection
		if err != nil {
			s.logger.Println(fmt.Sprintf("connection failed: %v", err))
			continue
		}
		_, err = fmt.Fprintf(conn, s.welcomemessage + "\n")
		if err != nil {
            s.logger.Println(fmt.Sprintf("sending message failed ; %v", err))
			continue
		}
       
		//initialize a reader i/o in a buffered way 
		// buffer will do this for us 
		//conn i/o reader
		r := bufio.NewReader(conn) // min 27
		username, err := r.ReadString('\n')
		if err != nil {
			s.logger.Println(fmt.Sprintf("reading username failed ; %v", err))
		}
		// remove trailing suffix after in the username 
		username = strings.TrimSuffix(username, "\n")
		c := &connection{
			conn : &conn,
			username : username, // we need to remember the username
		}
		// channel send statement 
		// channel of pointers to connection
		//it's not buffered so we do block on sending to that until
		// someone's ready to read from it 
		s.register <- c

	}

		





	}
}


func ListenAndServe(addr string) (err error) {
	return ListenAndServeWithLogger(addr, log.New(os.Stderr, "", log.LstdFlags))
}

func ListenAndServeWithLogger(addr string, logger lorgus.StdLogger) (err error) {
	s := &server{
		welcomemessage: "Welcome to the Jungle Homies ! PLease enter your name : 	"
	}
    s.logger = logger
	s.register = make(chan *connection, 1)
	s.unregister = make(chan *net.Conn, 1)
	s.clients = map(chan[*net.Conn]string)
	s.broadcast = make(chan)
	
	s.listner, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.logger.Println(fmt.Sprint(" Listening on %v", addr))
	s.run()

	return nil

}