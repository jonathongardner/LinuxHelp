read -p "$1? " CONT
if [[ "$CONT" =~ [Yy] ]]; then
  exit 0;
else
  exit 1;
fi
