#!/bin/bash

./bench_wrapper -o --command="rkt api-service" | tee bench_api_service.log
