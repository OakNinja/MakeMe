# MakeMe shell integration for Zsh
mm() {
  local selected_command
  selected_command=$(command mm --print-command "$@")
  if [ -n "$selected_command" ]; then
    print -z "$selected_command"
  fi
}
