# Steps to Run
### Prerequisite
1. protobuf 2.4.1
2. go modules

### Build
`./build.sh`

Binaries are generated in `./bin`

### Run Server
`./bin/server`

### Run Client
`./bin/client`

### Info
Client sends 100 messages to server and receives response.

It sends 100 Login messages. `./proto/auth.proto/Login`
and gets response `LoginResponse`

