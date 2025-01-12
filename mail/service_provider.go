package mail

import (
	"github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	app := Application{}
	facades.Mail = app.Init()
}

func (route *ServiceProvider) Boot() {
	facades.Queue.Register([]queue.Job{
		&SendMailJob{},
	})
}
