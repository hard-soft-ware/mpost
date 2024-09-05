#!/bin/bash
# Original source: https://github.com/Bookshelf-Writer/scripts-for-integration/blob/main/_run/push-hook.sh
echo "[HOOK]" "Push"

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
script_dir="$script_dir/scripts"

OLD_VER=$(bash "$script_dir/sys.sh" -v)
VERSION=$(bash "$script_dir/sys.sh" -i -pa)

echo "Updated patch-ver:" "$OLD_VER >> $VERSION"
#############################################################################

go mod tidy
go mod vendor
bash "$script_dir/creator_const_Go.sh"


#############################################################################
exit 0

