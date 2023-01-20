package api

import (
	"log"
	"net"
	"strconv"
	"test1/module"
)

var Bank = &module.Admin{
	ID:   "1",
	Name: "Bank",
}

func BankTransfer(args []string, conn net.Conn) {
	if len(args) < 4 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	sId := args[1]
	rId := args[2]
	amount, err := strconv.Atoi(args[3])
	if err != nil {
		log.Println(err)
	}
	Transfer(sId, rId, amount, conn)
}

func BankBalance(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]
	for _, item := range Users {
		if item.ID == user_id {
			amount := strconv.Itoa(item.Balance)
			conn.Write([]byte("> Balance is:" + amount + "\n"))
		}
	}
}

func BankTrans(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]

	for _, item := range Users {
		if item.ID == user_id {
			trans := item.Transections
			Showtrans(trans, conn, 50)
		}
	}

}
