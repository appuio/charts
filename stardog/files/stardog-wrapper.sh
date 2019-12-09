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

OPTIONS=sahn:p:dv
LONGOPTS=server,disable-user,set-admin-password,help,new-user-name:,new-user-pass-from:,new-database:,grant-user:,verbose

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

  -s,--server              VALUE                Stardog server URL. Uses \$STARDOG_SERVER_URL if unset.
    ,--disable-user        USER                 Disable the given user.
  -a,--set-admin-password                       Sets the admin password. Needs \$STARDOG_ADMIN_PW set as ENV var.
  -n,--new-user-name       USER                 Adds a new user name with the given name. Needs -p.
  -p,--new-user-pass-from  PATH                 When using --new-user-name, set a path to a file which contains the new user password.
  -p,--new-user-pass-from PATH   When using --new-user-name, set a path to a file which contains the new user password.
  -d,--new-database        VALUE                Creates a new database with the given name. Pass additional options after --, e.g. "--new-database stardog -- -o search.enabled=true"
  -h,--help                                     Display this help.
  -v,--verbose                                  Verbose mode.
EOM
	exit 2
}

admin_bin="${STARDOG_ADMIN_BIN:-/opt/stardog/bin/stardog-admin}"
admin_pw="${STARDOG_ADMIN_PW}"
server="${STARDOG_SERVER_URL:-http://localhost:5820}"

disable_user=false
set_admin_password=false
verbose=false
new_user=false
new_database=false
new_user=false
options=""

# now enjoy the options in order and nicely split until we see --
while true; do
    case "$1" in
        -s|--server)
            server="${2}"
            shift 2
            ;;
           --disable-user)
            disable_user=true
            user_name="${2}"
            shift 2
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
        -d|--new-database)
            new_database=true
            new_database_name="${2}"
            shift 2
            ;;
        -v|--verbose)
            verbose=true
            shift
            ;;
        --)
            shift
            options="${@}"
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

if $disable_user; then
    log "Disable user ${user_name}"
    ${admin_bin} --server "${server}" user disable ${user_name}
fi

if $set_admin_password; then
    log "Setting admin password"
    ${admin_bin} --server "${server}" user passwd --new-password "${admin_pw}"
fi

if $new_database; then
    log "Checking if database ${new_database_name} exists"
    ${admin_bin} --server "${server}" db status --passwd "${admin_pw}" "${new_database_name}" || \
    log "Database does not exist - creating database" && \
    ${admin_bin} --server "${server}" db create --name "${new_database_name}" --passwd "${admin_pw}" ${options}
fi

if $new_user; then
    log "Adding user ${new_user_name} to stardog"
    ${admin_bin} --server "${server}" user add --new-password $(cat "${new_user_pass_from}") --passwd "${admin_pw}" "${new_user_name}" || \
    log "User already exists - updating password" && \
    ${admin_bin} --server "${server}" user passwd --new-password $(cat "${new_user_pass_from}") --passwd "${admin_pw}" "${new_user_name}"
fi
