#!/bin/bash
# Variables --------------------------------------------------------------------
VERSION=0
# ### COLORS ### #
green="\\033[1;32m"
red="\\033[1;31m"
basic="\\033[0;39m"
blue="\\033[0;34m"
#bblue="\\033[1;34m"
# ### ### #
# ------------------------------------------------------------------------------
# Version check functions ------------------------------------------------------
version_comp () {
    if [[ "$1" == "$2" ]]
    then
        return 0
    fi
    local IFS=.
    local i ver1=($1) ver2=($2)
    # fill empty fields in ver1 with zeros
    for ((i=${#ver1[@]}; i<${#ver2[@]}; i++))
    do
        ver1[i]=0
    done
    for ((i=0; i<${#ver1[@]}; i++))
    do
        if [[ -z ${ver2[i]} ]]
        then
            # fill empty fields in ver2 with zeros
            ver2[i]=0
        fi
        if ((10#${ver1[i]} > 10#${ver2[i]}))
        then
            return 0
        fi
        if ((10#${ver1[i]} < 10#${ver2[i]}))
        then
            return 2
        fi
    done
    return 0
}

test_version_comp () {
    version_comp "$1" "$2"
    case $? in
        0) op='>=';;
        *) op='<';;
    esac
    if [[ $op = '<' ]]
    then
        echo -e "$red FAIL: Your version is older than require.,  '$1', '$2' $basic"
        return 1
    else
        echo -e "$green Pass: '$1 $op $2'. If greater, please report issue when not working. $basic"
        return 0
    fi
}
# -------------------------------------------------------------------------------
# Command Functions ------------------------------------------------------------
docker_ver () {
  echo "Checking requirement"
  echo -e "##################### INSTALLATIONS AND VERSIONS TESTS #########################"
  # Test Docker installation -----------------------------------------------------
  echo
  echo -e "$blue Checking docker installation .... $basic"
  VERSION=$(docker version --format '{{.Server.Version}}')
  if [ $? -eq 0 ]
  then
    test_version_comp "$VERSION" "$DOCKER_REQUIRE"
    if [ $? -eq 0 ]
    then
      echo -e "$green Docker well installed $basic"
    else
      echo -e "$red Please update Docker from https://docs.docker.com/engine/installation/ $basic"
      RETURN_CODE=1
    fi
  else
    echo -e "$red Please Install docker from https://docs.docker.com/engine/installation/ or make it run without sudo. Read the full doc :) $basic"
    RETURN_CODE=1
  fi
  echo
  VERSION=0
  # ------------------------------------------------------------------------------
  # Test Docker compose installation ---------------------------------------------
  echo -e "$blue Checking docker-compose installation .... $basic"
  VERSION=$(docker-compose version --short)
  if [ $? -eq 0 ]
  then
    echo
    test_version_comp "$VERSION" "$COMPOSE_REQUIRE"
    if [ $? -eq 0 ]
    then
      echo -e "$green Docker well installed $basic"
    else
      echo -e "$red Please update Docker compose from https://docs.docker.com/compose/install/ $basic"
      RETURN_CODE=1
    fi
  else
    echo -e "$red Please Install docker compose from https://docs.docker.com/compose/install/ $basic"
    RETURN_CODE=1
  fi
  echo
  VERSION=0
  # ------------------------------------------------------------------------------
  echo
  VERSION=0

  if [[ $RETURN_CODE -eq 1 ]]
  then
    echo -e "$red They are some problems with your installation, please fix it before trying again$basic"
    return 1
  else
    echo -e "$basic ################################################################################"
    echo
    echo -e "$blue Installation seems good. Setting up.$basic"
  fi
  echo
}
# ------------------------------------------------------------------------------
