#!/bin/bash

pushd ..

grep -E "\(https://documentation.elasticpath.com/" *.go | sed -E 's/^.+\((https:[^)]+)\).+$/\1/g' | xargs -d "\n" -n 1 wget -x -P $(pwd)/docs/build/