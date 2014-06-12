#!/usr/bin/expect

set timeout -1
spawn  scp  ../bin/WXServer  root@115.29.107.154:/data
expect "*password"
send "cd9e46d9\n"
interact
