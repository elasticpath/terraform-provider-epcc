#!/bin/bash

RESOURCES=(
  v2/currencies
  v2/accounts
  pcm/pricebooks
  pcm/hierarchies
  pcm/products
  v2/customers
  v2/files
)



token=$(curl -X POST $EPCC_API_BASE_URL/oauth/access_token -d "client_id=$EPCC_CLIENT_ID" -d "client_secret=$EPCC_CLIENT_SECRET"     -d 'grant_type=client_credentials'  -s  2>/dev/null | jq -r .access_token);

for RESOURCE in ${RESOURCES[@]};
do

  while : ; do
    AMOUNT_REMAINING=$(curl -H "Authorization: Bearer $token" -H "EP-Beta-Features: $EPCC_BETA_API_FEATURES"  "$EPCC_API_BASE_URL/$RESOURCE" -s | jq -r '.data | length')
    echo "Processing $RESOURCE Currently has $AMOUNT_REMAINING remaining (if this number is greater than 25 there may be more)"


    if [[ AMOUNT_REMAINING -gt 0 ]]; then
      if [[ ("$RESOURCE" == "v2/currencies") && AMOUNT_REMAINING -le 1 ]]; then
        echo -e "\t => Only Default Currency Left (Probably)"
        break
      fi


      curl -H "Authorization: Bearer $token" -H "EP-Beta-Features: $EPCC_BETA_API_FEATURES"  "$EPCC_API_BASE_URL/$RESOURCE" -s | jq -r .data[].id | awk '{ system("sleep 0.1"); print $1 }' | xargs -d "\n"  -I '{}' curl -X DELETE -H "Authorization: Bearer $token" -H "EP-Beta-Features: $EPCC_BETA_API_FEATURES"  "$EPCC_API_BASE_URL/$RESOURCE/{}"
    else
      break
    fi
  done
done


find -iname "terraform.tfstate*" -delete
