import os

from MakeMe.command_line import get_makefile_rows, get_makefile_targets


def test_load_makefile():
    makefile_path = f'{os.getcwd()}/sample_files/Makefile'
    rows = get_makefile_rows(makefile_path)
    assert rows


def test_get_makefile_targets():
    makefile_path = f'{os.getcwd()}/sample_files/Makefile'
    rows = get_makefile_rows(makefile_path)
    targets = get_makefile_targets(rows)
    assert len(targets) == 8


def test_get_makefile_targets_remove_phony_markers():
    rows = ['foo: ', '.PHONY: ']
    targets = get_makefile_targets(rows)
    assert len(targets) == 1


def test_get_makefile_targets_remove_non_targets():
    rows = ['foo: ', '.mcflurry: ']
    targets = get_makefile_targets(rows)
    assert len(targets) == 1


def test_get_makefile_targets_remove_duplicates():
    rows = ['foo: ', 'foo: ']
    targets = get_makefile_targets(rows)
    assert len(targets) == 1


def test_get_makefile_targets_matching_keyword():
    keyword = 'd'
    rows = ['dab: ', 'lol: ', 'daz: ']
    targets = get_makefile_targets(rows, keyword)
    assert len(targets) == 2


def test_get_makefile_targets_matching_keywords_sort_order():
    keyword = 'dev'
    rows = ['dev-c: ', 'lol: ', 'foo-dev: ', 'bar-dev: ', 'dev-b: ', 'dev-a: ']
    targets = get_makefile_targets(rows, keyword)
    assert targets == ['dev-a', 'dev-b', 'dev-c', 'bar-dev', 'foo-dev']


