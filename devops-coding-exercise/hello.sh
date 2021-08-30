#! /bin/bash

# Using the scripting language of your choice, write a script that does the following:
# Take a path to a text file as an argument (e.g. hello.sh path/to/file.txt)
# Add “Hello World” to a new line at the beginning of the file
# Add “Goodbye” to a new line at the end of the file

filePath=($1)

sed -i '1s/^/Hello World \n/' $filePath

# sed ensures it's on a newline, 
# echo "Goodbye" >> file is an option too
sed -i '$ a Goodbye' $filePath