package dynamic

import (
	"fmt"
	"testing"

	"github.com/sgago/algomon/col/queue"
)

/*
	Coin Game
	https://algo.monster/problems/coin_game

	You are trying to get the highest score possible.
	The enemy player is trying to do the same, and always
	plays perfectly. However, you have an advantage and get to go first.

	Each player, you and the enemy, can pick coins from either end
	of an array. Then, it's the other players turn.

	Given coins 4, 4, 9, 4, play the game.

	Turns are:
	You: Pick 4 from the left. We want that 9.
	Enemy: Pick 4 from either side, doesn't matter. Say the pick 4 from right.
	You: Pick the 9 on the right.
	Enemy: Picks the only remaining 4.
	Yay, we win, output is 13.

	We are trying to get the highest score possible.
	Selecting a number from either side removes a number
	from either side too, until the game is done.

	We need to develop a recurrence relation, the key to all these problems.
	That's the hard part in all this. Once that's unlocked, we can
	proceed.

	So, the smallest thing is picking an element from either side.
	"Picking" is interesting here. We have to choose. And, from the guide obviously,
	we probably want to store a memo[left][right] = choice.

	Given 1, 2, 3, 4, we want to maximize our score and minimize the opponents.

	Us: 1 vs 4, we choose 4. How'd we do this? max(1, 4).
	Them: 1 vs 3, they choose 3. max(1, 3)
	Us: 1 vs 2, we choose 2. max(1, 2)
	Them: 1, they choose 1. Output 6.

	Scramble the numbers a bit, want something more interesting:
	3, 4, 1, 0, 2

	I think, in this example, want the 2.
	Us: 2
	Them: 0
	Us: 3
	Them: 4
	Us: 1

	So we want to maximize our pick and reduce their pick as much as possible.
	This'll be like:
	ourPick := max(l, r) // We want the highest
	ourPickedSide := l or r // The side we picked
	theirPick := max(min(l, r), other(l+1 or r-1)) // What we leave them with
	theirPickedSide := max(l+1 or r-1, l or r)

	ourPick(l, r) - theirPick(l, r) -> and then do it again

	So, we can memoize the results of each pick... and try to max it out

	Unsure how to represent this in perfect math terms...
*/

func TestCoinGameDfs(t *testing.T) {
	coins := []int{4, 4, 9, 4, 4}

	result := CoinGameDfs(coins)
	fmt.Println("score:", result)
}

type stateDfs struct {
	total    int
	leftIdx  int
	rightIdx int
	turn     bool
}

// This is like a 2^N solution. We need a different approach to get to N^2.
func CoinGameDfs(coins []int) int {
	q := queue.New[stateDfs](4, stateDfs{
		total:    0,
		leftIdx:  0,
		rightIdx: len(coins) - 1,
		turn:     true,
	})

	score := 0

	for !q.Empty() {
		curr := q.DeqHead()

		fmt.Println("state: ", curr)

		score = max(score, curr.total)

		if curr.turn {
			// Our turn!
			if curr.rightIdx-curr.leftIdx > 0 {
				q.EnqHead(
					stateDfs{
						total:    curr.total + coins[curr.leftIdx],
						leftIdx:  curr.leftIdx + 1,
						rightIdx: curr.rightIdx,
						turn:     !curr.turn,
					},
					stateDfs{
						total:    curr.total + coins[curr.rightIdx],
						leftIdx:  curr.leftIdx,
						rightIdx: curr.rightIdx - 1,
						turn:     !curr.turn,
					})
			} else {
				// One option remaining, so we always take that one
				score = max(score, curr.total+coins[curr.leftIdx])
			}
		} else {
			// Opponents turn!
			// They *always* take the higher number.
			if curr.rightIdx-curr.leftIdx > 0 {
				// Enemy does not want to see *both*, they pick the optimal, max choice
				if coins[curr.leftIdx] >= coins[curr.rightIdx] {
					q.EnqHead(stateDfs{
						total:    curr.total,
						leftIdx:  curr.leftIdx + 1,
						rightIdx: curr.rightIdx,
						turn:     !curr.turn,
					})
				} else {
					q.EnqHead(stateDfs{
						total:    curr.total,
						leftIdx:  curr.leftIdx,
						rightIdx: curr.rightIdx - 1,
						turn:     !curr.turn,
					})
				}

			}

			// With one option remaining, enemy always take that one
		}
	}

	return score
}

func TestCoinGameDp(t *testing.T) {

}

func CoinGameDp(t *testing.T) {
	// 1. Define problem. We're playing coin game -> pick coins maxing our score and mining opponents
	// 2. Identify the substates. The set of params that change as the problem progresses.
	//    For the coin game, this is left index, right index, score, and turn. We could also sum the total and subtract for that to determine remaining.
	// 3. Define the base case. Mmk, so:
	//    Pick -> max(coins[l], coins[r])
	//    Opponent ->
}
