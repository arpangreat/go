#!/bin/bash
ROOT_PATH=$(cd $(dirname $0) && pwd)

docker run -it -p 5050:8080 $1
