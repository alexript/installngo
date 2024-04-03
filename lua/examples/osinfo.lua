print("\n-- Example: platform usage ------")
local platform = require("platform")
local log = require("log")
local info = log.new()

print("Current platform is: " .. platform.OsName .. ", current architecture is: " .. platform.Arch)

info:printf("Is it Windows: %s", tostring(platform.IsWindows()))
info:printf("Is it Linux: %s", tostring(platform.IsLinux()))
info:printf("Is it MacOS: %s", tostring(platform.IsMacOS()))
info:printf("Distributive: %s [%s]", platform.Distributive, platform.DistributiveVersion)
