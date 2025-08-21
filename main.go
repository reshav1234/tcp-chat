package main

type sendCommand struct {
	message string
}

type nameCommand struct {
	name string
}

type messageCommand struct {
	message string
	name    string
}

type commandWriter struct {
	writer io.Writer
}

func newCommandWriter(writer io.Writer) *commandWriter {
	return &commandWriter{
		writer: writer,
	}
}

func (w *commandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}

func (w *commandWriter) Write(command interface{}) error {
	var err error
	switch v := command.(type) {
	case sendCommand:
		err = w.writeString(fmt.Sprintf("SEND %v\n", v.message))
	case nameCommand:
		err = w.writeString(fmt.Sprintf("Name %v\n", v.name))
	case messageCommand:
		err := w.writeString(fmt.Sprintf("MESSAGE %v\n", v.name, v.message))
	default:
		err = UnknownCommand
	}
	return err
}
