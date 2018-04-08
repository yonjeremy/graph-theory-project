# Regex Compiler in Golang

## Introduction

This project is done by Jeremy Yon G00330435 for 3rd Year Software development GMIT, graph theory module.

## Description

The following program is written in the GO programming language that builds a non-deterministic finite automaton (NFA) from a regular expression, and the NFA can be used to check if any regular expression matches any given string of text. The regular expression can contain normal characters (a-z,A-Z,0-9), special characters(+*?|.) and brackets. 

The program has several functions:
1. The program can convert an infix string to postfix notation
2. The program will build a series of NFAs to create the final NFA
3. The program will be able to use a matching algorithm to check string matches