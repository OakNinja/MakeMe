import os
from setuptools import setup, find_packages


def read(filename):
    return open(os.path.join(os.path.dirname(__file__), filename)).read()


setup(
    name="MakeMe",
    version="0.1.0",
    author="Esse Woods",
    author_email="esse.woods@gmail.com",
    description=("Easing the usage of Makefiles"),
    license="MIT",
    keywords="Makefile MakeMe MM",
    url="http://packages.python.org/makeme",
    install_requires=[
        'questionary',
    ],
    packages=['makeme', 'tests'],
    long_description=read('README'),
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Topic :: Utilities",
        "License :: MIT License",
    ],
    entry_points={
        'console_scripts': ['mm=makeme.command_line:main'],
    }
)
