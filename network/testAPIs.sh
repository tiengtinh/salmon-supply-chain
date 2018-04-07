#!/bin/bash

jq --version > /dev/null 2>&1
if [ $? -ne 0 ]; then
	echo "Please Install 'jq' https://stedolan.github.io/jq/ to execute this script"
	echo
	exit 1
fi

starttime=$(date +%s)

echo "POST request signin as fisherman of Fredrick  ..."
echo
ORG1_TOKEN=$(curl -s -X POST \
  http://localhost:8080/signin \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'orgName=fredrick')
echo $ORG1_TOKEN
ORG1_TOKEN=$(echo $ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG1 token is $ORG1_TOKEN"

