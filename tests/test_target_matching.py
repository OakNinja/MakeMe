import re
from pathlib import Path
import os

from MakeMe.command_line import get_makefile_rows, get_makefile_targets


def test_load_makefile():
    makefile_path = f'{os.getcwd()}/Makefile'
    rows = get_makefile_rows(makefile_path)
    assert rows


def test_get_makefile_targets():
    makefile_path = f'{os.getcwd()}/Makefile'
    rows = get_makefile_rows(makefile_path)
    targets = get_makefile_targets(rows)
    assert len(targets) == 17


def test_get_makefile_targets_matching_keyword():
    makefile_path = f'{os.getcwd()}/Makefile'
    keyword = 'blah'
    rows = get_makefile_rows(makefile_path)
    targets = get_makefile_targets(rows, keyword)
    assert len(targets) == 3