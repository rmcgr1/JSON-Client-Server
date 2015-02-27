#!/bin/bash

mkdir -p $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/temp

cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client.go $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client_p1_rmcgr1
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/client_p1_rmcgr1
go install

cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server.go $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server_p1_rmcgr1
cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/server_p1_rmcgr1
go install

cd $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/temp
rm $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/temp/submission_rmcgr1.tar

cp $GOPATH/bin/server_p1_rmcgr1 $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/temp
cp $GOPATH/bin/client_p1_rmcgr1 $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/temp
cp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/config.txt $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/temp


tar -czf submission_rmcgr1.tar client_p1_rmcgr1 server_p1_rmcgr1 config.txt
scp $GOPATH/src/github.com/rmcgr1/JSON-Client-Server/submission_rmcgr1.tar user@192.168.2.112:~
ssh user@192.168.2.112 'tar -xzf ~/submission_rmcgr1.tar'
