# LinuxWC - Recreation of 'wc' Linux command

## Requirements

The following is some guidance to get you started in thinking about the behavior of wc:
 
All options must precede all input filenames.
 
If no input files are given as command-line arguments, wc defaults to standard input.
 
wc always writes to standard output.
 
Options can be given individually and in any order (e.g., -m -l or -l -m),  or combined (e.g., -lm or -wml).
 
The order in which the options are supplied has no effect on the order in which the counters are displayed.
The number of lines are always printed first, followed by the number of words and characters.
 
If no options are given, wc prints the number of lines, words, and characters (same as, e.g., wc -mwl).
 
If an invalid option or filename is given, your program must print the same error message wc would print to standard error (stderr) in that particular situation and halt with the same non-zero exit status.
