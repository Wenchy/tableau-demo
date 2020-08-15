#!/bin/bash

# set -eux
set -e
set -o pipefail

output_dir="../"

test_dir="../protobuf"
for item in "$test_dir"/* ; do
    echo "$item"
    if [ -f "$item" ]; then
        protoc -I"$test_dir" --go_out="$output_dir" "$item"
    fi
done

testpb_dir="../github.com/Wenchy/tableau-demo/testpb"

# update testpb
rsync -avz "$testpb_dir" ../

# remove
rm -rf "../github.com"