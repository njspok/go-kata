package money

type IExpression interface {
	Reduce(bank *Bank, to string) *Money
	Plus(IExpression) IExpression
	Times(uint) IExpression
}
