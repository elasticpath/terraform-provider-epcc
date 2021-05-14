#!/bin/bash

token=$(curl -X POST $EPCC_API_BASE_URL/oauth/access_token -d "client_id=$EPCC_CLIENT_ID" -d "client_secret=$EPCC_CLIENT_SECRET"     -d 'grant_type=client_credentials'  -s | jq -r .access_token);

curl -X GET $EPCC_API_BASE_URL/$1 \
    -H "Authorization: Bearer $token" | jq -r

