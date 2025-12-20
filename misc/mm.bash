# MakeMeGo shell integration for Bash
mm() {
  local selected_command
  selected_command=$(command mm --print-command "$@")
  if [ -n "$selected_command" ]; then
    history -s "$selected_command"
    eval "$selected_command"
  fi
}
