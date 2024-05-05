package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/OPC-16/CLI-messaging/message"
	"github.com/OPC-16/CLI-messaging/user"
)

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    fmt.Println("1. Send Message between two users")
    fmt.Println("2. Broadcast Message to all users")
    fmt.Println("3. View Message Log of a user")
    fmt.Println("4. Exit")

    for {
        io.WriteString(out, "Enter your choice: ")
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        line := scanner.Text()

        switch line {
            case "1":
                io.WriteString(out, "Enter sender ID: ")
                scanner.Scan()
                senderID, _ := strconv.Atoi(scanner.Text())

                io.WriteString(out, "Enter receiver ID: ")
                scanner.Scan()
                receiverID, _ := strconv.Atoi(scanner.Text())

                io.WriteString(out, "Enter message content: ")
                scanner.Scan()
                msg := scanner.Text()

                message.SendMessage(senderID, receiverID, msg)

            case "2":
                io.WriteString(out, "Enter message content: ")
                scanner.Scan()
                msg := scanner.Text()

                message.BroadcastMessage(msg)

            case "3":
                io.WriteString(out, "Enter UserID: ")
                scanner.Scan()
                userID, _ := strconv.Atoi(scanner.Text())

                message.LogMessages(userID)

            case "4":
                io.WriteString(out, "Printing all users' Message logs and exiting the application.")

                allUsers := user.FetchAllUsers()
                for _, user := range *allUsers {
                    message.LogMessages(user.UserID)
                }

                return

            default:
                io.WriteString(out, "Please enter a valid choice.\n")
        }
    }
}
