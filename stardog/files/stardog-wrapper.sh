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

OPTIONS=hv
LONGOPTS=disable-user:,set-admin-password,help,ensure-user:,user-pass-from:,new-database:,ensure-role:,ensure-role-grants:,assign-role-to-user:,verbose

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

     --disable-user        USER                 Disable the given user.
     --set-admin-password                       Sets the admin password. Needs \$STARDOG_ADMIN_PW set as ENV var.
     --ensure-user         USER                 Ensures a user exists with the given name. Needs --user-pass-from flag.
                                                Accepts additional options after -- (e.g. '-- -s' for superuser)
     --user-pass-from      PATH                 When using --ensure-user, set a path to a file which contains the user password.
     --ensure-role         ROLE                 Ensures a role exists with the given name.
     --ensure-role-grants  ROLE,ACTION,OBJECT   Ensures an existing role has the given permissions and grant objects (e.g. role_name,write,db:*)
     --assign-role-to-user ROLE,USER
     --new-database        VALUE                Creates a new database with the given name. Pass additional options after --,
                                                e.g. "--new-database stardog -- -o search.enabled=true"
                                                Will skip existing databases.
  -h,--help                                     Display this help and exit.
  -v,--verbose                                  Verbose mode.
EOM
	exit 2
}

admin_bin="${STARDOG_ADMIN_BIN:-/opt/stardog/bin/stardog-admin}"
admin_pw="${STARDOG_ADMIN_PW}"

disable_user=false
set_admin_password=false
verbose=false
new_database=false
ensure_user=false
ensure_role=false
ensure_role_grants=false
assign_role=false
options=""

# now enjoy the options in order and nicely split until we see --
while true; do
    case "$1" in
           --disable-user)
            disable_user=true
            user_name="${2}"
            shift 2 ;;
          --set-admin-password)
            set_admin_password=true
            shift ;;
        -h|--help)
            usage
            exit 1 ;;
           --ensure-user)
            ensure_user=true
            user_name="${2}"
            shift 2 ;;
           --user-pass-from)
            user_pass_from="${2}"
            shift 2 ;;
           --new-database)
            new_database=true
            new_database_name="${2}"
            shift 2 ;;
           --ensure-role)
            ensure_role=true
            role_name="${2}"
            shift 2 ;;
           --ensure-role-grants)
            ensure_role_grants=true
            grant_arr=(${2//,/ })
            shift 2 ;;
           --assign-role-to-user)
            assign_role=true
            assign_arr=(${2//,/ })
            shift 2 ;;
        -v|--verbose)
            verbose=true
            shift ;;
        --)
            shift
            options="${@}"
            break ;;
        *)
            >&2 echo "Programming error"
            exit 3 ;;
    esac
done

log() {
    if $verbose; then
        echo $@
    fi
}

if $disable_user; then
    log "Disable user ${user_name}"
    ${admin_bin} user disable --passwd "${admin_pw}" ${user_name}
fi

if $set_admin_password; then
    log "Setting admin password"
    ${admin_bin} user passwd --new-password "${admin_pw}"
fi

if $new_database; then
    log "Checking if database ${new_database_name} exists..."
    ${admin_bin} db status --passwd "${admin_pw}" "${new_database_name}" || \
    ${admin_bin} db create --name "${new_database_name}" --passwd "${admin_pw}" ${options}
fi

if $ensure_user; then
    log "Adding user ${user_name} to stardog"
    ${admin_bin} user add --new-password $(cat "${user_pass_from}") --passwd "${admin_pw}" "${user_name}" ${options} || \
    log "User already exists - updating password" && \
    ${admin_bin} user passwd --new-password $(cat "${user_pass_from}") --passwd "${admin_pw}" "${user_name}" ${options}
fi

if $ensure_role; then
    add_command="${admin_bin} role add ${role_name} --passwd ${admin_pw}"
    log "Creating role ${role_name} if not existing"
    grep -E "(already exists)|(Successfully)" - < <(${add_command}) || ${add_command}
fi

if $ensure_role_grants; then
    role_name=${grant_arr[0]}
    action_name=${grant_arr[1]}
    grant_object=${grant_arr[2]}
    log "Granting '${action_name}' permission to role '${role_name}' for '${grant_object}'"
    grant_command="${admin_bin} role grant --passwd ${admin_pw} --action ${action_name} --object ${grant_object} ${role_name}"
    grep -E "(already has)|(Successfully)" - < <(${grant_command}) || ${grant_command}
fi

if $assign_role; then
    role_name=${assign_arr[0]}
    user_name=${assign_arr[1]}
    check_command="${admin_bin} role list --verbose --passwd ${admin_pw}"
    assign_command="${admin_bin} user addrole --role ${role_name} --passwd ${admin_pw} ${user_name}"
    log "Assigning role ${role_name} to user ${user_name}"

    grep -qE "(${role_name}).*(${user_name})" - < <(${check_command}) && log "user ${user_name} already has role ${role_name}" \
      || ${assign_command}

fi
