package lb

import (
	"context"

	"github.com/Melon-Network-Inc/gateway-service/pkg/config"
)

type lb struct {
	accountServiceLB
	paymentServiceLB
}

func New(ctx context.Context, lbConfig config.LoadBalancerConfigProvider) (Accessor, error) {
	accountLB := newAccountLB(lbConfig.GetAccountServiceAddressesList())
	paymentLB := newPaymentLB(lbConfig.GetPaymentServiceAddressesList())
	return &lb{
		accountServiceLB: accountLB,
		paymentServiceLB: paymentLB,
	}, nil
}