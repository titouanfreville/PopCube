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
# ### OPTION DEFAULT VALUES ### #
quiet=1
interactive=0
debug=1
env_file=
running_environment="dev"
# ### ### #
SCRIPT_EXECUTION_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source "$SCRIPT_EXECUTION_DIR/scripts/version_checking.sh"
source "$SCRIPT_EXECUTION_DIR/scripts/runner.sh"
source "$SCRIPT_EXECUTION_DIR/scripts/env_management.sh"
# ------------------------------------------------------------------------------
## Running process --------------------------------------------------------------
# ### Process Variables ### #
# Get options passed
TEMP=$(getopt -o dhqye:f: --long --debug,help,non-interactive,quiet,env-file: -n 'Softgallery Tools Initialisation' -- $*)
# Help message to print for -h option (or when not providing correctly)
HELP_MESSAGE="Usage: ./**/init.sh [OPTIONS] [COMMANDS]

A quick way to configure and run PopCube's container tools.
Not providing command will default on executing all process.

This script require getopt to run.
NOTE : On mac environement, install gnu-getopts.

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

Command:
  all             Running full stack for choosen env.
  back 						Running back.
  front           Running popcube front server.
  set_env         Set up the env file.
  test 						Running test.
  test_docker     Check if you have well setted docker.
  test_env        Testing that you configure correctly the tool.
"

# Set tasks variables to false and quick_conf to interactive (td : docker_verion, te : test_env, b=back, t=test, f=front i: interactivity)
td=1; se=1; te=1; b=1; f=1; t=1; i=0; dex=0;
# ### ### #
################################################################################
##################### GETTING ARGS #############################################
# REQUIRE : getopt. Should be available on most bash terminal ##################
# ENSURE : args variable setted with value provided by user   ##################
################################################################################
eval set -- "$TEMP"
echo "Provided arguments AFTER flag reading : $*"
echo "$#"
echo "$1"
while true
do
  case "${1}" in
    -h|--help)
      echo "$HELP_MESSAGE"; exit 0;;
    --dev)
      running_environment="dev"; shift;;
    -p| --prod)
      running_environment="prod"; shift;;
    -t| --test)
      running_environment="teste"; shift;;
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
    *) echo "You provided a wrong option"; echo "$HELP_MESSAGE"; exit 1;;
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
[ $debug -eq 0 ] && echo "Provided arguments AFTER flag reading : $*"
echo "Provided arguments AFTER flag reading : $*"
echo "$#"
echo "$1"
if [ $# -eq 0 ]
then
  [ $debug -eq 0 ] && echo "No command provided. Running all tasks.";
  td=0; se=0; te=0; b=0; f=0; t=0;
else
  for cmd in "$@"
  do
    [ $debug -eq 0 ] && echo "Read command : $cmd";
    case "$cmd" in
      all)
        [ $debug -eq 0 ] && echo "ALL command readed"
        b=0; f=0; t=0;
        [ $debug -eq 0 ] && echo "Should provide 0, 0, 0 :::: Actual : $b, $f, $t"
        ;;
      back)
        [ $debug -eq 0 ] && echo "Back command readed"
        b=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $b"
        ;;
      front)
        [ $debug -eq 0 ] && echo "Front command readed"
        f=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $f"
        ;;
      set_env)
        [ $debug -eq 0 ] && echo "Set-Env command readed"
        se=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $se"
        ;;
      test)
        [ $debug -eq 0 ] && echo "Test command readed"
        t=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $t"
        ;;
      test_docker)
        [ $debug -eq 0 ] && echo "Test-Docker command readed"
        td=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $td"
        ;;
      test_env)
        [ $debug -eq 0 ] && echo "Test-Env command readed"
        te=0
        [ $debug -eq 0 ] && echo "Should provide 0 :::: Actual : $te"
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
    i=1; se=1;
    [ $debug -eq 0 ] && echo "Interractive indice should be set to 1. Expect 1 ::::: Actual $i"
    [ $debug -eq 0 ] && echo "Set service have to be disabled. Expect 1 ::::: Actual $se"
fi
################################################################################
##################### Add ENV_VAR from -e ######################################
# ENSURE : Env var provided from -e flags are added to wished env file #########
################################################################################
if [ ! -z "$env_args" ]
  then
    [ $debug -eq 0 ] && echo "Provided environement variable through -e tag. Values : ${env_args[*]}"
    for env in "${env_args[@]}"
    do
      [ $debug -eq 0 ] && echo  "Readed arguments : $env"
      echo "$env" >> "$env_file"
      [ $debug -eq 0 ] && echo  "Argument should be add to $env_file ::::: cat $env_file $(cat "$env_file")"
    done
fi
################################################################################
##################### ENV_FILE SETTING #########################################
# ENSURE : Setted correctly env file on the wishes one #########################
################################################################################
if [ ! -z "$env_file" ]
  then
    [ $debug -eq 0 ] && echo  "Env file provided. Setting docker compose to use this file"
    deb_tmp=$(sed "s/#env_file: __ENV_FILE_NAME__/env_file: $env_file/g" -- *.yml)
    sed -i "s/#env_file: __ENV_FILE_NAME__/env_file: $env_file/g" -- *.yml
    [ $debug -eq 0 ] && echo "Shoul have replaced #env_file:.... with env_file value. :::::: $deb_tmp"
fi
################################################################################
##################### RUNNING COMMAND ##########################################
# ENSURE : Running wishes command quietly if required ##########################
################################################################################
if [ $quiet -eq 1 ]
  then
    [ $debug -eq 0 ] && echo "Running script as verbose. You should see output."
    td=0; te=0; b=0; f=0; t=0;
    [ $td -eq 0 ] && docker_ver
    [ $se -eq 0 ] && set_env
    [ $te -eq 0 ] && test_env $i
    [ $b -eq 0 ] && dex=$((dex+1))
    [ $f -eq 0 ] && dex=$((dex+2))
    [ $t -eq 0 ] && dex=$((dex+4))
    dockerexec $dex $running_environment $i
  else
    source "./scripts/spinner.sh"
    [ $debug -eq 0 ] && echo "Running script as quiet. You should see spinners."
    # Checking docker
    [ $td -eq 0 ] && start_spinner "Checking docker version"
    [ $td -eq 0 ] && docker_ver > /dev/null
    [ $td -eq 0 ] && stop_spinner $? "Checking docker version"
    # Testing env
    [ $te -eq 0 ] && start_spinner "Checking that environment is correctly setted"
    [ $te -eq 0 ] && test_env 1 > /dev/null
    [ $te -eq 0 ] && stop_spinner $? "Checking that environment is correctly setted"
    [ $b -eq 0 ] && dex=$((dex+1))
    [ $f -eq 0 ] && dex=$((dex+2))
    [ $t -eq 0 ] && dex=$((dex+4))
    start_spinner "Running correct stack"
    dockerexec $dex $running_environment 1
    stop_spinner $? "Running correct stack"
fi
# ------------------------------------------------------------------------------