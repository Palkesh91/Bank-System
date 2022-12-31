package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var users = make([]user, 0)

type admin struct {
	ID   string
	Name string
}

var bank = admin{
	ID:   "1",
	Name: "Bank",
}

type user struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Balance      int           `json:"Balance"`
	Transections []transection `json:"transections"`
}

type transection struct {
	SenderID   string `json:"sender"`
	RecieverID string `json:"reciver"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()

	log.Printf("server started on :8080")
	loadFile()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		log.Printf("client has connected %s", conn.RemoteAddr().String())
		loadFile()
		go readInput(conn)
	}

}

func readInput(conn net.Conn) {
	for {
		conn.Write([]byte("->Type "))
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/create":
			createUser(args, conn)
		case "/balance":
			Balance(args, conn)
		case "/transfer":
			userTrans(args, conn)
		case "/transections":
			transections(args, conn)
		case "/bankTransfer":
			bankTransfer(args, conn)
		case "/bankBalance":
			bankBalance(args, conn)
		case "/bankTrans":
			bankTrans(args, conn)
		default:
			conn.Write([]byte("Invalid command"))
		}

		//log.Print(users)
		//conn.Write([]byte("> " + msg + "\n"))
	}

}

func createUser(args []string, conn net.Conn) {
	if len(args) < 3 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}

	id := uuid.New().String()
	amount, err := strconv.Atoi(args[2])
	if err != nil {
		log.Println(err)
	}
	tranfer := transection{
		SenderID:   bank.ID,
		RecieverID: id,
		Amount:     amount,
		Status:     "credit",
	}
	trans := []transection{tranfer}
	u := user{
		ID:           id,
		Name:         args[1],
		Balance:      amount,
		Transections: trans,
	}
	log.Print(u)
	users = append(users, u)
	updateFile()
}

func Balance(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]
	for _, item := range users {
		if item.ID == user_id {
			amount := strconv.Itoa(item.Balance)
			conn.Write([]byte("> Balance is:" + amount + "\n"))
		}
	}
}

func userTrans(args []string, conn net.Conn) {
	if len(args) < 4 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	sId := args[1]
	rId := args[2]
	amount, err := strconv.Atoi(args[3])
	if err != nil {
		log.Println(err)
	}
	transfer(sId, rId, amount, conn)
}
func transfer(senderid string, receiverid string, amount int, conn net.Conn) {
	flag := false
	for ind, item := range users {
		if item.ID == senderid {
			if item.Balance >= amount {
				item.Balance = item.Balance - amount
				flag = !flag
				t := transection{
					SenderID:   senderid,
					RecieverID: receiverid,
					Amount:     amount,
					Status:     "debit",
				}
				item.Transections = append(item.Transections, t)
				users[ind].Balance = item.Balance
				users[ind].Transections = item.Transections

			} else {
				conn.Write([]byte("> Insufficient Balance\n"))
			}

		}
	}
	if flag {
		for ind, item := range users {
			if item.ID == receiverid {
				item.Balance = item.Balance + amount
				t := transection{
					SenderID:   senderid,
					RecieverID: receiverid,
					Amount:     amount,
					Status:     "credit",
				}
				item.Transections = append(item.Transections, t)
				users[ind].Balance = item.Balance
				users[ind].Transections = item.Transections

				log.Println(item.Balance)
				conn.Write([]byte("> Transection Successfull\n"))
			}
		}
	}
	updateFile()
}

func transections(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]

	for _, item := range users {
		if item.ID == user_id {
			trans := item.Transections
			showtrans(trans, conn, 10)
		}
	}

}

func showtrans(trans []transection, conn net.Conn, cnt int) {
	for _, item := range trans {
		if cnt == 0 {
			break
		}
		sID := item.SenderID
		rID := item.RecieverID
		amount := strconv.Itoa(item.Amount)
		status := item.Status
		cnt--
		conn.Write([]byte("> SenderID:" + sID + ", " + "RecieverID:" + rID + ", " + "Amount:" + amount + ", " + "Status:" + status + "\n"))
	}
}

func bankTransfer(args []string, conn net.Conn) {
	if len(args) < 4 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	sId := args[1]
	rId := args[2]
	amount, err := strconv.Atoi(args[3])
	if err != nil {
		log.Println(err)
	}
	transfer(sId, rId, amount, conn)
}

func bankBalance(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]
	for _, item := range users {
		if item.ID == user_id {
			amount := strconv.Itoa(item.Balance)
			conn.Write([]byte("> Balance is:" + amount + "\n"))
		}
	}
}

func bankTrans(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]

	for _, item := range users {
		if item.ID == user_id {
			trans := item.Transections
			showtrans(trans, conn, 50)
		}
	}

}

func updateFile() {
	content, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("userfile.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	loadFile()
}
func loadFile() {
	content, err := ioutil.ReadFile("userfile.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &users)
	if err != nil {
		log.Fatal(err)
	}
}
