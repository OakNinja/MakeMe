# MakeMe
import re
import questionary
import os

makefile_content = []
with open('/Users/esse/Projects/MMD/Makefile', 'rt') as makefile:
    makefile_content = makefile.readlines()

make_targets = []
for row in makefile_content:
    if not row.startswith('.'):
        make_target_match = re.search(r'^(.+?):', row)
        if make_target_match:
            make_targets.append(make_target_match.group(1))

answer = questionary.select(
    "Choose a target",
    choices=make_targets).ask()  # returns value of selection

print(f'make {answer}')
os.system(f'make {answer}')
