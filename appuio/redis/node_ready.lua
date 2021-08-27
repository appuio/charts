local raw_state = redis.call("info", "replication")

local split = function(text, delim)
    return text:gmatch("[^"..delim.."]+")
end

local collect = function(iter)
    local elements = {}
    for s in iter do table.insert(elements, s); end
    return elements
end

local has_prefix = function(text, prefix)
    return text:find(prefix, 1, true) == 1
end

local replication_state = {}
for s in split(raw_state, "\r\n") do
    (function(s)
        if has_prefix(s,"#") then
            return
        end

        local kv = collect(split(s, ":"))
        replication_state[kv[1]] = kv[2]
    end)(s)
end

local isSlave = replication_state["role"] == "slave"
local isMasterLinkDown = replication_state["master_link_status"] == "down"
local isSyncing = replication_state["master_sync_in_progress"] == "1"

if isSlave and isMasterLinkDown then
    if isSyncing then
        return redis.error_reply("node is syncing")
    else
        return redis.error_reply("link to master down")
    end
end

return redis.status_reply("ready")
