#!/bin/bash

docker stop GUIcontroll
docker rm GUIcontroll
make all
docker build -t suchy11/guisocket:latest .
docker-compose -f run.yml up -d
