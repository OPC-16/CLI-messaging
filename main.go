// Go-based command-line messaging application facilitating communication among multiple users via a central channel. Users can send and receive messages, targeting 
// specific users or broadcasting to all users.

package main

import (
	"os"

	"github.com/OPC-16/CLI-messaging/cli"
)

func main() {
    cli.Start(os.Stdin, os.Stdout)
}
