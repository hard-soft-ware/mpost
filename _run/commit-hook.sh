#!/bin/bash
# Original source: https://github.com/Bookshelf-Writer/scripts-for-integration/blob/main/_run/commit-hook.sh
echo "[HOOK]" "Commit"

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
script_dir="$script_dir/scripts"

VERSION=$(bash "$script_dir/sys.sh" -v)
NAME=$(bash "$script_dir/git.sh" -b)

echo -e "$NAME [$VERSION] \n" $(cat "$1") > "$1"
#############################################################################

#Your code

#############################################################################
exit 0

