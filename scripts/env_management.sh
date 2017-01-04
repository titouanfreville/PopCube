#!/bin/bash
##
##
function set_env () {
	clear
  docker pull titouanfreville/whiptails:1.0
  if [ ! -f .env ]
    then
      cp -f .env.dist .env
  fi

  cp -f .env .env.old

  docker-compose -f docker-compose.whiptails.yml run whiptails
  RETURN_CODE="$?"
  docker-compose -f docker-compose.whiptails.yml rm
  if [ $RETURN_CODE -eq 1 ]
  then
    echo -e "$red An error occur during setup. Please correct it before running again.$basic"
    exit 1
  else
    echo -e "$basic ################################################################################"
    echo
    echo -e "$blue Setting done. Now checking that it will work.$basic"
  fi
  echo
  clear
}

function test_env () {
	echo "Testing env"

}