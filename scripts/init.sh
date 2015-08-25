#!/usr/bin/env bash
zenv system --init `pwd`
function zenv_cd(){
    local before = `pwd`
    cd $1
    local after = `pwd`
    zenv system --cd $before $after
}
alias cd='zenv_cd'
