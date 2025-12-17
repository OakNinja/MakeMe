#!/usr/bin/env fish

function install_mm
    set -q INSTALL_DIR; or set INSTALL_DIR /usr/local/bin

    echo "Building mm..."
    go build -o mm ./cmd/makemego

    echo "Installing mm to $INSTALL_DIR..."
    mkdir -p "$INSTALL_DIR"
    if test -w "$INSTALL_DIR"
        mv mm "$INSTALL_DIR/mm"
    else
        sudo mv mm "$INSTALL_DIR/mm"
    end

    echo "mm installed successfully!"

    echo "Installing Fish shell integration..."
    set -q XDG_CONFIG_HOME; or set XDG_CONFIG_HOME ~/.config
    set -l functions_dir "$XDG_CONFIG_HOME/fish/functions"
    mkdir -p "$functions_dir"

    set -l mm_function_file "$functions_dir/mm.fish"
    begin
        echo "function mm --description 'Interactively select and run make targets'"
        echo "  set -l selected_command (command mm --print-command \$argv)"
        echo "  if test -n \"\$selected_command\""
        echo "    commandline -r \"\$selected_command\""
        echo "  end"
        echo "end"
    end > "$mm_function_file"

    echo "Fish shell integration installed. Please restart your shell or run 'source $mm_function_file'."
end

install_mm