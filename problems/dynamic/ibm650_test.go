package dynamic

import "testing"

/*
	Story problem is:
	Given a collection of punchcards with a value, starttime, and endtime,
	maximize the value of all punchcards run.

	Subproblem is:
	The maximum value schedule from punchcards i thru n, such that punchcards
	are sorted by time.

	Let's do the breakdown game:
	Solve max by solving all 0 to n cards.
	Solve 0,n-1 first, then 0,n-2.
	Note that we never ever want to recompute the solution to the subproblem.

	To get to i+1, what do we need to know at i?

	We can either pick or skip a punch card.
*/

type punchcard struct {
	val   int
	start int
	end   int
}

func Ibm650MaxValue(t *testing.T) {

}
