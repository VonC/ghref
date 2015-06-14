#!/bin/sh
v=$(git describe --all --long --abbrev=7)
n=$(git ls-files -m|grep -v version.go| awk '{ print } END { print NR }')
if [[ ${n} -gt 0 ]]; then
  sed -e "s;xx\";\";g" $1
else
  sed -e "s;\"[^\"]*\?\";\"${v}\";g" $1
fi
