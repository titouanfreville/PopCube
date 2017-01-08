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
    sha=$(ls -lR "$path" | sha1sum)
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
	if [[ -f "${1}" ]]
	then
		go test -v "${1}"
		exit $?
	fi

	for el in "${1}"/*
	do
		if [[ -d "${el}" ]]
		then
			case ${el} in
				*/bin*) ;;
				*/pkg*) ;;
        */src) check_file "${el}";;
				*.*) ;;
				*)
          echo "Testing : ${el}"
				  go test -cover -v "${el}"/*.go;
          failures=$((failures+$?))
					check_file "${el}"
				;;
			esac
		fi
	done
}

check_fixed_packages () {
  echo "Testing FIXED PACKAGES : "
  echo ">>> API "
  go test -v -cover -covermode=count -coverprofile=/home/coverage/api.cover api
  failures=$((failures+$?))
  echo ">>> DATA_STORES "
  go test -v -cover -covermode=count -coverprofile=/home/coverage/data_stores.cover data_stores
  failures=$((failures+$?))
  echo ">>> MODELS "
  go test -v -cover -covermode=count -coverprofile=/home/coverage/models.cover models
  failures=$((failures+$?))
  echo ">>> UTILS "
  go test -v -cover -covermode=count -coverprofile=/home/coverage/utils.cover utils
  failures=$((failures+$?))
  echo "Generating coverage html reports"
  go tool cover -html=/home/coverage/api.cover -o /home/docs/api_cover.html
  go tool cover -html=/home/coverage/data_stores.cover -o /home/docs/data_stores_cover.html
  go tool cover -html=/home/coverage/models.cover -o /home/docs/models_cover.html
  go tool cover -html=/home/coverage/utils.cover -o /home/docs/utils_cover.html
  godoc -html cmd/api > /home/docs/api_documentation.html
  godoc -html cmd/data_stores > /home/docs/data_stores_documentation.html
  godoc -html cmd/models > /home/docs/models_documentation.html
  godoc -html cmd/utils > /home/docs/utils_documentation.html

}

watching=${2:-0}
failures=0
if [[ -d /home/coverage && -d /home/docs ]]
then
  CMD="check_fixed_packages"
else
  CMD="check_file ${1}"
fi

if [ "$watching" -eq 0 ]
then
	watcher /go "$CMD"
else
	"$CMD"
  exit "$failures"
fi

