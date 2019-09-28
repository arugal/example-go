#!/usr/bin/env bash

# insert
curl -X POST http://127.0.0.1:9000/dao/insert  -H "Content-Type:application/json" -d '{"name":"curl", "age":30}' -v

# update
curl http://127.0.0.1:9000/dao/update -X POST -H "Content-Type:application/json" -d '{"id":1, "name":"curl", "age":30}' -v

# select
curl http://127.0.0.1:9000/dao/select -v