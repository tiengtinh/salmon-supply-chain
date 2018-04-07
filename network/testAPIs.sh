#!/bin/bash

jq --version > /dev/null 2>&1
if [ $? -ne 0 ]; then
	echo "Please Install 'jq' https://stedolan.github.io/jq/ to execute this script"
	echo
	exit 1
fi

starttime=$(date +%s)

echo "========="
echo "Hey there! I'm fisherman of Fredrick"
echo "signing in"
echo
RESP_TOKEN=$(
curl -s -X POST \
  http://localhost:8080/signin \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'orgName=fredrick'
)
echo $RESP_TOKEN
TOKEN=$(echo $RESP_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "My token is $TOKEN"
echo
echo "I will now add 1 salmon data to ledger"
echo
RESP=$(
curl -s -X POST \
  http://localhost:8080/salmons \
  -H "authorization: Bearer $TOKEN" \
  -H "content-type: application/json" \
  -d '{
    "id": "101",
    "vessel": "Octopus Ship",
    "datetime": "2018-04-07T15:05:00+00:00",
    "location": "Ha Long Bay"
}'
)
echo $RESP
echo
echo "And I retrieve it just to check the information... you know"
echo
RESP=$(
curl -s -X GET \
  http://localhost:8080/salmons/101 \
  -H "authorization: Bearer $TOKEN" \
  -H "content-type: application/json"
)
echo $RESP
echo
echo "Fisherman of Fredrick out!"
echo



echo "============"
echo
echo "Hey there! It's member of restauranteur Alice. It's my turn now"
echo
echo "signing in"
echo
RESP_TOKEN=$(
curl -s -X POST \
  http://localhost:8080/signin \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'orgName=alice'
)
echo $RESP_TOKEN
TOKEN=$(echo $RESP_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "My token is $TOKEN"
echo
echo "The story is my restaurant has received the salmon id 101 from Fredrick. So.. it's blockchain time! don't wanna be a cheater! Let's declare our ownership with that 101 salmon with the network"
echo
RESP=$(
curl -s -X POST \
  http://localhost:8080/salmons/101/ownerships/claim \
  -H "authorization: Bearer $TOKEN" \
  -H "content-type: application/json" \
)
echo $RESP
echo
echo "coolio man! now let's check who that fish belong to (not that I don't trust the system, but just to be sure, you know)"
echo
RESP=$(
curl -s -X GET \
  http://localhost:8080/salmons/101 \
  -H "authorization: Bearer $TOKEN" \
  -H "content-type: application/json"
)
echo $RESP
echo
echo "Member of restauranteur Alice out!"
echo



echo "============"
echo

echo "Yo man! I'm a regulator here. It's my time to shine"
echo
echo "signing in"
echo
RESP_TOKEN=$(
curl -s -X POST \
  http://localhost:8080/signin \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'orgName=regulator'
)
echo $RESP_TOKEN
TOKEN=$(echo $RESP_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "My token is $TOKEN"
echo
echo "I'm a regulation and my job is to maintain regulation. That's what I do. Therefore, I will check all them salmons on the ledger right now and right there! gota query them all"
echo
RESP=$(
curl -s -X GET \
  http://localhost:8080/salmons \
  -H "authorization: Bearer $TOKEN" \
  -H "content-type: application/json"
)
echo $RESP
echo
echo "(Had a quick glance) Look like all is good! Regulator retires for the night!"
echo