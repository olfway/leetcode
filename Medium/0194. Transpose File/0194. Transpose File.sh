# Read from the file file.txt and print its transposed content to stdout.
for c in $(seq $(head -n 1 file.txt | wc -w))
do
    cut -f $c -d" " file.txt | tr '\n' ' '
    echo
done | sed 's/ $//'
