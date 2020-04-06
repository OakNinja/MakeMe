import signal
import subprocess
import os
import sys
import re
from pathlib import Path

import questionary
from prompt_toolkit.styles import Style

style = Style([
    ('qmark', 'fg:#673ab7 bold'),        # token in front of the question
    ('question', 'bold fg:#268bd2'),                # question text
    ('answer', 'fg:#f44336 bold'),       # submitted answer text behind the question
    ('pointer', 'fg:#673ab7 bold'),      # pointer used in select and checkbox prompts
    ('highlighted', 'fg:#cb4b16 bold'),  # pointed-at choice in select and checkbox prompts
    ('selected', 'fg:#cc5454'),          # style for a selected item of a checkbox
    ('separator', 'fg:#cc5454'),         # separator in lists
    ('instruction', 'fg:#6c71c4'),       # user instructions for select, rawselect, checkbox
    ('text', 'fg:#d33682'),                 # plain text
    ('disabled', 'fg:#858585 italic')    # disabled choices for select and checkbox prompts
])

make_target_pattern = re.compile(r'^([a-zA-Z0-9][^$#\/\t=]+?):[^$#\/\t=].*$')


def get_makefile_rows(makefile_path):
    with open(makefile_path) as makefile:
        return makefile.readlines()


def get_makefile_targets(makefile_rows, keyword=None):
    make_targets = []
    if makefile_rows:
        for row in makefile_rows:
            make_target_match = re.search(make_target_pattern, row)
            if make_target_match:
                make_target = make_target_match.group(1)
                if keyword and keyword not in make_target.lower() or make_target in make_targets:
                    continue
                make_targets.append(make_target)
    if keyword:
        return sorted(sorted(make_targets, key=str.lower), key=lambda target: 0 if target.lower().startswith(keyword) else 1)
    return make_targets


def generate_choices(makefile_targets, keyword):
    base_choices = ['Quit MakeMe (ctrl+c)']
    if makefile_targets:
        title = f"Targets matching '{keyword}'" if keyword else 'Choose a target'
        answer = questionary.select(
            title,
            choices=base_choices + makefile_targets,
            style=style,
        ).ask()
        if not answer or answer == base_choices[0]:
            sys.exit()  # Bail out
        print(answer)
        if '%' in answer:
            follow_up_answer = questionary.text(
                f'Replace % in {answer} with:'
            ).ask()
            answer = answer.replace('%', follow_up_answer, 1)
        return answer


def process_makefile(makefile_path, keyword):
    makefile_rows = get_makefile_rows(makefile_path)
    makefile_targets = get_makefile_targets(makefile_rows, keyword)
    if makefile_targets:
        answer = generate_choices(makefile_targets, keyword)
        return answer


def call_make_target(target):
    try:
        subprocess.call(['make', target])
    except KeyboardInterrupt:
        pass
    except subprocess.CalledProcessError as e:
        print(f'{e.cmd} returned {e.returncode} with error: {e.output}')


def main():
    if len(sys.argv) >= 2:
        keyword = sys.argv[1].lower()
    else:
        keyword = None
    makefile_path = f'{os.getcwd()}/Makefile'
    if Path(makefile_path).exists():
        target = process_makefile(makefile_path, keyword)
        if target:
            print(f'make {target}')
            call_make_target(target)
        else:
            print('No matching targets found in Makefile')
    else:
        print('No Makefile found in current directory')


if __name__ == '__main__':
    main()
