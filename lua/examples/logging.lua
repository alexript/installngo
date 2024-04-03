print("\n-- Example: logger usage ------")
local log = require("log")
local info = log.new()
info:printf("%s %d", "ok", 42)
