#!/bin/bash

# tshark - the terminal wireshark
# 
# writes all streams to stdout

F=$1
if [ -z $1 ]; then
  exit 1
fi
END=$(tshark -r $F -T fields -e tcp.stream | sort -n | tail -1);
for ((i=0;i<=END;i++)); do 
  tshark -r $F -qz follow,tcp,ascii,$i 
done
