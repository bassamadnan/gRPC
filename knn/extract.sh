#!/bin/bash

extract_timing() {
    local file=$1
    local k=$(echo $file | cut -d'_' -f1)
    local dataset=$(echo $file | cut -d'_' -f2)
    local real=$(grep "real" $file | awk '{print $2}')
    local user=$(grep "user" $file | awk '{print $2}')
    local sys=$(grep "sys" $file | awk '{print $2}')
    echo "$dataset,$k,$real,$user,$sys"
}

echo "Dataset,K,Real,User,Sys"

for file in *_*_benchmark.txt; do
    extract_timing $file
done | sort -t',' -k1n -k2n
