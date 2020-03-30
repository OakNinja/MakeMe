import os
import re
from pathlib import Path

import questionary


def main():
    makefile_path = f'{os.getcwd()}/Makefile'
    if Path(makefile_path).exists():
        with open(makefile_path) as makefile:
            makefile_content = makefile.readlines()

        make_targets = set()
        for row in makefile_content:
            if not row.startswith('.') and not row.startswith('.PHONY:'):
                make_target_match = re.search(r'^(.+?):', row)
                if make_target_match:
                    make_targets.add(make_target_match.group(1))

        answer = questionary.select(
            "Choose a target",
            choices=sorted(make_targets)).ask()  # returns value of selection
        if answer:
            print(f'make {answer}')
            os.system(f'make {answer}')
    else:
        print('No Makefile found in current directory')


if __name__ == '__main__':
    main()
