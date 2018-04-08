# Regex Compiler in Golang

## Introduction

This project is done by Jeremy Yon G00330435 for 3rd Year Software development GMIT, graph theory module.

## Description

The following program is written in the GO programming language that builds a non-deterministic finite automaton (NFA) from a regular expression, and the NFA can be used to check if any regular expression matches any given string of text. The regular expression can contain normal characters (a-z,A-Z,0-9), special characters(+*?|.) and brackets. 

The program has several functions:
1. The program can convert an infix string to postfix notation
2. The program will build a series of NFAs to create the final NFA
3. The program will be able to use a matching algorithm to check string matches

## Prerequisites

1. go1.10.* https://golang.org/dl/

## Compile program

1. Git clone this project to local machine
```
git clone https://github.com/yonjeremy/graph-theory-project/
```

2. Build the program
```
go build RegexCompiler.go
```

3. Run the executable
```
./RegexCompiler
```

## Run program

1. User is prompted to enter infix string to compile NFA, or enter -1 to quit program
2. User is then prompted to enter regex string to check whether string matches nfa. This will keep looping til the user exits with -1.
3. User can then enter in another infix string to generate another NFA

## Testing the program
a.b|c* --->  ccccc (any number of c accepted)
a.b.c ---> abc (a joined with b joined with c accepted) 
## Research

When I first started the project, I researching the approaches that would be used to code the project. I started doing my own research on the thompsons algorithm that is used to create the nfa, as well as the shunting yard algorithm that is used to convert the infix string to postfix. At first, I started coding the project in an object-oriented approach, but then I realised that coding an object-oriented program in golang would not be the best option.

Hence, I used video resources posted on https://web.microsoftstream.com/video/96e6f4cc-b390-4531-ba7f-84ad6ab01f47 by Dr Ian McLoughlin, as well as a paper written by Russ Cox https://swtch.com/~rsc/regexp/regexp1.html. This provided in depth research to how the code was to be coded and the different approach used to generate an NFA. 

I also used the regexp package on https://golang.org/pkg/regexp/ to understand more on the concept of regular expressions.

## How the program works

1. Shunting Algorithm
The first part of the project is converting an infix string to a postfix string. The function intopost() converts an infix regex string into postfix. This is done with the usage of stacks and popping elements off the stack. The shunting algorithm uses the precedence of the special characters to ensure the characters with higher precedence are pushed to the front of the array.

2. Thompsons algorithm
The second part of the project is using the thompsons algorithm to generate an nfa. The function regtonfa() takes in an infixstring and converts it to postfix, then the thompson's algorithm takes over. It reads the non special characters on the strings and converts them into small individual NFAs. When it reads a special character, for example a concat character ".", it joins the two small NFAs into an overall NFAs. It does this until all the runes have been read and one overall NFA has been generated.

3. Matching REGEX strings
The third part of the project is checking if the regex strings are a match with the NFA that has been generated. It does this by first calling the addState() function, which adds in all the missing states, for example an initial state with empty string. Then, it reads the regex string, and if the character matches the accept state of the final state, then the string is matched.

##Resources

https://swtch.com/~rsc/regexp/regexp1.html
https://web.microsoftstream.com/video/96e6f4cc-b390-4531-ba7f-84ad6ab01f47
https://web.microsoftstream.com/video/d08f6a02-23ec-4fa1-a781-585f1fd8c69e
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b

## Credits

Dr Ian McLoughlin
Russ Cox
