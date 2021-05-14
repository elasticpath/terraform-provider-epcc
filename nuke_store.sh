#!/bin/bash

RESOURCES=(
  v2/accounts
  pcm/hierarchies
  pcm/products
  v2/customers
)



token=$(curl -X POST $EPCC_API_BASE_URL/oauth/access_token -d "client_id=$EPCC_CLIENT_ID" -d "client_secret=$EPCC_CLIENT_SECRET"     -d 'grant_type=client_credentials'  -s | jq -r .access_token);

for RESOURCE in ${RESOURCES[@]};
do

  while : ; do
    AMOUNT_REMAINING=$(curl -H "Authorization: Bearer $token" -H "EP-Beta-Features: $EPCC_BETA_API_FEATURES"  "$EPCC_API_BASE_URL/$RESOURCE" -s | jq -r '.data | length')
    echo "Processing $RESOURCE Currently has $AMOUNT_REMAINING remaining (if this number is greater than 25 there may be more)"

    if [[ AMOUNT_REMAINING -ne 0 ]]; then
      curl -H "Authorization: Bearer $token" -H "EP-Beta-Features: $EPCC_BETA_API_FEATURES"  "$EPCC_API_BASE_URL/$RESOURCE" -s | jq -r .data[].id | awk '{ system("sleep 0.1"); print $1 }' | xargs -d "\n"  -I '{}' curl -X DELETE -H "Authorization: Bearer $token" -H "EP-Beta-Features: $EPCC_BETA_API_FEATURES"  "$EPCC_API_BASE_URL/$RESOURCE/{}"
    else
      break
    fi
  done





done