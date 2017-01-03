#!/bin/bash
watcher () {
# script:  watch
# author:  Mike Smullin <mike@smullindesign.com>
# license: GPLv3
# description:
#   watches the given path for changes
#   and executes a given command when changes occur
# usage:
#   watch <path> <cmd...>
#
  path=$1
  shift
  cmd=$*
  sha=0
  update_sha() {
    sha=`ls -lR $path | sha1sum`
  }
  update_sha
  previous_sha=$sha
  build() {
    echo -en " building...\n\n"
    $cmd
    echo -en "\n--> resumed watching."
  }
  compare() {
    update_sha
    if [[ $sha != $previous_sha ]] ; then
      echo -n "change detected,"
      build
      previous_sha=$sha
    else
      echo -n .
    fi
  }
  trap build SIGINT
  trap exit SIGQUIT

  echo -e  "--> Press Ctrl+C to force build, Ctrl+\\ to exit."
  echo -en "--> watching \"$path\"."
  while true; do
    compare
    sleep 1
  done
}

check_file () {
	local home="/go"

	if [[ -f "${1}" ]]
	then
		go test -v ${1}
		exit $?
	fi

	for el in "${1}"/*
	do
		if [[ -d "${el}" ]]
		then
			case ${el} in
				*/bin*) ;;
				*/pkg*) ;;
				*.*) ;;
				*)
          echo "Testing : ${el}"
					go test -cover -v ${el}/*.go;
          failures=$[$failures+$?]
					check_file ${el}
				;;
			esac
		fi
	done
}

watching=${2:-0}
failures=0
CMD="check_file ${1}"

if [ $watching -eq 0 ]
then
	watcher /go $CMD
else
	$CMD
  exit $failures
fi

