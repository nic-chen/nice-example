package srv

import(
	"github.com/nic-chen/nice/micro/registry"
)

func RunAll(register registry.Registry) {
	go RunGreeter(register)
	RunApi(register)
}
