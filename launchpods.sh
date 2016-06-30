#!/bin/bash

for i in `seq 1 $1`;
do
  sudo rkt run --insecure-options=image,ondisk --stage1-name=coreos.com/rkt/stage1-coreos docker://gcr.io/google_containers/pause:2.0 &
done

function cleanup {
  sudo rkt gc --grace-period=0 --expire-prepared=0
}

trap cleanup EXIT
