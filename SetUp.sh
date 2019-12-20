#!/bin/bash
cd "$( dirname "$0" )"
SCRIPT_DIR="$( dirs -c; dirs )"
MY_BASH="$SCRIPT_DIR/.jonathon_bash_profile"
MY_VIM="$SCRIPT_DIR/.jonathon_vimrc"
FULL_PATH="$(pwd)"
MY_GIT_C="$FULL_PATH/.gitconfig"

if [ -f $(eval echo $MY_BASH) ]; then
  if [ -f $(eval echo $MY_VIM) ]; then
    if [ -f $(eval echo $MY_GIT_C) ]; then
      # bash rc
      if [ "$1" = "-y" ] || ./script/prompt.sh "Source bash profile"; then
        echo "Sourcing $MY_BASH...":
        echo "" >> ~/.bashrc
        echo "source $MY_BASH" >> ~/.bashrc
      fi
      # vim rc
      if [ "$1" = "-y" ] || ./script/prompt.sh "Source vmrc"; then
        echo "Sourcing $MY_VIM..."
        echo "" >> ~/.vimrc
        echo "source $MY_VIM" >> ~/.vimrc
      fi
      # git config
      if [ "$1" = "-y" ] || ./script/prompt.sh "Source git config"; then
        echo "Including $MY_GIT_C..."
        echo "[include]" >> ~/.gitconfig
        echo "  path = $MY_GIT_C" >> ~/.gitconfig
      fi
    else
      echo "ERROR: Could not find git config ($MY_GIT_C)."
    fi
  else
    echo "ERROR: Could not find vimrc ($MY_VIM)."
  fi
else
  echo "ERROR: Could not find bash profile ($MY_BASH)."
fi
