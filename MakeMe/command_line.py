import os
import re
from pathlib import Path

import questionary


def main():
    makefile = f'{os.getcwd()}/Makefile'
    if Path(makefile).exists():
        with open(makefile) as makefile:
            makefile_content = makefile.readlines()

        make_targets = []
        for row in makefile_content:
            if not row.startswith('.') and not row.startswith('.PHONY:'):
                make_target_match = re.search(r'^(.+?):', row)
                if make_target_match:
                    make_targets.append(make_target_match.group(1))

        answer = questionary.select(
            "Choose a target",
            choices=make_targets).ask()  # returns value of selection
        if answer:
            print(f'make {answer}')
            os.system(f'make {answer}')
    else:
        print('No Makefile found in current directory')


if __name__ == '__main__':
    main()
