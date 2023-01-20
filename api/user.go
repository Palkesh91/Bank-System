package api

import (
	"log"
	"net"
	"strconv"
	"test1/module"

	"github.com/google/uuid"
)

var Users = make([]*module.User, 0)

func CreateUser(args []string, conn net.Conn) {
	if len(args) < 3 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}

	id := uuid.New().String()
	amount, err := strconv.Atoi(args[2])
	if err != nil {
		log.Println(err)
	}
	tranfer := module.Transection{
		SenderID:   Bank.ID,
		RecieverID: id,
		Amount:     amount,
		Status:     "credit",
	}
	Trans := []module.Transection{tranfer}
	u := &module.User{
		ID:           id,
		Name:         args[1],
		Balance:      amount,
		Transections: Trans,
	}
	log.Print(u)
	Users = append(Users, u)
	UpdateFile()
}

func Balance(args []string, conn net.Conn) {
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

func UserTrans(args []string, conn net.Conn) {
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
func Transfer(senderid string, receiverid string, amount int, conn net.Conn) {
	flag := false
	for ind, item := range Users {
		if item.ID == senderid {
			if item.Balance >= amount {
				item.Balance = item.Balance - amount
				flag = !flag
				t := &module.Transection{
					SenderID:   senderid,
					RecieverID: receiverid,
					Amount:     amount,
					Status:     "debit",
				}
				item.Transections = append(item.Transections, *t)
				Users[ind].Balance = item.Balance
				Users[ind].Transections = item.Transections

			} else {
				conn.Write([]byte("> Insufficient Balance\n"))
			}

		}
	}
	if flag {
		for ind, item := range Users {
			if item.ID == receiverid {
				item.Balance = item.Balance + amount
				t := &module.Transection{
					SenderID:   senderid,
					RecieverID: receiverid,
					Amount:     amount,
					Status:     "credit",
				}
				item.Transections = append(item.Transections, *t)
				Users[ind].Balance = item.Balance
				Users[ind].Transections = item.Transections

				log.Println(item.Balance)
				conn.Write([]byte("> Transection Successfull\n"))
			}
		}
	}
	UpdateFile()
}

func Transections(args []string, conn net.Conn) {
	if len(args) < 2 {
		conn.Write([]byte("> Invalide command" + "\n"))
	}
	user_id := args[1]

	for _, item := range Users {
		if item.ID == user_id {
			trans := item.Transections
			Showtrans(trans, conn, 10)
		}
	}

}

func Showtrans(trans []module.Transection, conn net.Conn, cnt int) {
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
