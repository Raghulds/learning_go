#!/bin/bash

for i in "$@"; do
    (sleep "$i" && echo "$i")&
done
wait