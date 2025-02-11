# myWc

Now we are able to find our files, but we may need more meta information about what is in these files. Let's implement a `wc`-like utility to collect basic statistics about our files.

First of all, let's assume that our files are utf-8 encoded text files, so your code should also work with text in Russian (forget about special cases like Arabic for now, only English and Russian are needed). You can also ignore punctuation and just use spaces as the only word delimiter.

You'll need to implement three mutually exclusive (only one can be specified at a time, otherwise you'll get an error message) flags for your code: `-l` for counting lines, `-m` for counting characters, and `-w` for counting words. Your program should be executable in this way:

```
# Count words in file input.txt
~$ ./myWc -w input.txt
777 input.txt
# Count lines in files input2.txt and input3.txt
~$ ./myWc -l input2.txt input3.txt
42 input2.txt
53 input3.txt
# Count characters in files input4.txt, input5.txt and input6.txt
~$ ./myWc -m input4.txt input5.txt input6.txt
1337 input4.txt
2664 input5.txt
3991 input6.txt
```

As you can see, the answer is always a calculated number and a filename separated by a tab (`\t`). If no flags are given, the `-w` behavior should be used.

**Important**: Since all files are independent, you should use goroutines to process them simultaneously. You can start as many goroutines as there are input files specified for the program.