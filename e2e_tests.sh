#!/usr/bin/env bash

cd ./docs/features/definitionsteps || exit
godog ../
cd ../../..