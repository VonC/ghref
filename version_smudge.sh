#!/bin/sh
sed -e 's;\([^ x]\)";\1xx";' $1
