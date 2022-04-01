// Jon Moran
// moranj13@udayton.edu
// CPS 444: UNIX/Linux Programming, University of Dayton
// Project 1
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	NOTICE:
This program is treated as if it were 'wc' in the command line ---
"wc" is not expected as an argument, this program takes its place in the command line

**When trying to exit standard input:
	If running on Windows, press Ctrl-Z-Enter to exit
*/

/*
	Requirements:

The program must be written in Go and compile/run without errors on a Linux system.

Your version of wc must behave exactly like the wc command installed on a Linux system in all aspects with the following exception.
You need only implement the -l, -w and -m options.
Your program must mime the behavior of wc on a Linux system and replicate it in your program (see the wc main page and experiment with the command until you completely understand its behavior).

The following is some guidance to get you started in thinking about the behavior of wc:

All options must precede all input filenames.

If no input files are given as command-line arguments, wc defaults to standard input.

wc always writes to standard output.

Options can be given individually and in any order (e.g., -m -l or -l -m),  or combined (e.g., -lm or -wml).

The order in which the options are supplied has no effect on the order in which the counters are displayed.
The number of lines are always printed first, followed by the number of words and characters.

If no options are given, wc prints the number of lines, words, and characters (same as, e.g., wc -mwl).

If an invalid option or filename is given, your program must print the same error message wc would print to standard error (stderr) in that particular situation and halt with the same non-zero exit status.
*/

func main() {
	argCount := len(os.Args[1:]) // store number of args

	if argCount == 0 { // if no arguments are given, default to standard input
		os.Args = append(os.Args, "-")
	}

	// DEBUG fmt.Println(os.Args[1])
	if string(os.Args[1]) == "wc" {
		fmt.Println("Don't include 'wc' as an argument for this program. Assume the program is acting as the argument 'wc' from Linux OS.")
		os.Exit(0)
	} else if string(os.Args[1]) == "--help" {
		fmt.Println(`Usage: wc [OPTION]... [FILE]...
	or:  wc [OPTION]... --files0-from=F
Print newline, word, and character counts for each FILE, and a total line if
more than one FILE is specified.  A word is a non-zero-length sequence of
characters delimited by white space.

With no FILE, or when FILE is -, read standard input.

The options below may be used to select which counts are printed, always in
the following order: newline, word, character.
-m,	--chars            	print the character counts
-l,	--lines	          	print the newline counts
	--files0-from=F		read input from the files specified by
					NUL-terminated names in file F;
					If F is - then read names from standard input
-w,	--words            	print the word counts
	--help     		display this help and exit

GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
Full documentation at: <http://www.gnu.org/software/coreutils/wc>
or available locally via: info '(coreutils) wc invocation'
	  `)
		os.Exit(0)
	}

	lBool, wBool, mBool := false, false, false
	var boolCount int
	for i := 1; i <= len(os.Args[1:]); i++ { // for every arg, decide what options are requested (cumulatively)
		arg := string(os.Args[i])
		if isOption(arg) {
			if strings.Contains(string(arg), "l") { // l was a requested option
				lBool = true
				boolCount++
			}
			if strings.Contains(string(arg), "w") { // w was a requested option
				wBool = true
				boolCount++
			}
			if strings.Contains(string(arg), "m") { // m was a requested option
				mBool = true
				boolCount++
			}
		}
	}

	// if no options are provide: default to use all options
	if boolCount == 0 {
		lBool, wBool, mBool = true, true, true
	}

	// for every arg see if its an invalid option, handle accordingling, also for each valid option track what options are requested
	for i := 1; i <= argCount; i++ {
		arg := string(os.Args[i])
		if !isOption(arg) && !isFile(arg) { // if it is an valid option or valid file
			if arg == "-" {
				continue
			} else if string(arg[0:1]) == "-" { // if first character of arg is a hyphen, check if its unrecognizable or invalid option
				arg = arg[1:]
				if string(arg[0:1]) == "-" { // if arg has at least two hypens in the beginning , it is not a recognizable option
					fmt.Println("wc: unrecognized option '" + arg + "'\nTry 'wc --help' for more information.")
					os.Exit(0)
				} else { // if arg has only one hyphen , it is an invalid option
					fmt.Println(string(arg[0:1]))
					fmt.Println("DEBUGGING1")
					fmt.Println("wc: invalid option -- '" + arg + "'\nTry 'wc --help' for more information.")
					os.Exit(0)
				}
			}
		}
		if isOption(arg) { //if arg is a valid option, make sure it is 'turned on' to be utilized
			if strings.Contains(string(arg), "l") { // l was a requested option
				lBool = true
			}
			if strings.Contains(string(arg), "w") { // w was a requested option
				wBool = true
			}
			if strings.Contains(string(arg), "m") { // m was a requested option
				mBool = true
			}
		}
	}

	var numFilesAndStdInputs int
	var numL, numW, numM, lTotal, wTotal, mTotal int
	for i := 1; i <= len(os.Args[1:]); i++ { // for every arg
		arg := string(os.Args[i])

		// if not valid/invalid option or a file...
		if !isOption(arg) && !isFile(arg) && arg != "-" {
			// already tested for invalid options cases
			// the arg is assumed to be a file that can't be found
			fmt.Println("wc: " + arg + ": No such file or directory")
		}

		//  PRINT IN THE FORMAT:
		// 5  40  380    FILE1.txt
		// 27  949	4700
		// 32  989  5080  total

		// if user types just a hyphen: take standard input and perform desired options on input
		if arg == "-" {
			var allInput = ""

			//take standard input
			//fmt.Printf("enter input: ")
			for !strings.Contains(string(allInput), string(`^D`)) { // loop until users types in ctrl+d command
				var reader = bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')

				allInput = allInput + input
				if err != nil && err != io.EOF {
					fmt.Println(err)
				}
				if err == io.EOF { // if standard input is exited
					break
				}
				//fmt.Println(io.EOF)
			}
			//fmt.Println(allInput) // DEBUG
			// l(allInput) // DEBUG
			numL, numW, numM = requestedOptions(lBool, wBool, mBool, allInput, argCount) // return number of lines, words, characters from the file
			//fmt.Println(numL) //DEBUG
			lTotal += numL
			wTotal += numW
			mTotal += numM
			//fmt.Println(arg + " is a file")
			numFilesAndStdInputs++
		}

		if isFile(arg) {
			//fmt.Println(l(arg)) //DEBUG
			numL, numW, numM = requestedOptions(lBool, wBool, mBool, arg, argCount) // return number of lines, words, characters from the file
			//fmt.Println(numL) //DEBUG
			lTotal += numL
			wTotal += numW
			mTotal += numM
			//fmt.Println(arg + " is a file")
			numFilesAndStdInputs++
		}

	}

	var totalString string
	if numFilesAndStdInputs > 1 { // if it is multiple files: make sure to toggle and display requested values for total of all files
		if lBool { // print number of lines
			totalString = totalString + strconv.Itoa(lTotal) + "  "
		}
		if wBool { // print number of lines
			totalString = totalString + strconv.Itoa(wTotal) + "  "
		}
		if mBool { // print number of lines
			totalString = totalString + strconv.Itoa(mTotal) + "  "
		}
		fmt.Println(totalString + "total")
	}
}

//function that decides whether an option is valid (true) or invalid (false)
func isOption(arg string) bool {
	var valid bool
	valid = false
	// if the first character is a hyphen, and the arg is 4 characters at most (e.g. -lwm, -mwl, etc.)
	if string(arg[0:1]) == "-" && len(arg) <= 4 {
		letter := ""
		for i := 1; i < len(arg); i++ { // check every letter after the hyphen to see if it is a valid option(s)
			if i == len(arg) {
				letter = arg[i:]
			} else {
				letter = arg[i : i+1]
			}
			//fmt.Println(letter) // DEBUG
			if letter == "l" || letter == "w" || letter == "m" { // three valid option that are accepted as input
				//fmt.Println(arg + " might be an option!") // debug
				valid = true
			} else {
				valid = false
				break
			}
		}
	}
	return valid // valid is defaulted to false (invalid), unless appropriately assigned as true (valid)
}

// function to decide whether an arg is a valid existing file: file is in the current directory or valid filepath
func isFile(arg string) bool {
	var valid bool
	valid = false
	if _, err := os.Stat(arg); err == nil { // check if argument is a valid file
		//fmt.Printf("File exists\n") //DEBUG
		valid = true
		return valid
	} else {
		//fmt.Printf("File does not exist\n") //DEBUG
	}
	return valid
}

// function takes in which options are requested for each argument (if any)
// if a valid option(s) is valid in an argument it will then perform the option for the file that is passed to the function
// function returns the values for each respective number for each respective requested option in regards to the file (these values are returned to add for the total # of lines, words, characters)
func requestedOptions(lBool bool, wBool bool, mBool bool, fileName string, argCount int) (int, int, int) {
	var numString, lString, wString, mString string

	var numL, numW, numM int
	if lBool { // if l was requested as option
		numL = l(fileName)           // get number of lines as integer value
		lString = strconv.Itoa(numL) // convert to string variable
		//lTotal += numL // increment total number of lines by respective amount
		numString = numString + lString + "  " //add num of lines to string for printing
	}
	if wBool { // same as before, but for if w (words) was requested
		numW = w(fileName) //
		wString = strconv.Itoa(numW)
		//WTotal += numW
		numString = numString + wString + "  "
	}
	if mBool { // same as before, but for if m (characters) was requested
		numM = m(fileName)
		mString = strconv.Itoa(numM)
		//MTotal += numM
		numString = numString + mString + "  "
	}
	if isFile(fileName) { // if it is a file
		fmt.Println(numString + fileName) // print respective numbers for that file
	} else { // must be standard input
		if argCount == 0 {
			fmt.Printf(numString + " \n") // no args given, print respective numbers for given standard input
		} else {
			fmt.Printf(numString + "-\n") // print respective numbers for given standard input
		}

	}

	return numL, numW, numM
}

// function for -l option: takes a fileName (string), counts the number of lines and returns that number
func l(fileName string) int { //edit type appropriately

	var lines int
	if isFile(fileName) {
		//open+read file
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Err ", err)
		}

		///scan file
		scanner := bufio.NewScanner(file)

		// count number of lines in file
		for scanner.Scan() {
			lines++
			//fmt.Println(lines) // DEBUG
		}

	} else { // if not file, it must be standard input
		// fmt.Println("This is standard input") //DEBUG

		nLine := "\n"
		lines = strings.Count(fileName, nLine)

	}

	return lines
}

// function for -w option: takes a fileName (string), counts the number of words and returns that number
func w(fileName string) int { //edit type appropriately
	var words int
	if isFile(fileName) {
		//open+read file
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Err ", err)
		}

		scanner := bufio.NewScanner(file)

		// count number of words in file
		for scanner.Scan() {

			line := scanner.Text()
			//fmt.Println(line) // DEBUG
			splitLines := strings.Split(line, " ")
			words += len(splitLines)
			//fmt.Println(words) // DEBUG
		}
	} else { // if not file, it must be standard input
		//fmt.Println("This is standard input") //DEBUG

		words = len(strings.Fields(fileName))
	}
	return words
}

// function for -m option: takes a fileName (string), counts the number of characters and returns that number
func m(fileName string) int {
	var characters int
	if isFile(fileName) {
		//open+read file
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Err ", err)
		}

		scanner := bufio.NewScanner(file)

		// count number of characters in file
		for scanner.Scan() {

			line := scanner.Text()
			characters += len(line)
			//fmt.Println(characters) // DEBUG
		}
	} else { // if not file, it must be standard input
		//fmt.Println("This is standard input") //DEBUG
		nLine := "\n"
		lines := strings.Count(fileName, nLine)
		characters = len(fileName) - (lines * 2) //characters excluding '/n' from string

	}
	return characters

}
