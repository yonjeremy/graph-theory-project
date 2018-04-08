// Regex NFA Compiler and Regex String Matcher
// Project by Jeremy Yon of Galway Mayo Institute of Technology
// Language used: Golang
// References: https://swtch.com/~rsc/regexp/regexp1.html
// 				https://web.microsoftstream.com/video/96e6f4cc-b390-4531-ba7f-84ad6ab01f47
//				https://web.microsoftstream.com/video/d08f6a02-23ec-4fa1-a781-585f1fd8c69e
//				https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
//				https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c
//				https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//				https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b



package main

import "fmt"

// struct for a state or also known as node
type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}

// struct for an nfa, consist of the initial state, and the accept state
type nfa struct{
	initial *state
	accept *state
}


// intopost function
// converts an infix string into a postfix string
// breaks string down into individual runes, and pushes it into a "stack" according to precedence
// default characters are "a-z,A-Z"
// special characters and brackets are given precedence, () being the highest, and then +, ?, *, ., | in that order  
// intopost function accepts an infix string and returns a postfix string
func intopost(infix string) string{

	// special characters that determine the precedence
	specials := map[rune]int{'*':10,'.':9,'|':8,'+':12,'?':11}

	// instantiate postfix slice of runes, and stack s which is also a slice of runes
	pofix,s := []rune{},[]rune{}

	// loop through every character(rune) in infix string
	for _,r := range infix {
		switch {
		// brackets are appended to the stack based on the index that they are at
		case r == '(':
			s = append(s,r)
		// push closing brackets to end of stack
		case r == ')':
			for s[len(s)-1] != '('{
				pofix = append(pofix,s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]

		// loop through the special character map, push them to stack based on precedence
		case specials[r] > 0:
			for (len(s) > 0 && specials[r] <= specials[s[len(s)-1]]){
				pofix = append(pofix,s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s,r)

		// default are normal characters, push them to the stack as per usual
		default:
			pofix = append(pofix,r)
		} 
	}

	for len(s) > 0{
		pofix = append(pofix,s[len(s)-1])
		s = s[:len(s)-1]
	}
	

	// returns a postfix string
	return string(pofix)
}


// post regular expression to non-deterministic finite automata function
// takes in a postfix string and converts it to an nfa struct
// basically, default characters and put onto the stack first as NFAs, followed by special characters
// special characters are tasked to "join" the NFAs together from the stack based on their characteristics
// at the end, there should only be one nfa left on the stack, which will be the output of the function
func regtonfa(infix string) *nfa{

	// instantiate the nfa stack that holds pointers of the nfas
	nfastack := []*nfa{}

	// calls the intopost function, converts a infix string to postfix
	pofix := intopost(infix)

	// loops through the pofix string
	for _,r := range pofix {
		switch r {
		// . indicates concatenation
		case '.':
			// pop second item on stack
			frag2 := nfastack[len(nfastack)-1]
			// remove it from stack
			nfastack = nfastack[:len(nfastack)-1]
			// pop first item on stack
			frag1 := nfastack[len(nfastack)-1]
			// remove from stack
			nfastack = nfastack[:len(nfastack)-1]
			// join two nfas 
			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})


		// | indicates alternation
		case '|':
			// pop second item on stack
			frag2 := nfastack[len(nfastack)-1]
			// remove it from stack
			nfastack = nfastack[:len(nfastack)-1]
			// pop first item on stack
			frag1 := nfastack[len(nfastack)-1]
			// remove item from stack
			nfastack = nfastack[:len(nfastack)-1]

			// new accept and initial state
			accept := state{}
			// both nfas have same initial but different accept
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		// * indicates at least 0 or more
		case '*':
			// pop one item
			frag := nfastack[len(nfastack)-1]
			// remove nfa from stack
			nfastack = nfastack[:len(nfastack)-1]
			// accept state and new initial state
			accept := state{}
			// one edge loops back to itself and another to the accept, and another to the nfa before this
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			
		// + inicates at least one 
		case '+':
			//pop one item
			frag :=  nfastack[len(nfastack) -1]
			//remove popped item from stack
			nfastack =  nfastack[:len(nfastack)-1]
			//accept state & new initial state
			accept := state{}
			//one edge going back to itself & another to the accept
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
		// The ? operator indicates zero or one
		case '?':
			//pop one item
			frag :=  nfastack[len(nfastack) -1]
			//remove popped item from stack
			nfastack =  nfastack[:len(nfastack)-1]
			//state pointing to popped item and accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			// add the nfa to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})

		// for non special characters
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	// return pointer to nfa
	return nfastack[0]
}

// adds missing states of the nfa
// takes list of pointers to state, the initial state and the accept state
func addState(l []*state, s *state, a *state) []*state{
	l = append(l,s)

	// if the s edge has edges coming from it
	if s!= a && s.symbol == 0{
		l = addState(l, s.edge1, a)
		if s.edge2 != nil{
			l = addState(l, s.edge2, a)
		}
	}

	return l
}


// function takes in an infix regex string and the regex string to match, and outputs a boolean to indicate whether there's a match
func pomatch( in string, s string) bool{
	ismatch := false

	// get the postfix nfa
	ponfa := regtonfa(in)

	// instantiate current states variable
	current := []*state{}
	// instantiate next state which moves current states to next state
	next := []*state{}

	// adds all the missing states
	current = addState(current[:],ponfa.initial, ponfa.accept)


	// loop through s , a character at a time
	for _,r := range s{
		// loop through current array
		for _, c := range current{
			//
			if c.symbol == r{
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// finally, swap current with next, and blank out next
		current, next = next, []*state{}
	}

	// loop through current array
	for _, c := range current{
		// if the state is equal to the accept state of the nfa, then match
		if c == ponfa.accept{
			ismatch = true
			break
		}
	}


	return ismatch
}


// main function
func main(){

	var input string

    // get the user to input nfa
	fmt.Println("Please enter infix string to compile NFA or -1 to quit:")
	fmt.Scanf("%s\n", &input)

	// loop through menu
	for input != "-1"{ 
		
		// get regex string to match
		fmt.Println("Please enter string to match NFA or -1 to compile new NFA")
		var regexMatch string
		fmt.Scanf("%s\n",&regexMatch)
	
		for (regexMatch != "-1"){


			fmt.Println(pomatch(input,regexMatch))
			fmt.Println("Please enter string to match NFA or -1 to compile new NFA")
			fmt.Scanf("%s\n",&regexMatch)
		}
		fmt.Println("Please enter infix to compile NFA or -1 to quit:")
		fmt.Scanf("%s\n", &input)
	}
	

}