#!/bin/bash

# saner programming env: these switches turn some bugs into errors
set -o errexit -o pipefail -o noclobber -o nounset

# -allow a command to fail with !’s side effect on errexit
# -use return value from ${PIPESTATUS[0]}, because ! hosed $?
! getopt --test > /dev/null
if [[ ${PIPESTATUS[0]} -ne 4 ]]; then
    >&2 echo 'I’m sorry, `getopt --test` failed in this environment.'
    exit 1
fi

OPTIONS=sdahn:p:v
LONGOPTS=server,disable-anonymous,set-admin-password,help,new-user-name:,new-user-pass-from:,verbose

# -regarding ! and PIPESTATUS see above
# -temporarily store output to be able to check for errors
# -activate quoting/enhanced mode (e.g. by writing out “--options”)
# -pass arguments only via   -- "$@"   to separate them correctly
! PARSED=$(getopt --options=$OPTIONS --longoptions=$LONGOPTS --name "$0" -- "$@")
if [[ ${PIPESTATUS[0]} -ne 0 ]]; then
    # e.g. return value is 1
    #  then getopt has complained about wrong arguments to stdout
    exit 2
fi
# read getopt’s output this way to handle the quoting right:
eval set -- "$PARSED"

function usage {
	cat <<EOM
Usage: $(basename "$0") [OPTION]...

  -s,--server             VALUE  Stardog server URL. Uses \$STARDOG_SERVER_URL if unset.
  -d,--disable-anonymous         Disable the anonymous user.
  -a,--set-admin-password        Sets the admin password. Needs \$STARDOG_ADMIN_PW set as ENV var.
  -v,--verbose                   Verbose mode.
  -n,--new-user-name      VALUE  Adds a new user name with the given name. Needs -p.
  -p,--new-user-pass-from PATH   When using --new-user-name, set a path to a file which contains the new user password.
  -h,--help                      Display this help.
EOM
	exit 2
}

admin_bin="${STARDOG_ADMIN_BIN:-/opt/stardog/bin/stardog-admin}"
admin_pw="${STARDOG_ADMIN_PW}"
server="${STARDOG_SERVER_URL:-http://localhost:5820}"

disable_anonymous=false
set_admin_password=false
verbose=false
new_user=false

# now enjoy the options in order and nicely split until we see --
while true; do
    case "$1" in
        -s|--server)
            server="${2}"
            shift 2
            ;;
        -d|--disable-anonymous)
            disable_anonymous=true
            shift
            ;;
        -a|--set-admin-password)
            set_admin_password=true
            shift
            ;;
        -h|--help)
            usage
            exit 1
            ;;
        -n|--new-user-name)
            new_user=true
            new_user_name="${2}"
            shift 2
            ;;
        -p|--new-user-pass-from)
            new_user_pass_from="${2}"
            shift 2
            ;;
        -v|--verbose)
            verbose=true
            shift
            ;;
        --)
            shift
            break
            ;;
        *)
            >&2 echo "Programming error"
            exit 3
            ;;
    esac
done

log() {
    if $verbose; then
        echo $@
    fi
}

if $disable_anonymous; then
    log "Disable anonymous user"
    ${admin_bin} --server "${server}" user disable anonymous
fi

if $set_admin_password; then
    log "Setting admin password"
    ${admin_bin} --server "${server}" user passwd --new-password "${admin_pw}"
fi

if $new_user; then
    log "Adding user ${new_user_name} to stardog"
    ${admin_bin} --server "${server}" user add --new-password $(cat "${new_user_pass_from}") --passwd "${admin_pw}" "${new_user_name}" || \
    log "User already exists - updating password" && \
    ${admin_bin} --server "${server}" user passwd --new-password $(cat "${new_user_pass_from}") --passwd "${admin_pw}" "${new_user_name}"
fi
