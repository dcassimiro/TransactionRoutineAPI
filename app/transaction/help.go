package transaction

const (
	COMPRA_A_VISTA   = 1
	COMPRA_PARCELADA = 2
	SAQUE            = 3
	PAGAMENTO        = 4
)

func amout(op int, am float32) float32 {
	if op != PAGAMENTO {
		am = am - (am + am)
	}
	return am
}
