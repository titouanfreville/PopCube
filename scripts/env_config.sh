#!/bin/bash
#################################################################################
#################################################################################
################################ CREATE_ENV_FILE ################################
#################################################################################
## This script provide an interactive solution to set correctly the .env file  ##
## required in SOCIETY's softgallery container tools.                          ##
#################################################################################
## Author: FREVILLE Titouan <titouanfreville@hotmail.fr>                       ##
#################################################################################
#################################################################################
# Local Variables --------------------------------------------------------------
# ### COLORS ### #
green="\\033[1;32m"
red="\\033[1;31m"
basic="\\033[0;39m"
bblue="\\033[1;34m"
blue="\\033[0;34m"
# ### ### #
# ------------------------------------------------------------------------------
# Source wipthails script ------------------------------------------------------
source "whiptails"
# ------------------------------------------------------------------------------
# Get variables to update ------------------------------------------------------
if [ ! -f .env ]
then
  cp -f .env.dist .env
fi

cp -f .env .env.old

source ".env.old"
OPTIONS_TABLE=(
  "DEBUG" "$DEBUG" "ON"
  "MYSQL_ROOT_PASSWORD" "$MYSQL_ROOT_PASSWORD" "OFF"
  "WATCHING" "$WATCHING" "ON"
  "API_PORT" "$API_PORT" "OFF"
  "COMPOSE_REQUIRE" "$COMPOSE_REQUIRE" "OFF"
  "DOCKER_REQUIRE" "$DOCKER_REQUIRE" "OFF"
  "PORT" "$PORT" "OFF"
)
RES=$(init_checklist "Setting ENVFILE For POPCUBE chat project" "PARAMETERS" "Select parameters to update -- space to select, enter to validate." ".env.old" $OPTIONS_TABLE 20 80)
LIST="${RES//\"}"
# ------------------------------------------------------------------------------
# Operate modification ---------------------------------------------------------
read -r -a array <<< "$LIST"
for key in "${array[@]}"
do
  case "$key" in
    # REQUIRED MODIFICATIONS ------------------------------------------------------------------------------------------------------------------------
    # Debug setting
    # Call for a single input box wipthails so you can choose correct value for Debug variable.
    # ENSURE THAT PROVIDED Debug value is putted in .env
    "DEBUG" )
      OPTIONS=(
        0 "True"
        1 "False"
      )
      # Run wipthails script witch provide a selector for OPTIONS table.
      CHOISE=$(init_select "Update .env" "DEBUG" "Select true if you want debug output, false else." $OPTIONS 20 80)
      CHOISE=${CHOISE:-1}
      sed -i 's|DEBUG=.*|DEBUG='"$CHOISE"'|g' .env && ERROR=0 || ERROR=1
      ;;
    # Mysql Root Password
    # Call for a single input box wipthails so you can input a Password for the MariaDB container.
    # ENSURE THAT PROVIDED PASSWORD is putted in .env
    "MYSQL_ROOT_PASSWORD" )
      INPUT="$(init_inputbox "Update .env" "MYSQL_ROOT_PASSWORD" "Enter your Password" "$MYSQL_ROOT_PASSWORD" 20 80)"
      sed -i 's|MYSQL_ROOT_PASSWORD=.*|MYSQL_ROOT_PASSWORD='"$INPUT"'|g' .env && ERROR=0 || ERROR=1
      MYSQL_ROOT_PASSWORD=$INPUT
      ;;
    # Watching setting
    # Call for a single input box wipthails so you can choose correct value for Watching variable.
    # ENSURE THAT PROVIDED Watching value is putted in .env
    "WATCHING" )
      OPTIONS=(
        0 "True"
        1 "False"
      )
      # Run wipthails script witch provide a selector for OPTIONS table.
      CHOISE=$(init_select "Update .env" "WATCHING" "Select true if you want to watch, false else." $OPTIONS 20 80)
      CHOISE=${CHOISE:-0}
      sed -i 's|WATCHING=.*|WATCHING='"$CHOISE"'|g' .env && ERROR=0 || ERROR=1
      ;;
    # -----------------------------------------------------------------------------------------------------------------------------------------------
    # OTHERS ----------------------------------------------------------------------------------------------------------------------------------------
    # Api Port
    # Call for a single input box wipthails so you can input a Default port for the api container.
    # ENSURE THAT PROVIDED PORT is putted in .env for api
    "API_PORT" )
      INPUT="$(init_inputbox "Update .env" "API_PORT" "Enter your Password" "$API_PORT" 20 80)"
      sed -i "s|API_PORT=.*|API_PORT=$INPUT|g" .env && ERROR=0 || ERROR=1
      API_PORT=$INPUT
      ;;
    # Compose require
    # Call for a single input box wipthails so you can input a Minimal required version for docker-compose
    # ENSURE THAT PROVIDED MIN VERSION is putted in .env for DC
    "COMPOSE_REQUIRE" )
      INPUT="$(init_inputbox "Update .env" "COMPOSE_REQUIRE" "Enter your Password" "$COMPOSE_REQUIRE" 20 80)"
      sed -i "s|COMPOSE_REQUIRE=.*|COMPOSE_REQUIRE=$INPUT|g" .env && ERROR=0 || ERROR=1
      COMPOSE_REQUIRE=$INPUT
      ;;
    # Docker Require
    # Call for a single input box wipthails so you can input a inimal required version for docker.
    # ENSURE THAT PROVIDED MIN VERSION is putted in .env for Docker
    "DOCKER_REQUIRE" )
      INPUT="$(init_inputbox "Update .env" "DOCKER_REQUIRE" "Enter your Password" "$DOCKER_REQUIRE" 20 80)"
      sed -i "s|DOCKER_REQUIRE=.*|DOCKER_REQUIRE=$INPUT|g" .env && ERROR=0 || ERROR=1
      DOCKER_REQUIRE=$INPUT
      ;;
    # Port
    # Call for a single input box wipthails so you can input a Default port for the front container.
    # ENSURE THAT PROVIDED PORT is putted in .env for front
    "PORT" )
      INPUT="$(init_inputbox "Update .env" "PORT" "Enter your Password" "$PORT" 20 80)"
      sed -i "s|PORT=.*|PORT=$INPUT|g" .env && ERROR=0 || ERROR=1
      PORT=$INPUT
      ;;
     * )
      echo "DONE"
      ;;
    # -----------------------------------------------------------------------------------------------------------------------------------------------
  esac
done
# ------------------------------------------------------------------------------