package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "ace":
    	return 11
    case "two":
    	return 2
    case "three":
    	return 3
    case "four":
    	return 4
    case "five":
    	return 5
    case "six":
    	return 6
    case "seven":
    	return 7
    case "eight":
    	return 8
    case "nine":
    	return 9
    case "ten", "jack", "queen", "king":
    	return 10
    default:
    	return 0
    }
}

func isFaceCard(card string) bool {
    return card == "jack" || card == "queen" || card == "king"
}

func isAce(card string) bool {
    return card == "ace"
}

func isBlackjack(card1 string, card2 string) bool {
    return (ParseCard(card1) + ParseCard(card2)) == 21
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var card1Value = ParseCard(card1)
    var card2Value = ParseCard(card2)
    var hand = card1Value + card2Value
    
    const STAND = "S"
    const HIT = "H"
    const SPLIT = "P"
    const WIN = "W"
    
	if isAce(card1) && isAce(card2) {
        return SPLIT
    }
    
    if isBlackjack(card1, card2) &&
        !isAce(dealerCard) &&
        !isFaceCard(dealerCard) &&
        dealerCard != "ten" {
        return WIN
    }

    if hand >= 17 && hand <= 20 {
        return STAND
    }
    
    if hand >= 12 && hand <= 16 {
        if ParseCard(dealerCard) >= 7 {
            return HIT
        } else {
            return STAND
        }
    }
    
    if hand <= 11 {
        return HIT
    }

    return STAND
}
