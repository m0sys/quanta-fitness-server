#!/usr/bin/env bash

unset db
while getopts d: flag
do
    case "${flag}" in
        d) db=${OPTARG};;
    esac
done


if [ -z "$db" ]
then
    echo -e $"\nUsage: $(basename "$0") {d} - dbname\n"
    exit 1
fi


echo 'Openning connection to' $db'!'
docker exec -it qf_postgres psql -U root $db
