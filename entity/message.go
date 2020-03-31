package entity

type RegisterMessage struct {
	Address     string `json:"address"`
	Port        string `json:"port"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

func (registerMessage *RegisterMessage) SetMessage(address, port, method, description string) {
	registerMessage.Method = method
	registerMessage.Port = port
	registerMessage.Address = address
	registerMessage.Description = description
}
