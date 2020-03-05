package v1

import fmt "fmt"

func (e *EOSAccountBalance) Key() string {
	return fmt.Sprintf("%s:%s:%s", e.TokenContract, e.Symbol, e.Account)
}

func (e *EOSToken) Key() string {
	return fmt.Sprintf("%s:%s", e.Contract, e.Symbol)
}

func (c *TokenCursor) Key() string {
	return fmt.Sprintf("%s:%s", c.Contract, c.Symbol)
}

func (c *AccountBalanceCursor) Key() string {
	return fmt.Sprintf("%s:%s:%s", c.Contract, c.Symbol, c.Account)

}
