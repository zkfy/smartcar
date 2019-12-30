
package services

import (
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type CarctlHandler struct {
	container component.IContainer
}

func NewCarctlHandler(container component.IContainer) (u *CarctlHandler) {
	return &CarctlHandler{container: container}
}

//Handle .
func (u *CarctlHandler) Handle(ctx *context.Context) (r interface{}) {
	Forward(time.Second)
	return "success"
}

func (u *CarctlHandler) ForwardHandle(ctx *context.Context) (r interface{}) {
	Forward(time.Second)
	return "success"
}

func (u *CarctlHandler) FLeftHandle(ctx *context.Context) (r interface{}) {
	FLeft(time.Second)
	return "success"
}

func (u *CarctlHandler) FRightHandle(ctx *context.Context) (r interface{}) {
	FRight(time.Second)
	return "success"
}

func (u *CarctlHandler) BLeftHandle(ctx *context.Context) (r interface{}) {
	BLeft(time.Second)
	return "success"
}

func (u *CarctlHandler) BRightHandle(ctx *context.Context) (r interface{}) {
	BRight(time.Second)
	return "success"
}

func (u *CarctlHandler) BackHandle(ctx *context.Context) (r interface{}) {
	Back(time.Second)
	return "success"
}
