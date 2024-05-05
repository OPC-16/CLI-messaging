package user

import "fmt"

type User struct {
    UserID              int
    Username            string
    ReceivedMessagesLog []string
    SentMessagesLog     []string
}

var TotalUsers []User

func CreateUsers() *[]User {

    user1 := User{
        UserID: 1,
        Username: "Manthan",
        ReceivedMessagesLog: []string{"test1"},
        SentMessagesLog: []string{},
    }

    user2 := User{
        UserID: 2,
        Username: "Radha",
        ReceivedMessagesLog: []string{"test2"},
        SentMessagesLog: []string{},
    }

    user3 := User{
        UserID: 3,
        Username: "John",
        ReceivedMessagesLog: []string{"test3"},
        SentMessagesLog: []string{},
    }

    TotalUsers = append(TotalUsers, user1, user2, user3)

    return &TotalUsers
}

// Fetches a single user based on the ID passed in
func FetchUser(userId int) *User {
    for _, user := range TotalUsers {
        if userId == user.UserID {
            return &user
        }
    }

    fmt.Printf("ERROR: User not found with %d UserID\n", userId)
    return nil
}

func FetchAllUsers() *[]User {
    return &TotalUsers
}
