## Easing the usage of makefiles
[![Demo](https://github.com/OakNinja/MakeMe/raw/master/docs/demo.gif)](https://github.com/OakNinja/MakeMe/raw/master/docs/demo.gif)

**Requirements:**
`python>=3.6`

**Install:**
`pip install makeme`

**Fish integration:**
`fisher add OakNinja/MakeMeFish` then run with `mmf`

**Usage:**
type `mm` in a terminal window - if there is a Makefile in the current folder, all targets will be listed and selectable.

type `mm <keyword>` to filter on targets matching keyword, eg `mm dev`


***Other options:***

`mm -h` to display the help menu

`mm -l` to list all make targets

`mm -i` to run in interactive mode, after a command is executed, the targets will be listed again until the user quits. 


