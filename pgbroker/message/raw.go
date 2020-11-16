package message

import "io"

type Raw struct{
	Type byte
	Data []byte
}

func (m *Raw) Reader() io.Reader {
	b := NewBaseFromBytes(m.Data)
	return b.SetType(m.Type).Reader()
}

func ReadRaw(msg_type byte, raw []byte) *Raw {
	msg := &Raw{
		Type: msg_type,
		Data: raw,
	}
	return msg
}
