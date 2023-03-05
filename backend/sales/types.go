package sales

type PaymentType string

const (
	PixPayment       PaymentType = "Pagamento no PIX"
	CashPayment      PaymentType = "Pagamento no Dinheiro"
	CardPayment      PaymentType = "Pagamento no Cartão"
	EmployersPayment PaymentType = "Pagamento de funcionarios"
	SuppliersPayment PaymentType = "Pagamento de fornecedores"
	AnotherPayments  PaymentType = "Outros pagamentos"
)
