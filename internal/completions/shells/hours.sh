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
          case "$command" in
            add-directory | remove-directory | add-dir | rm-dir)
              mapfile -t COMPREPLY < <(compgen -W "$(hours ls -c)" -- "$cur")
              ;;
          esac
          case "$command" in
            completion | comp)
              mapfile -t COMPREPLY < <(compgen -W "bash fish zsh" -- "$cur")
              ;;
          esac
          ;;
      esac
      ;;
    *)
      case $cword in
        1)
          mapfile -t COMPREPLY < <(compgen -W "add complete config delete get list new remove start stop task time version" -- "$cur")
          ;;
      esac
      ;;
  esac
}

complete -r hours 2>/dev/null
complete -F _hours_completion hours
