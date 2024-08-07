# Rules of Blackjack

I will try to summarize all rules of Blackjack.

### Main objective

Objective is to win with dealer by being closest to value 21 without exceeding it.

### deck of cards

Several decks of cards are shuffled together. Most of times is 6 decks so there is 312 cards (1 deck = 52). Dealer designates one player to insert blank plastic cards that indicates when is time to reshuffle.

### Values:

Cards are valued:

- Cards from 2-10 are valued at their face value.

- King, Queen, and Jack are valued at 10 points each.

- Aces are valued at 1 or 11 points, depending on what is more beneficial for the holder of the card.

## Key Terms:

- Soft Hand: Any hand that contains an Ace, since the Ace can be counted as either 1 or 11.

- Hard Hand: Any hand that does not contain an Ace.

- Blackjack: A player has blackjack if their initial two cards total 21 (an Ace and a 10-value card). Blackjack pays 3:2.

- Dealer's Play: The dealer must draw cards until reaching a soft 17

- Insurance: A side bet that a player can make if the dealer's upcard is an Ace. The insurance bet is half of the player's original bet. Insurance pays 2:1 if the dealer has a blackjack (Note: Not always implemented).

- Bust: When a player or dealer's hand exceeds the value of 21.

- Push/Tie: When the player and dealer have hands of equal value. The player's bet is returned.

## Game

Dealing Cards: The dealer deals cards to each player starting from the left, giving one card to each player and then one to themselves. The dealer repeats this process, giving another card to each player and one more to themselves. The dealer's second card is dealt face down.

#### Player actions:

- Hit - draw another card.

- Stand - Keep the current hand and end their turn.

- Split - If the first two cards have the same value, the player can split them into two separate hands.

- Double down - The player can double their bet and receive only one more card.

About spliting and doubling down. Casinos most of the times has a rules of maximum times of splitting the cards. This means if you have 2 same cards and you want to split them and you draw another card from deck that is the same it depends on casino rules if you can split them again or not. That rules also are aplied to double down diffrent casinos has diffrent rules and some of them forbids to double dwon after split.

If all players choose to stand dealer reveals his/her 2nd card. Next if value of cards are lower than 17 dealer draws cards to the moment that values of them is 17 or higher. In this moment dealer stops and it has few outcomes:

- Dealer Busts: The dealer's hand exceeds 21.
- Dealer Wins: The dealer's total is higher than the player's without busting.
- Push/Tie: The dealer and player have the same total
- Player Wins: The player's total is higher than the dealer's without busting.
- Player Busts: The player's hand exceeds 21.

#### Important Note

If both the player and dealer bust, the house wins and all bets go to the dealer.

## Now about code

Objectives:

- Betting - implemented
- Blank card - implemented
- Hit - implemented
- Stand - implemented
- Split - Not implemented (will be added later)
- Double down - implemented
- Insurance - Not implemented (will be added later)
- game logic - implemented

In the end, I hope that some of you will use this code for educational purposes. I don't encourage anyone to gamble money. Don't do it.
