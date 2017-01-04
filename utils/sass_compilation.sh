#!/bin/bash

src=$1
dest=$2
mode=${3:-none}
watching=${4:-0}

if [[ $watching -eq 0 ]]
then
	compass="compass watch"
else
	compass="compass compile"
fi

if [[ $mode = "none" ]]
then
	echo "Mode de Compilation : Debug(d) ou Production(P)"

	read mode

	case "$mode" in
	  d) $compass --sass-dir "$src" --force -s expand "$dest";;
	  *) $compass --sass-dir "$src" --force -s compressed --no-line-comments "$dest";;
	esac;
else
	[[ $mode -eq 0 ]] && "$compass" --sass-dir "$src" --force -s expand "$dest" || $compass --sass-dir "$src" --force --no-line-comments "$dest"
fi
