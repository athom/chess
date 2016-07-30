#!/usr/bin/env bash
export app_name=yeerchesswebserver
sudo docker run -d --name $app_name -p 3000:3000 athom/chess /bin/bash -c "/go/bin/$app_name"
