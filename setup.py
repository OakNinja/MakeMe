import os
from setuptools import setup


# Utility function to read the README file.
# Used for the long_description.  It's nice, because now 1) we have a top level
# README file and 2) it's easier to type in the README file than to put a raw
# string in below ...
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
    packages=['makeme', 'tests'],
    long_description=read('README'),
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Topic :: Utilities",
        "License :: MIT License",
    ],
)
