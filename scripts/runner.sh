#!/bin/bash

function dockerexec () {
	local ex=$1
	local running_environement=${2:-dev}
	local detached=${3:-0}
	local compose_radical="docker-compose"
	local compose_detached=""
	if [[ $detached -eq 1 ]]
		then compose_detached="-d"
	fi

	case $running_environement in
		prod) compose_radical="$compose_radical -f docker-compose.yml -f docker-compose.prod.yml";;
		teste) compose_radical="$compose_radical -f docker-compose.yml -f docker-compose.test.yml";;
		*) ;;
	esac

	case $ex in
		1) $compose_radical build && $compose_radical run front $compose_detached;;
		2) $compose_radical build && $compose_radical run back $compose_detached;;
		4) $compose_radical build && $compose_radical run back-test $compose_detached;;
		3) $compose_radical build && { $compose_radical run front -d; $compose_radical run back -d; };;
		5) $compose_radical build && { $compose_radical run front -d; $compose_radical run back-test -d; };;
		6) $compose_radical build && { $compose_radical run back -d; $compose_radical run back-test -d; };;
		7) $compose_radical build && $compose_radical up $compose_detached;;
		*) ;;
	esac
}