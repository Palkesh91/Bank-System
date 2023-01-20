package input

import (
	"bufio"
	"net"
	"strings"
	"test1/api"
)

func ReadInput(conn net.Conn) {
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
			api.CreateUser(args, conn)
		case "/balance":
			api.Balance(args, conn)
		case "/transfer":
			api.UserTrans(args, conn)
		case "/transections":
			api.Transections(args, conn)
		case "/bankTransfer":
			api.BankTransfer(args, conn)
		case "/bankBalance":
			api.BankBalance(args, conn)
		case "/bankTrans":
			api.BankTrans(args, conn)
		default:
			conn.Write([]byte("Invalid command"))
		}

		//log.Print(users)
		//conn.Write([]byte("> " + msg + "\n"))
	}

}
