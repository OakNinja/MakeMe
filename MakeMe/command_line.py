import os
import sys
import re
from pathlib import Path

import questionary
from prompt_toolkit.styles import Style

custom_style_fancy = Style([
    ('qmark', 'fg:#673ab7 bold'),       # token in front of the question
    ('question', 'bold'),               # question text
    ('answer', 'fg:#f44336 bold'),      # submitted answer text behind the question
    ('pointer', 'fg:#673ab7 bold'),     # pointer used in select and checkbox prompts
    ('highlighted', 'fg:#673ab7 bold'), # pointed-at choice in select and checkbox prompts
    ('selected', 'fg:#cc5454'),         # style for a selected item of a checkbox
    ('separator', 'fg:#cc5454'),        # separator in lists
    ('instruction', ''),                # user instructions for select, rawselect, checkbox
    ('text', ''),                       # plain text
    ('disabled', 'fg:#858585 italic')   # disabled choices for select and checkbox prompts
])

make_target_pattern = re.compile(r'^([a-zA-Z0-9][^$#\/\t=]+?):.*$')


def get_makefile_rows(makefile_path):
    with open(makefile_path) as makefile:
        return makefile.readlines()


def get_makefile_targets(makefile_rows, keyword=None):
    make_targets = []
    if makefile_rows:
        for row in makefile_rows:
            make_target_match = re.search(r'^([a-zA-Z0-9][^$#\/\t=]+?):.*$', row)
            if make_target_match:
                make_target = make_target_match.group(1)
                if keyword and keyword.lower() not in make_target.lower():
                    continue  # Bail out if keyword is not matching the target
                make_targets.append(make_target)
    return make_targets


def generate_choices(makefile_targets):
    if makefile_targets:
        answer = questionary.select(
            "Choose a target",
            choices=makefile_targets,
            style=custom_style_fancy,
        ).ask()
        return answer


def main():
    if len(sys.argv) >= 2:
        keyword = sys.argv[1]
    else:
        keyword = None
    makefile_path = f'{os.getcwd()}/Makefile'
    if Path(makefile_path).exists():
        makefile_rows = get_makefile_rows(makefile_path)
        makefile_targets = get_makefile_targets(makefile_rows, keyword)
        if makefile_targets:
            answer = generate_choices(makefile_targets)
            if answer:
                print(f'make {answer}')
                os.system(f'make {answer}')
        else:
            print('No matching targets found in Makefile')
    else:
        print('No Makefile found in current directory')


if __name__ == '__main__':
    main()
