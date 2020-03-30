import os
from setuptools import setup, find_packages


def read(filename):
    return open(os.path.join(os.path.dirname(__file__), filename)).read()


setup(
    name="MakeMe",
    version="0.1.1",
    author="Esse Woods",
    author_email="esse.woods@gmail.com",
    description="Easing the usage of Makefiles",
    keywords="Makefile MakeMe MM",
    url="http://packages.python.org/makeme",
    install_requires=[
        'questionary',
    ],
    packages=['makeme', 'tests'],
    long_description=read('README'),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Development Status :: 3 - Alpha",
        "Operating System :: OS Independent",
        "Topic :: Utilities",
    ],
    python_requires='>=3.6',
    entry_points={
        'console_scripts': ['mm=makeme.command_line:main'],
    }
)
