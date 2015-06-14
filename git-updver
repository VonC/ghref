#!/bin/sh
v=$(git describe --all --long --abbrev=7)
git update-index --assume-unchanged -- version.go
n=$(git ls-files -m|wc -l)
git update-index --no-assume-unchanged -- version.go
if [[ ${n} -eq 0 ]]; then
  sed -i "s;\"[^\"]*\?\";\"${v}\";g" version.go
fi
