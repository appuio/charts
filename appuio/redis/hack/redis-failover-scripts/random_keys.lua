local random_string = function(length)
    local res = ""
    for i = 1, length do
        res = res .. string.char(math.random(97, 122))
    end
    return res
end

-- Seeds random
-- https://redis.io/commands/eval#selective-replication-of-commands
redis.replicate_commands()

for _ = 1, 100000, 1 do
    local str = random_string(10)
    redis.call("SET", "RAND_"..str.."key", str);
end

return redis.status_reply("ok")
