#!/usr/bin/env bash
export PATH=~/.zenv/vbin:$PATH

zenv system --cd-after $$ $(pwd)
hash -r
function zenv_cd(){
    local before=`pwd`
    builtin cd "$@"
    local after=`pwd`
    zenv system --cd $$ "$before" "$after"
    hash -r
}
alias cd="zenv_cd"
