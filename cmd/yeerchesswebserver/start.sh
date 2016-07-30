#!/usr/bin/env bash
app_name=yeerchesswebserver

sudo docker rm -f $app_name

# debug
#sudo docker run -it \
            #-p 2999:3000 \
            #-v /Users/yeer/go/src:/go/src\
            #athom/chess \
            #/bin/bash
#exit

sudo docker run -d --name $app_name \
            -p 3000:3000 \
            -v /Users/yeer/go/src:/go/src\
            athom/chess \
            /bin/bash \
            -c "cd /go/src/github.com/athom/chess/cmd/$app_name; go build; go install; /go/bin/$app_name"

sudo docker logs -f $app_name
