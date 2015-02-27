
	#!/bin/bash


cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client.go $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client_p1_rmcgr1
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client_p1_rmcgr1
go install
cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server.go $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server_p1_rmcgr1
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server_p1_rmcgr1
go install
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/
rm ./submission_rmcgr1.tar
cp $GOPATH/bin/server_p1_rmcgr1 $GOPATH/bin/client_p1_rmcgr1 .
tar -czf submission_rmcgr1.tar client_p1_rmcgr1 server_p1_rmcgr1 config.txt
rm ./client_p1_rmcgr1 ./server_p1_rmcgr1
scp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/submission_rmcgr1.tar user@192.168.2.112:~
ssh user@192.168.2.112 'tar -xzf ~/submission_rmcgr1.tar'
