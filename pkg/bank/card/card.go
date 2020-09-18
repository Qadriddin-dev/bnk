package card

import (
	"bank/pkg/bank/types"
)

func PaymentSource(cards []types.Card) []types.PaymentSource {
	var payment []types.PaymentSource
	for _, card := range cards {
		if card.Balance <= 0 {

			continue
		}
		if !card.Active {
			continue

		}
		payment = append(payment, types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		})

	}
	return payment
}
