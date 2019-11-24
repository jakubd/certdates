#!/usr/bin/env bash
# shellcheck disable=SC2164
cert_handling "$(dirname "$0")"
 ./certdates --domains="domains.txt" --threshold=29
