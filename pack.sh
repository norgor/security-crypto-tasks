#!/bin/sh

echo "Packing task $1..."
echo "Enumerating task files..."
FILES=$(git ls-files | grep $1/ | tr '\n' ' ')
echo "Packing $(echo $FILES | tr -cd ' ' | wc -c) files..."
tar -czvf $1.tar.gz $FILES
