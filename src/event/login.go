package event

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type OnLogin struct {
}

func (OnLogin) EventID() uint16 {
	return OnLoginID
}

func (OnLogin) Event() {
}
