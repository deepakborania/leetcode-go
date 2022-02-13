#!/bin/bash
NL=$'\n'

result="Leetcode solutions in Go lang\n\n| |Problem|\n|:---:|:---:|\n"
for d in [0-9][0-9][0-9][0-9]*/; do
    echo "Generating links for $d"

    IFS='_' read -r -a parts <<< "${d%/}"
    strpart="| [${parts[0]}](https://leetcode.com/problems/${parts[1]}) |[${parts[1]}]($d)|\n"
    # echo $strpart
    result=$result$strpart
done

echo -e $result > README.md
# echo -e $result