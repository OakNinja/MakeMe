import os
from setuptools import setup, find_packages


def read(filename):
    return open(os.path.join(os.path.dirname(__file__), filename)).read()


setup(
    name="MakeMe",
    version="0.2.1",
    author="Esse Woods",
    author_email="esse.woods@gmail.com",
    description="Easing the usage of Makefiles",
    long_description=read('README.MD'),
    keywords="Makefile MakeMe MM",
    url="https://github.com/OakNinja/MakeMe",
    project_urls={
        "Bug Tracker": "https://github.com/OakNinja/makeme/issues",
        "Documentation": "https://github.com/OakNinja/MakeMe",
        "Source Code": "https://github.com/OakNinja/MakeMe",
    },
    install_requires=[
        'questionary',
    ],
    packages=['makeme', 'tests'],
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
