package factory_iface


// This pattern allows us to abstract away the creation of specific implementations of an interface.
// The caller just needs to specify the method, and the factory returns the appropriate implementation.
func main() {
	processor := GetPaymentProcessor("paypal")
	if processor != nil {
		if err := processor.Process(100.0); err != nil {
			panic(err)
		}
	} else {
		panic("Payment processor not found")
	}
}

func GetPaymentProcessor(method string) PaymentProcessor {
	switch method {
	case "paypal":
		return PaypalProcessor{}
	case "stripe":
		return StripeProcessor{}
	default:
		return nil
	}
}

type PaymentProcessor interface {
	Process(amount float64) error
}

type PaypalProcessor struct{}

func (pp PaypalProcessor) Process(amount float64) error {
	return nil
}

type StripeProcessor struct{}

func (sp StripeProcessor) Process(amount float64) error {
	return nil
}
