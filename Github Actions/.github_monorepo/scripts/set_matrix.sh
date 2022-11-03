#!/bin/bash

#Author: Wade
#Date: 2022-Aug
#Description: To provide a json matrix for github actions

##init function
function init() {
    # init value
    JSON="{\"include\":["
    diff=$1
    dir=$2
}

# read path from diff files
function read_path () {
    # Loop by lines
    while read path; do
        # set parent folders of each kind of program as type
        type="$( echo "$path" | cut -d'/' -f1 -s )"
        if [ -z "$type" ]; then
            continue # Exclude root module
        fi

        if [ "$type"_flag == "$dir"_flag ]; then
            # Set sub dir as module in each parent folder
            module="$( echo $path | cut -d'/' -f2 -s )"
            # Add build to the matrix only if it is not already included
            JSONline="{\"module\": \"$module\", \"type\": \"$type\"},"
            if [[ "$JSON" != *"$JSONline"* ]]; then
                JSON="$JSON$JSONline"
            fi
        fi

    done <<< $diff
    JSON=$(modify_suffix "$JSON")
    echo $JSON
}

# Remove last "," and add closing brackets
function modify_suffix () {
    JSON=$1
    if [[ "$JSON" == *, ]]; then
      JSON="${JSON%?}"
    fi
    JSON="$JSON]}"
    echo "$JSON"
}

#main
function main() {
    init "$@"
    result=$(read_path)
    echo $result
}

main "$@"
