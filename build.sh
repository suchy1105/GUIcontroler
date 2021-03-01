#!/bin/bash

docker stop gui
docker rm gui
docker build -t guisocket .
docker-compose -f run.yml up -d
