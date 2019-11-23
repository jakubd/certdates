#!/usr/bin/env bash
# shellcheck disable=SC2164
cd "$(dirname "$0")"
 ./certdates --domains="domains.txt" --threshold=29
