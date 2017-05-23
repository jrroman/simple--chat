package main

import "os"
import "net"
import "fmt"
import "time"
import "bufio"
import "sync"

type Session struct {
	sid		int
	connection	net.Conn
	reader		*bufio.Reader
	writer		*bufio.Writer
	indata		chan string
	outdata		chan string
	sessionMutex	sync.Mutex
}

func NewSession(sid int, conn net.Conn) *Session {
	reader := *bufio.NewReader(conn)
	writer := *bufio.NewWriter(conn)

	session := &Session {
		sid: sid,
		connection: conn,
		reader: reader,
		writer: writer,
		indata: make(chan string)
		outdata: make(chan string)
	}
	fmt.Printf("New Host connected, session id: %d\n ", sid)
	return session
}

func (s *Session) Read() {
	for {
		str, err := s.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read string error")
			time.Sleep(100 * time.Millisecond)
			continue
		}
		s.indata <- str
	}
}

func (s *Session) Write() {
	for {
	}
}

func main() {
	var port string = "3490"
	servListener, listenErr := net.Listen("tcp", ":" + port)
	if listenErr != nil {
		fmt.Println("Listen server error")
		os.Exit(1)
	}

	fmt.Println("Server listening on port " + port)
	for {
		conn, connErr := servListener.Accept()
		if connErr != nil {
			fmt.Println("Error connecting to server")
			conn.Close()
			time.Sleep(100 * time.Millisecond)
			continue
		}

		go handleConn(conn)
	}
}
