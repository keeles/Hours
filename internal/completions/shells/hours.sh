#!/bin/bash

_hours_completion() {
  local cur prev words cword
  _init_completion || return

  case "${words[1]}" in
    new | delete | task | get)
      case "$cword" in
        2)
          mapfile -t COMPREPLY < <(compgen -W "$(hours ls -c)" -- "$cur")
          ;;
      esac
      ;;
    add | remove | rm | complete | start)
      case "$cword" in
        2)
          mapfile -t COMPREPLY < <(compgen -W "$(hours ls -c)" -- "$cur")
          ;;
        3)
          local client="${words[2]}"
          mapfile -t COMPREPLY < <(compgen -W "$(hours get "$client" -t)" -- "$cur")
          ;;
      esac
      ;;
    config)
      case "$cword" in
        2)
          mapfile -t COMPREPLY < <(compgen -W "add-directory remove-directory list completion" -- "$cur")
          ;;
        3)
          local command="${words[2]}"
          if [ "$command" != "list" ]; then
            mapfile -t COMPREPLY < <(compgen -W "$(hours ls -c)" -- "$cur")
          fi
          ;;
      esac
      ;;
    *)
      mapfile -t COMPREPLY < <(compgen -W "add complete config delete get list new remove start stop task time version" -- "$cur")
      ;;
  esac
}

complete -r hours 2>/dev/null
complete -F _hours_completion hours
