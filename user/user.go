package user

import "fmt"

type User struct {
    UserID              int
    Username            string
    ReceivedMessagesLog []string
    SentMessagesLog     []string
}

// slice of pointers to User struct
var TotalUsers []*User

// this will create 3 users and store their pointers in a slice and return that slice
func CreateUsers() []*User {

    user1 := User{
        UserID: 1,
        Username: "Manthan",
        ReceivedMessagesLog: []string{},
        SentMessagesLog: []string{},
    }

    user2 := User{
        UserID: 2,
        Username: "Radha",
        ReceivedMessagesLog: []string{},
        SentMessagesLog: []string{},
    }

    user3 := User{
        UserID: 3,
        Username: "John",
        ReceivedMessagesLog: []string{},
        SentMessagesLog: []string{},
    }

    TotalUsers = append(TotalUsers, &user1, &user2, &user3)

    return TotalUsers
}

// Fetches a single user based on the ID passed in
func FetchUser(userId int) *User {
    for _, user := range TotalUsers {
        if userId == user.UserID {
            return user
        }
    }

    fmt.Printf("ERROR: User not found with UserID %d\n", userId)
    return nil
}

func FetchAllUsers() []*User {
    return TotalUsers
}

func (u *User) AddReceivedMessage(msg string) {
    u.ReceivedMessagesLog = append(u.ReceivedMessagesLog, msg)
}

func (u *User) AddSentMessage(msg string) {
    u.SentMessagesLog = append(u.SentMessagesLog, msg)
}

func (u *User) PrintSentMessages() {
    for _, msg := range u.SentMessagesLog {
        fmt.Println("\t" + msg)
    }
}

func (u *User) PrintReceivedMessages() {
    for _, msg := range u.ReceivedMessagesLog {
        fmt.Println("\t" + msg)
    }
}
