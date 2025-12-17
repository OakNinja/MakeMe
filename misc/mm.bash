# MakeMeGo shell integration for Bash
mm() {
  local selected_command
  selected_command=$(command mm --print-command "$@")
  if [ -n "$selected_command" ]; then
    READLINE_LINE="$selected_command"
    READLINE_POINT="${#selected_command}"
  fi
}
