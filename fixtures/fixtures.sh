
apeupres_set_env() {
  if [ -z ${APEUPRES_NAME} ]; then
    APEUPRES_NAME="$1"
  else
    APEUPRES_NAME="$APEUPRES_NAME - $1"
  fi
  export APEUPRES_NAME
}

_ape_unset_bash() {
  OIFS=$IFS
  IFS=:
  for ape_item in $APEUPRES_TO_CLEAN_ENV; do
    unset $ape_item
  done
  IFS=$OIFS
}

_ape_unset_zsh() {
  for ape_item in ${(s/:/)APEUPRES_TO_CLEAN_ENV}; do
    unset $ape_item
  done
}

conf_unset() {
  ape_sh=`cat /proc/$$/comm`
  case "$ape_sh" in
    bash)
      _ape_unset_bash
      ;;
    zsh)
      _ape_unset_zsh
      ;;
  esac
  unset APEUPRES_NAME
  unset APEUPRES_TO_CLEAN_ENV
}

apeupres_set_clean_env() {
  if [ -z "$APEUPRES_TO_CLEAN_ENV" ]; then
    APEUPRES_TO_CLEAN_ENV="$@"
  else
    APEUPRES_TO_CLEAN_ENV="$APEUPRES_TO_CLEAN_ENV:$@"
  fi
  export APEUPRES_TO_CLEAN_ENV
}

