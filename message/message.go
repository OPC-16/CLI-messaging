package message

import (
	"fmt"

	"github.com/OPC-16/CLI-messaging/user"
)

func SendMessage(sender, receiver int, msg string) {
    if msg == "" {
        // then do cat fact api call and put that string in msg var
    }

    senderUser := user.FetchUser(sender)
    receiverUser := user.FetchUser(receiver)

    if senderUser != nil && receiverUser != nil {
        senderUser.AddSentMessage(msg)
        receiverUser.AddReceivedMessage(msg)

        fmt.Printf("User %d received message from User %d: %s\n", receiver, sender, msg)
    }
}

func BroadcastMessage(msg string) {
    if msg == "" {
        // then do cat fact api call and put that string in msg var
    }

    allUsers := user.FetchAllUsers()
    for _, singleUser := range allUsers {
        singleUser.AddReceivedMessage(msg)
    }

    fmt.Printf("Message broadcasted to all users: %s\n", msg)
}

func LogMessages(userId int) {
    u := user.FetchUser(userId)

    if u != nil {
        fmt.Printf("\nFetching Message logs of %s (UserID: %d) ...\n", u.Username, u.UserID)

        fmt.Println("Sent Messages are:")
        u.PrintSentMessages()

        fmt.Println("Received Messages are:")
        u.PrintReceivedMessages()
    }
}
