#!/usr/bin/env bash

cd ./docs/features/definitionSteps || exit
godog ../
cd ../../..