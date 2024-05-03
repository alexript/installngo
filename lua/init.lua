print("Install&Go examples")
print("\n-- Example: require from folder ------")
require("examples.require")
require("examples.hello")
require("examples.logging")
require("examples.osinfo")

print("\n-- Example: require from zipped bundle ------")
require("examples-zip")

print("\n-- use 3rdparty lua scripts. Catppuccin theme.")
local palette = require("catppuccin.macchiato")
print(palette.name)
require("catppuccin.term")
print("--- EndOfExamples ------")
