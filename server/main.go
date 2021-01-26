package main

import (
	_ "ShockChatServer/conf"
	"ShockChatServer/hook"
	"ShockChatServer/router"
	"github.com/aceld/zinx/znet"
)

func main() {
	server := znet.NewServer()
	server.AddRouter(0x100, &router.LegalCheckSaltRouter{})
	server.AddRouter(0x102, &router.SaltCheck{})
	server.AddRouter(0x200, &router.RegisterRouter{})
	server.SetOnConnStop(hook.AfterConnectionStopped)
	server.SetOnConnStart(hook.ConnStart)
	server.Serve()
}
