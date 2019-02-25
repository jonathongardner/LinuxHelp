#!/bin/bash

cd "$( dirname "$0" )"
SCRIPT_DIR="$( dirs -c; dirs )"
MY_BASH="$SCRIPT_DIR/.jonathon_bash_profile"
FULL_PATH="$(pwd)"
MY_GIT_C="$FULL_PATH/.gitconfig"

if [ -f $(eval echo $MY_BASH) ]; then
  echo "" >> ~/.bashrc
  echo "source $MY_BASH" >> ~/.bashrc
  cp $MY_GIT_C ~/.gitconfig
else
  echo "ERROR: Could not find bash ($MY_BASH)."
fi

