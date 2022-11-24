package balances

type OperationType string

const (
	OperationSale    OperationType = "Venda"
	OperationPayment OperationType = "Pagamentos"
)

func (op OperationType) String() string {
	return string(op)
}
