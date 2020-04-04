# Read from the file words.txt and output the word frequency list to stdout.
cat words.txt | tr ' ' '\n' | egrep '[a-z]' | sort | uniq -c | sort -n -r | sed -E 's/^ *([0-9]+) ([a-z]+)$/\2 \1/'
