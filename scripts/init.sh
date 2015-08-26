#!/usr/bin/env bash
`zenv system --cd-after $(pwd)`
function zenv_cd(){
    local before=`pwd`
    builtin cd "$@"
    local after=`pwd`
    `zenv system --cd "$before" "$after"`
}
alias cd="zenv_cd"
