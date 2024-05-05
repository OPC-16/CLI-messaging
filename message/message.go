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
        senderUser.SentMessagesLog = append(senderUser.SentMessagesLog, msg)
        receiverUser.ReceivedMessagesLog = append(receiverUser.ReceivedMessagesLog, msg)

        fmt.Printf("User %d received message from User %d: %s\n", receiver, sender, msg)
    }
}

func BroadcastMessage(msg string) {
    if msg == "" {
        // then do cat fact api call and put that string in msg var
    }

    allUsers := user.FetchAllUsers()
    for _, user := range *allUsers {
        user.ReceivedMessagesLog = append(user.ReceivedMessagesLog, msg)
    }

    fmt.Printf("Message broadcasted to all users: %s\n", msg)
}

func LogMessages(userId int) {
    user := user.FetchUser(userId)

    if user != nil {
        fmt.Printf("Fetching Message logs of %s (UserID: %d) ...\n", user.Username, user.UserID)

        fmt.Println("Sent Messages are:")
        for _, msg := range user.SentMessagesLog {
            fmt.Println(msg)
        }

        fmt.Println("Received Messages are:")
        for _, msg := range user.ReceivedMessagesLog {
            fmt.Println(msg)
        }
    }
}
