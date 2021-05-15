#!/bin/bash

cat > docs.go << EOM
package epcc_docs

var DOCS_CSV=\`
EOM

grep -R -E "<tr>.+<code>.+</code>.*<code>.+</code>.+</tr>" ./build/ | sed -E "s|^([^:]+):.+>([^<]+)<.+>([^<]+)<.+>([^<]+)<.+$|'\1','\2','\3','\4'|g" >> docs.go

echo "\`" >> docs.go