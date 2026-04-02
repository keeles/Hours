# Top-level commands
complete -c hours -n "__fish_use_subcommand" \
  -a "add complete config delete get list new remove rm start stop task time version"

# Commands that take a client as 2nd arg
for cmd in new delete task get add remove rm complete start
  complete -c hours -n "__fish_seen_subcommand_from $cmd" \
    -a "(hours ls -c)"
end

# Commands that take client + task
for cmd in add remove rm complete start
  complete -c hours -n "__fish_seen_subcommand_from $cmd; and test (count (commandline -opc)) -eq 3" \
    -a "(hours get (commandline -opc)[2] -t)"
end

# Config subcommands
complete -c hours -n "__fish_seen_subcommand_from config" \
  -a "add-directory remove-directory list"

# Config client arg (except list)
complete -c hours -n "__fish_seen_subcommand_from config; and not __fish_seen_subcommand_from list" \
  -a "(hours ls -c)"
