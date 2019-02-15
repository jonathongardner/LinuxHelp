#!/bin/bash

cd "$( dirname "$0" )"
SCRIPT_DIR="$( dirs -c; dirs )"
MY_BASH="$SCRIPT_DIR/.jonathon_bash_profile"

if [ -f $(eval echo $MY_BASH) ]; then
  echo "" >> ~/.bashrc
  echo "source $MY_BASH" >> ~/.bashrc
else
  echo "ERROR: Could not find bash ($MY_BASH)."
fi

