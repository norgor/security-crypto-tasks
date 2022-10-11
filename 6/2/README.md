# How to build
Ensure you have GNU Make and Go (>= 1.19) installed and simply run make. This will build both the server and client.

# How to run
There is a default user with username `test` and password `test`.\ in the users.json file.

For client: `./client <username> <password>`
For server: `./server`
Adding a user: `./server -add -username <username> -password <password>`
