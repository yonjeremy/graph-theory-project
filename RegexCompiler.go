package main

import "fmt"
// import "strconv"

// This `person` struct type has `name` and `age` fields.
type node struct {
	isAcceptState bool
	nextState map[string][]string
	id int
}

type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct{
	initial *state
	accept *state
}

func intopost(infix string) string{

	specials := map[rune]int{'*':10,'.':9,'|':8,'+':12,'?':11}

	pofix,s := []rune{},[]rune{}

	for _,r := range infix {
		switch {
		case r == '(':
			s = append(s,r)

		case r == ')':
			for s[len(s)-1] != '('{
				pofix = append(pofix,s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]

		case specials[r] > 0:
			for (len(s) > 0 && specials[r] <= specials[s[len(s)-1]]){
				pofix = append(pofix,s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s,r)

		default:
			pofix = append(pofix,r)
		} 
	}

	for len(s) > 0{
		pofix = append(pofix,s[len(s)-1])
		s = s[:len(s)-1]
	}
	

	return string(pofix)
}

func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}

	pofix = intopost(pofix)
	for _,r := range pofix {
		switch r {
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})


		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			
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

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfastack[0]
}

func addState(l []*state, s *state, a *state) []*state{
	l = append(l,s)

	if s!= a && s.symbol == 0{
		l = addState(l, s.edge1, a)
		if s.edge2 != nil{
			l = addState(l, s.edge2, a)
		}
	}

	return l
}



func pomatch( po string, s string) bool{
	ismatch := false

	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:],ponfa.initial, ponfa.accept)


	for _,r := range s{
		for _, c := range current{
			if c.symbol == r{
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}

		current, next = next, []*state{}
	}

	for _, c := range current{
		if c == ponfa.accept{
			ismatch = true
			break
		}
	}


	return ismatch
}


func main(){

	var input string

    
	fmt.Println("Please enter infix to compile NFA or -1 to quit:")
	fmt.Scanf("%s\n", &input)

	for input != "-1"{ 
		

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