#!/bin/bash

docker stop GUIcontroll
docker rm GUIcontroll
make all
docker build -t guiimage .
#docker-compose -f run.yml up -d
