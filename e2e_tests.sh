#!/usr/bin/env bash

cd ./features/definitionsteps || exit
godog ../
cd ../../..