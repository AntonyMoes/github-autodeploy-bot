#!/bin/bash
git checkout master
git pull
cd ..
go build -o new_bot *.go
kill $(cat .bot.id)
mv new_bot bot
./bot &
cat $! > .bot.id
disown $!
