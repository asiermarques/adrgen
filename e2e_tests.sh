#!/usr/bin/env bash

cd ./docs/features/features_definition_steps || exit
godog ../
cd ../../..