package day07

import (
	"aoc/utils/aocfuncs"
	"aoc/utils/aocinput"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"sort"
	"strings"
)

const (
	fiveOfAKind  = 6
	fourOfAKind  = 5
	fullHouse    = 4
	threeOfAKind = 3
	twoPair      = 2
	onePair      = 1
	highCard     = 0
)

func handToString(handType int) string {
	switch handType {
	case fiveOfAKind:
		return "Five of a kind"
	case fourOfAKind:
		return "Four of a kind"
	case fullHouse:
		return "Full house"
	case threeOfAKind:
		return "Three of a kind"
	case twoPair:
		return "Two pair"
	case onePair:
		return "One pair"
	case highCard:
		return "High card"
	default:
		log.Fatalf("Unknown hand type: %d", handType)
	}
	return ""
}

func Part1() string {
	lines := aocinput.ReadInputAsLines(7)
	// lines := aocinput.ReadSampleAsLines(7)
	hands := aocfuncs.Map[string, hand](lines, parseHand)

	sort.Sort(byCardValue(hands))

	totalWinnings := 0
	for i, hand := range hands {
		winnings := (i + 1) * hand.bet
		totalWinnings += winnings
		log.Printf("%s %d => %s => %d", hand.cards, hand.bet, handToString(identifyCards(hand.cards)), winnings)
	}
	return fmt.Sprint(totalWinnings)
}

type hand struct {
	cards string
	bet   int
}

type byCardValue []hand

func (a byCardValue) Len() int           { return len(a) }
func (a byCardValue) Less(i, j int) bool { return compareHands(a[i], a[j]) }
func (a byCardValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func parseHand(line string) hand {
	parts := strings.Split(line, " ")
	return hand{
		cards: parts[0],
		bet:   aocparse.MustAtoi(parts[1]),
	}
}

func countTypes(hand string) map[rune]int {
	typeCounts := make(map[rune]int, 0)
	for _, r := range hand {
		if _, ok := typeCounts[r]; !ok {
			typeCounts[r] = 0
		}
		typeCounts[r] += 1
	}
	return typeCounts
}

func identifyCards(cards string) int {
	counts := countTypes(cards)
	countValues := make([]int, 0, len(counts))
	for _, v := range counts {
		countValues = append(countValues, v)
	}
	if len(countValues) == 1 {
		return fiveOfAKind
	} else if len(countValues) == 2 {
		// four of a kind or full house
		if countValues[0] == 4 || countValues[1] == 4 {
			return fourOfAKind
		} else {
			return fullHouse
		}
	} else if len(countValues) == 3 {
		// three of a kind or two pair
		if countValues[0] == 3 || countValues[1] == 3 || countValues[2] == 3 {
			return threeOfAKind
		} else {
			return twoPair
		}
	} else if len(countValues) == 4 {
		return onePair
	} else {
		return highCard
	}
}

func cardToInt(card rune) int {
	switch card {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return aocparse.MustAtoi(string(card))
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		log.Fatalf("Unknown card: %c", card)
	}
	return -1
}

func compareHands(a, b hand) bool {
	aType := identifyCards(a.cards)
	bType := identifyCards(b.cards)
	if aType < bType {
		return true
	} else if bType < aType {
		return false
	} else {
		for i, aChar := range a.cards {
			bChar := rune(b.cards[i])
			aValue := cardToInt(aChar)
			bValue := cardToInt(bChar)
			if aValue < bValue {
				return true
			} else if bValue < aValue {
				return false
			}
		}
		return true // arbitrary; they are equal
	}
}

func Part2() string {
	// lines := aocinput.ReadInputAsLines(7)
	lines := aocinput.ReadSampleAsLines(7)
	hands := aocfuncs.Map[string, hand](lines, parseHand)
	hands = aocfuncs.Map[hand, hand](hands, upgradeCards)

	sort.Sort(byCardValue(hands))

	totalWinnings := 0
	for i, hand := range hands {
		winnings := (i + 1) * hand.bet
		totalWinnings += winnings
		//log.Printf("%s %d => %s => %d", hand.cards, hand.bet, handToString(identifyCards(hand.cards)), winnings)
	}
	// 250545786 is too high
	// 250621238 is too high
	// 250682036 is too high
	// 250439423 is wrong
	return fmt.Sprint(totalWinnings)
}

func upgradeCards(h hand) hand {
	counts := countTypes(h.cards)

	// most common non-joker
	mostCommonRune := 'A'
	mostCommonCount := 0
	nJokers := 0
	for k, v := range counts {
		if k == 'J' {
			nJokers++
		} else if mostCommonCount < v || (mostCommonCount == v && cardToInt2(k) > cardToInt2(mostCommonRune)) {
			// if this card is more common, or it's just as common but valued higher
			mostCommonRune = k
			mostCommonCount = v
		}
	}
	if mostCommonCount == '1' {
		mostCommonRune = 'A'
	}
	upgradedCards := strings.ReplaceAll(h.cards, "J", string(mostCommonRune))

	if nJokers > 0 {
		log.Printf("%s => %s", h.cards, upgradedCards)
	}

	return hand{upgradedCards, h.bet}
}

func cardToInt2(card rune) int {
	switch card {
	case 'J':
		return 1
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return aocparse.MustAtoi(string(card))
	case 'T':
		return 10
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		log.Fatalf("Unknown card: %c", card)
	}
	return -1
}
