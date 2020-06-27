#!/usr/bin/expect

cd cmd/vault-api
GOOS=linux GOARCH=mips GOMIPS=softfloat go build
spawn scp ./vault-api root@192.168.8.1:/tmp
expect {
password: {send "$vault-pass\r"; exp_continue}
}
