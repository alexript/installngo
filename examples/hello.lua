print("Hello, Install&Go!")
print("Current platform is: " .. platform.OsName .. ", current architecture is: " .. platform.Arch)

local log = require("log")
local info = log.new()
info:printf("%s %d", "ok", 42)
info:printf("Is it Windows: %s", tostring(platform.IsWindows()))
info:printf("Is it Linux: %s", tostring(platform.IsLinux()))
info:printf("Is it MacOS: %s", tostring(platform.IsMacOS()))
