killall -9 go
nohup go run main.go >> $HOME/chess-server.log 2>&1 &
