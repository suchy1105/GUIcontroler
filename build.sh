#!/bin/bash

docker stop GUIapi
docker rm GUIapi
make
docker build -t guisocket .
docker-compose -f run.yml up -d
