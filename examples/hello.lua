print("Hello, Install&Go!")
print("Current platform is: " .. platform.OsName .. ", current architecture is: " .. platform.Arch)

local log = require("log")
local info = log.new()
info:printf("%s %d", "ok", 42)
