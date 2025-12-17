function mm --description 'Interactively select and run make targets'
  set -l selected_command (command mm --print-command $argv)
  if test -n "$selected_command"
    commandline -r "$selected_command"
  end
end
