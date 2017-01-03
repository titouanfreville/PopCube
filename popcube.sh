#!/bin/bash

#################################################################################
#################################################################################
################################## POP CUBE  ####################################
#################################################################################
## This script is used to run the PopCube chat project. It has dev, test and   ##
## prod options. It is mostly here to ease project configuration and           ##
## deploiement                                                                 ##
#################################################################################
## Author: FREVILLE Titouan <tfreville@nexway.com>                             ##
#################################################################################
#################################################################################
# Local Variables --------------------------------------------------------------
RETURN_CODE=0
# ### COLORS ### #
green="\\033[1;32m"
red="\\033[1;31m"
basic="\\033[0;39m"
blue="\\033[1;34m"
# ### ### #
# ### VERSION REQUIRED ### #
ME=$(whoami)
# ### ### #
# ### OPTION DEFAULT VALUES ### #
quiet=1
interactive=0
debug=1
watching=1
env_file=
running_environment="dev"
# ### ### #
# ------------------------------------------------------------------------------
## Running process --------------------------------------------------------------
# ### Process Variables ### #
# Get options passed
TEMP=`getopt -o dhqye:f: --long --debug,help,non-interactive,quiet,env-file: -n 'Softgallery Tools Initialisation' -- "$@"`
# Help message to print for -h option (or when not providing correctly)
HELP_MESSAGE="Usage: ./**/init.sh [OPTIONS] [COMMANDS]
A quick way to configure and run PopCube's container tools.
Not providing command will default on executing all process.
This script require getopt to run.
Options:
  -d, --debug             Debug. Run the script with debug information.
  --dev 									Run for developemnt. (default)
  -e                      Environment variable in string format as : \"NAME=VALUE\".
  -f, --env-file          Name of the environment file to use in the script. Default : .env.
  -h, --help              Print this help.
  -p, --prod 							Run for production.
  -q, --quiet             Silencing scripts. Render testing and getting resources non interactive by default.
  -t, --test 							Run for testing only.
  -y, --non-interactive   Run process non interactively.
  -w, --watch             Make test task watch for modification.

Command:
  all             Running full stack for choosen env.
  back 						Running back.
  front           Running popcube front server.
  test 						Running test.
  test_docker     Check if you have well setted docker.
  test_env        Testing that you configure correctly the tool.
"

# Set tasks variables to false and quick_conf to interactive (dv : docker_verion, se : set_env, te : test_env, gr: ger_resources, i: interactivity)
dv=1; se=1; te=1; gr=1; i=0;
# ### ### #
################################################################################
##################### GETTING ARGS #############################################
# REQUIRE : getopt. Should be available on most bash terminal ##################
# ENSURE : args variable setted with value provided by user   ##################
################################################################################
[ $debug -eq 0 ] && echo "Provided arguments BEFORE Flag reading : $@"
eval set -- "$TEMP"
while true
do
  case "${1}" in
    -h|--help)
      echo "$HELP_MESSAGE"; exit 0;;
    -w|--watch)
      watching=0;shift;;
    -d|--debug)
      debug=0;shift;;
    -q|--quiet)
      quiet=0;shift;;
    -y|--non-interactive)
      interactive=1;shift;;
    -e)
      env_args+=($2);shift 2;;
    -f|--env-file)
      env_file=$2;shift 2;;
    --) shift; break;;
    *) echo "You provided a wrong option"; echo $HELP_MESSAGE; exit 1;;
  esac
done
################################################################################
##################### WELCOME :) ###############################################
# ./scripts/welcome.sh
################################################################################
##################### GETTING COMMANDS #########################################
# ENSURE : command var are setted to true if command was provided by user ######
# DEFAULT SETTING : If no command provided, will set all to true          ######
################################################################################
[ $debug -eq 0 ] && echo "Provided arguments AFTER flag reading : $@"
if [ $# -eq 0 ]
then
  dv=0; se=0; te=0; gr=0;
else
  for cmd in $@
  do
    [ $debug -eq 0 ] && echo "Read command : $cmd";
    case "$cmd" in
      test_docker)
        [ $debug -eq 0 ] && echo "Test-Docker command readed"
        dv=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $dv"
        ;;
      set_env)
        [ $debug -eq 0 ] && echo "Set-Env command readed"
        se=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $se"
        ;;
      test_env)
        [ $debug -eq 0 ] && echo "Test-Env command readed"
        te=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $te"
        ;;
      get_resources)
        [ $debug -eq 0 ] && echo "Get-Resources command readed"
        gr=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $gr"
        ;;
      *) echo "Unavailable command."; echo "$HELP_MESSAGE"; exit 1;;
    esac
  done
fi
################################################################################
##################### INTERRACTIVE SETUP #######################################
# ENSURE : Set up interactivity of the process according to user wishes ########
# DEFAULT SETTING : Process is interactive by default                   ########
################################################################################
if [ $interactive -eq 1 ]
  then
    [ $debug -eq 0 ] && echo "Non interractive mode activated."
    se=1;
    [ $debug -eq 0 ] && echo "Set environement should not be execute. Expect 1 ::::: Actual $se"
    i=1;
    [ $debug -eq 0 ] && echo "Interractive indice should be set to 1. Expect 1 ::::: Actual $se"
fi
################################################################################
##################### Add ENV_VAR from -e ######################################
# ENSURE : Env var provided from -e flags are added to wished env file #########
################################################################################
if [ ! -z $env_args ]
  then
    [ $debug -eq 0 ] && echo "Provided environement variable through -e tag. Values : ${env_args[@]}"
    for env in ${env_args[@]}
    do
      [ $debug -eq 0 ] && echo  "Readed arguments : $env"
      echo "$env" >> $env_file
      [ $debug -eq 0 ] && echo  "Argument should be add to $env_file ::::: cat $env_file $(cat $env_file)"
    done
fi
################################################################################
##################### ENV_FILE SETTING #########################################
# ENSURE : Setted correctly env file on the wishes one #########################
################################################################################
if [ ! -z $env_file ]
  then
    [ $debug -eq 0 ] && echo  "Env file provided. Setting docker compose to use this file"
    deb_tmp=$(sed "s/#env_file: __ENV_FILE_NAME__/env_file: $env_file/g" *.yml)
    sed -i "s/#env_file: __ENV_FILE_NAME__/env_file: $env_file/g" *.yml
    [ $debug -eq 0 ] && echo "Shoul have replaced #env_file:.... with env_file value. :::::: $deb_tmp"
fi
################################################################################
##################### RUNNING COMMAND ##########################################
# ENSURE : Running wishes command quietly if required ##########################
################################################################################
if [ $quiet -eq 1 ]
  then
    [ $debug -eq 0 ] && echo "Running script as verbose. You should see output."
    [ $dv -eq 0 ] && docker_ver
    [ $se -eq 0 ] && set_env
    [ $te -eq 0 ] && test_env $i
    [ $gr -eq 0 ] && get_resources $i
  else
    source "./scripts/spinner.sh"
    [ $debug -eq 0 ] && echo "Running script as quiet. You should see spinners."
    # Checking docker
    [ $dv -eq 0 ] && start_spinner "Checking docker version"
    [ $dv -eq 0 ] && docker_ver > /dev/null
    [ $dv -eq 0 ] && stop_spinner $? "Checking docker version"
    # Setting env
    [ $se -eq 0 ] && set_env
    # Testing env
    [ $te -eq 0 ] && start_spinner "Checking that environment is correctly setted"
    [ $te -eq 0 ] && test_env 1 > /dev/null
    [ $te -eq 0 ] && stop_spinner $? "Checking that environment is correctly setted"
    # Getting resources
    [ $gr -eq 0 ] && start_spinner "Getting resources"
    [ $gr -eq 0 ] && get_resources 1 > /dev/null
    [ $gr -eq 0 ] && stop_spinner $? "Getting resources"
fi
# ------------------------------------------------------------------------------