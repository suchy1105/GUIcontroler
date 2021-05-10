#!/bin/bash

docker stop GUIapi
docker rm GUIapi
make all
docker build -t guisocket .
docker-compose -f run.yml up -d
