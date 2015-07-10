#!/bin/bash

prompt() {
	read -p "$1"
}

declare -a arr=(
"cd mypkg"
"tree"
"cat mypkg.go"
"export ANDROID_HOME=${ANDROID_HOME}"
"gomobile bind ."
"tree"
"unzip -l mypkg.aar"
)

for i in "${arr[@]}"
do
	prompt "\$ $i"
	$i
	echo
done

rm mypkg.aar
