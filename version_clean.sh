#!/bin/sh
v=$(git describe --all --long --abbrev=7)
git update-index --assume-unchanged -- $1
n=$(git ls-files -m|wc -l)
git update-index --no-assume-unchanged -- $1
if [[ ${n} -gt 0 ]]; then
  sed -e "s;xx\";\";g" $1
else
  sed -e "s;\"[^\"]*\?\";\"${v}\";g" $1
fi
