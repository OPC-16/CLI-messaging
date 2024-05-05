# CLI messaging application

It facilitates communication among multiple users via a central channel. Users can send and receive messages, targeting specific users or broadcasting to all users.

## Screenshots of interaction

#### sending a message
![sending a message](assets/send-message.png)

#### view message logs
![view message log](assets/view-messageLog.png)

#### all users' message logs while exiting
![broadcast and exit](assets/broadcast-and-exit.png)

## Key Points
* If a user sends an empty message, the application fetches a random fact from the [catfact API](https://catfact.ninja/fact) and uses it as the message content.
* The application stores user data in in-memory storage during runtime.
