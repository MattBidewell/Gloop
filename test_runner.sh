#!/bin/sh

set -eu

GOEXEC=$(which go)

cd "$(dirname "$0")"

for dir in ./*;
  do (
    if [ -d "$dir" ];then
      echo "checking for tests in $dir"

      # urgh dont judge. But slowing it down to make it viewable.
      sleep 0.5
      # please...dont judge.

      for file in $dir/*;
        do (
          if grep -q "test" "$file";then
            printf '%s\n' "$file"
            $GOEXEC test $dir
            break
          fi
        )
      done
    fi
  )
done

echo "finished"
        # echo "*********"
        # echo "checking for tests in ${PWD##*/}"
        # echo "*********"
        # exec $GOEXEC test ./${PWD##*/}
