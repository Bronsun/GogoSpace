package server

import "github.com/Bronsun/GogoSpace/config"

func Init() {
	r := NewRouter()
	r.Run(config.GetPort())
}
