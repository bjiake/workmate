package something

import (
	"github.com/google/wire"
	"sync"
	"workmate/internal/domain/interfaces"
	SomethingSvc "workmate/internal/service/something"
)

var (
	svc     *SomethingSvc.Service
	svcOnce sync.Once
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideSomethingService,

	wire.Bind(new(interfaces.SomethingService), new(*SomethingSvc.Service)),
)

func ProvideSomethingService() *SomethingSvc.Service {
	svcOnce.Do(func() {
		svc = &SomethingSvc.Service{}
	})

	return svc
}
