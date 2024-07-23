package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// This function creates the stack of deck for the game.
func newDeck() []int {
	cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
	var deck []int

	for i := 0; i < 24; i++ {
		deck = append(deck, cards...)
	}

	return deck
}

// Shuffle the deck of cards.
// "Fisher–Yates shuffle".
func shuffleDeck(deck []int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

// Draw a first card and remove it from list.
func drawCard(deck *[]int) int {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}

// give's card in order: player, dealer, player, dealer.
func initializeGame(deck []int) ([]int, []int, []int) {
	playerHand := []int{drawCard(&deck)}
	dealerHand := []int{drawCard(&deck)}
	playerHand = append(playerHand, drawCard(&deck))
	dealerHand = append(dealerHand, drawCard(&deck))
	return deck, playerHand, dealerHand
}

// checks the % of deck.
func deckUsedUp(deck []int, reshufflePercentage int) bool {
	return len(deck) <= len(newDeck())*(100-reshufflePercentage)/100
}

// Reshuffle the new deck
func reinitializeDeckIfNeeded(deck []int, reshufflePercentage int) []int {
	if deckUsedUp(deck, reshufflePercentage) {
		fmt.Println("Deck is", reshufflePercentage, "% used. Reshuffling")
		return shuffleDeck(newDeck())
	}
	return deck
}

func hit(deck *[]int, hand *[]int) {
	*hand = append(*hand, drawCard(deck))
}

func stand(deck *[]int, dealerHand []int) []int {
	dealerHand = dealerLogic(deck, dealerHand)
	return dealerHand
}

func dealerLogic(deck *[]int, hand []int) []int {
	for calculateHand(hand) < 17 {
		hand = append(hand, drawCard(deck))
	}
	return hand
}

func Blackjack(hand []int) bool {
	if len(hand) != 2 {
		return false
	}
	return (hand[0] == 11 && hand[1] == 10) || (hand[0] == 10 && hand[1] == 11)
}

func bust(hand []int) bool {
	return calculateHand(hand) > 21
}

func doubleDown(deck *[]int, hand *[]int, bet *int, playerMoney *int) {
	*bet *= 2
	hit(deck, hand)
	fmt.Println("Player doubled down and received one card.")
}

func calculateHand(hand []int) int {
	total := 0
	aceCount := 0

	for _, card := range hand {
		if card == 11 {
			aceCount++
		}
		total += card
	}

	for total > 21 && aceCount > 0 {
		total -= 10
		aceCount--
	}
	return total
}

func winCondition(playerHand []int, dealerHand []int, bet int, playerMoney *int) string {
	playerTotal := calculateHand(playerHand)
	dealerTotal := calculateHand(dealerHand)

	if Blackjack(playerHand) && Blackjack(dealerHand) {
		return "Push dealer's and player's blackjack"
	} else if Blackjack(playerHand) {
		*playerMoney += int(float64(bet) * 1.5)
		return "Player wins, Blackjack"
	} else if Blackjack(dealerHand) {
		*playerMoney -= bet
		return "Dealer wins, Blackjack "
	}

	if bust(playerHand) {
		*playerMoney -= bet
		return "Dealer wins! Player bust."
	} else if bust(dealerHand) {
		*playerMoney += bet
		return "Player wins! Dealer bust."
	}

	if playerTotal > dealerTotal {
		*playerMoney += bet
		return "Player wins!"
	} else if dealerTotal > playerTotal {
		*playerMoney -= bet
		return "Dealer wins!"
	} else {
		return "It's a tie!"
	}
}

func printHand(playerHand []int, dealerHand []int, showDealerFullHand bool) {
	fmt.Println("Player Hand:", playerHand, "Total:", calculateHand(playerHand))
	if showDealerFullHand {
		fmt.Println("Dealer Hand:", dealerHand, "Total:", calculateHand(dealerHand))
	} else {
		fmt.Println("Dealer Hand: [", dealerHand[0], ", ? ]")
	}
}

func placeBet(reader *bufio.Reader, playerMoney *int) int {
	for {
		fmt.Printf("You have $%d. How much would you like to bet? ", *playerMoney)
		betStr, _ := reader.ReadString('\n')
		betStr = strings.TrimSpace(betStr)
		bet, err := strconv.Atoi(betStr)
		if err != nil || bet <= 0 || bet > *playerMoney {
			fmt.Println("Invalid bet amount.")
		} else {
			return bet
		}
	}
}

func playGame(reader *bufio.Reader, deck []int, playerMoney *int, reshufflePercentage int) []int {
	bet := placeBet(reader, playerMoney)

	deck, playerHand, dealerHand := initializeGame(deck)
	printHand(playerHand, dealerHand, false)

	if Blackjack(playerHand) || Blackjack(dealerHand) {
		result := winCondition(playerHand, dealerHand, bet, playerMoney)
		fmt.Println(result)
		return deck
	}

	for {
		fmt.Println("Do you want to (h)it, (s)tand, or (d)ouble down?")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		if action == "h" {
			hit(&deck, &playerHand)
			printHand(playerHand, dealerHand, false)

			if bust(playerHand) {
				*playerMoney -= bet
				return deck
			}
		} else if action == "s" {
			dealerHand = stand(&deck, dealerHand)
			printHand(playerHand, dealerHand, true)
			result := winCondition(playerHand, dealerHand, bet, playerMoney)
			fmt.Println(result)
			return deck
		} else if action == "d" {
			doubleDown(&deck, &playerHand, &bet, playerMoney)
			printHand(playerHand, dealerHand, false)

			if bust(playerHand) {
				*playerMoney -= bet
				return deck
			}
			dealerHand = stand(&deck, dealerHand)
			printHand(playerHand, dealerHand, true)
			result := winCondition(playerHand, dealerHand, bet, playerMoney)
			fmt.Println(result)
			return deck
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	deck := newDeck()
	deck = shuffleDeck(deck)

	fmt.Print("Enter the amount of money you have: ")
	startMoneyStr, _ := reader.ReadString('\n')
	startMoneyStr = strings.TrimSpace(startMoneyStr)
	startMoney, _ := strconv.Atoi(startMoneyStr)

	playerMoney := startMoney

	var reshufflePercentage int
	for {
		fmt.Print("Enter where the blank card will be inserted. from 55 to 80: ")
		reshufflePercentageStr, _ := reader.ReadString('\n')
		reshufflePercentageStr = strings.TrimSpace(reshufflePercentageStr)
		reshufflePercentage, _ = strconv.Atoi(reshufflePercentageStr)
		if reshufflePercentage >= 55 && reshufflePercentage <= 80 {
			break
		} else {
			fmt.Println("Invalid percentage.")
		}
	}

	for {
		if playerMoney <= 0 {
			fmt.Println("You are out of money! Game over.")
			break
		}
		fmt.Println("\nStarting a new round...")
		deck = reinitializeDeckIfNeeded(deck, reshufflePercentage)
		deck = playGame(reader, deck, &playerMoney, reshufflePercentage)
		fmt.Printf("You have $%d remaining.\n", playerMoney)

		fmt.Println("Do you want to play another round? (y/n)")
		playAgain, _ := reader.ReadString('\n')
		playAgain = strings.TrimSpace(playAgain)
		if playAgain != "y" {
			break
		}
	}
	fmt.Printf("You left the game with $%d.\n", playerMoney)
}
