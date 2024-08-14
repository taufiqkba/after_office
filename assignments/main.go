package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(arraySign([]int{2, 1}))                    //1
	fmt.Println(arraySign([]int{-2, 1}))                   //-1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) //1
	fmt.Println(arraySign([]int{1, 2, 0, -1, 2}))          //0
	fmt.Println("")

	fmt.Println(isAnagram("anak", "kana"))       // true
	fmt.Println(isAnagram("anak", "mana"))       // false
	fmt.Println(isAnagram("anagram", "aagrnam")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println("")

	fmt.Println(findTheDifference("abcd", "abcde")) // 'e'
	fmt.Println(findTheDifference("abcd", "abced")) // 'e'
	fmt.Println(findTheDifference("", "y"))         // 'y'
	fmt.Println("")

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	// initialize as a positive number
	x := 1

	for _, numb := range nums {
		if numb == 0 {
			return 0
		}
		if numb < 0 {
			x *= -1
		}
	}
	// This return will be 1 if positive, -1 if negative, and return 0 if array contain number 0
	return x
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// to check length, anagram must have same length between s and t variable
	if len(s) != len(t) {
		return false
	}

	// initialize count map
	// key: rune
	// value: int
	charSum := make(map[rune]int)

	// sum character on string s
	for _, char := range s {
		charSum[char]++
	}

	// check string t, if string t has character more than s, it can't be anagram, return false
	for _, char := range t {
		charSum[char]--
		if charSum[char] < 0 {
			return false
		}
	}

	// If the return is not false, it means the string s&t has the same number of characters, it is an anagram
	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) string {
	var result byte

	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}

	for i := 0; i < len(t); i++ {
		result ^= t[i]
	}
	fmt.Print(result, " -> ")
	return string(result)
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// Sort the array using library sort
	sort.Ints(arr)

	// Check if the difference between each pair of adjacent elements is constant
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	d.cards = make([]Card, 52)
	index := 0

	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards[index] = Card{symbol: symbol, number: number}
			index++
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
	//rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
	if n > len(d.cards) {
		n = len(d.cards)
	}
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Printf("---\n\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Printf("---\n\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Printf("---\n\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
