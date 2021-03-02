#!/bin/bash

docker stop GUIapi
docker rm GUIapi
docker build -t guisocket .
docker-compose -f run.yml up -d
