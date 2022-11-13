#!/bin/bash

# Set secrets to fly.io
cat .env | tr '\n' ' ' | xargs flyctl secrets set