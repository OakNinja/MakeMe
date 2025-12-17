# MakeMeGo shell integration for Zsh
mm() {
  local selected_command
  selected_command=$(command mm --print-command "$@")
  if [ -n "$selected_command" ]; then
    BUFFER="$selected_command"
    CURSOR=${#BUFFER}
  fi
}
