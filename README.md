# MXProtocol Server

Texting tool: Postman

First check the ip address of your lora app server container.
Edit the url (ip address) in wallet_controller.go (GetWalletBalance).
run ./mxprotocol and ./lora-app-server
Use Postman to send the login request(post) to lora app server (http://xxx.xxx.xxx.xxx:8080/api/internal/login), and copy the jwt that return from lora app server.
Use Postman to send the getbalance request + jwt to m2m, then m2m will send the request + jwt to lora app server for verfication.
If jwt is correct then you will get the user details back (show on your terminal) and get the balance in the postman, otherwise you get errors.