#!/bin/bash

prompt() {
	read -p "$1"
}

cd hello

declare -a arr=(
"wc -l main.go"
"go run main.go"
"gomobile install -target=android"
)

for i in "${arr[@]}"
do
	prompt "\$ $i"
	$i
	echo
done
