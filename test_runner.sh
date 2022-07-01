#!/bin/bash

set -eu

GOEXEC=$(which go)

cd "$(dirname "$0")"

for dir in ./*;
  do (
    if [ -d "$dir" ];then
      echo "\n*********"
      echo "checking for tests in $dir"

      # urgh dont judge. But slowing it down to make it viewable.
      sleep 0.5
      # please...dont judge.

      for file in $dir/*;
        do (
          if [["$file" == *"test"*]];then
            printf '%s\n' "$file"
            exec $GOEXEC test $dir
            break
          fi
        )
      echo "*********"
      done
    fi
  )
done

echo "\n*********"
echo "finished"
echo "*********"

        # cd $dir
        # echo "*********"
        # echo "checking for tests in ${PWD##*/}"
        # echo "*********"
        # exec $GOEXEC test ./${PWD##*/}