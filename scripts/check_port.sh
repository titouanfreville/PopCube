#!/bin/bash
# Local Variables --------------------------------------------------------------
# ### COLORS ### #
green="\\033[1;32m"
red="\\033[1;31m"
basic="\\033[0;39m"
blue="\\033[1;34m"
# ### ### #
# ------------------------------------------------------------------------------
source .env
# Check if 0.0.0.0:80 is free --------------------------------------------------
echo -e "$blue Checking if 0.0.0.0:$PORT is free ... $basic"
sudo netstat -tlnp |grep :$PORT
if [ $? -eq 0 ]
then
  echo -e "$red Please shut the process running.$basic"
  exit 1
else
  echo -e "$green You are Ok ;)$basic"
fi
echo
# ------------------------------------------------------------------------------