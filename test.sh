#!/bin/bash


cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client.go $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client_p1_rmcgr1
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client_p1_rmcgr1
go install
cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server.go $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server_p1_rmcgr1
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server_p1_rmcgr1
go install
rm $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/submission_rmcgr1.tar
tar -czf $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/submission_rmcgr1.tar $GOPATH/bin/client_p1_rmcgr1 $GOPATH/bin/server_p1_rmcgr1 $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/config.txt
scp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/submission_rmcgr1.tar user@192.168.2.112:~
ssh user@192.168.2.112 'tar -xzf ~/submission_rmcgr1.tar'
