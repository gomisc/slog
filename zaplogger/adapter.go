package zaplogger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"git.corout.in/golibs/fields"
)

type zapMessage struct {
	message string
	fields  []zap.Field
}

func (m *zapMessage) addFields(fds ...fields.Field) {
	for i := 0; i < len(fds); i++ {
		fds[i].Value().Extract(fds[i].Key(), m)
	}
}

func getMessage(args ...any) *zapMessage {
	var (
		message []string
		msg     = &zapMessage{}
	)

	for n := range args {
		switch args[n].(type) {
		case fields.Field:
			if field, ok := args[n].(fields.Field); ok {
				msg.addFields(field)
			}
		case []fields.Field:
			if flds, ok := args[n].([]fields.Field); ok {
				msg.addFields(flds...)
			}
		default:
			message = append(message, fmt.Sprint(args[n]))
		}
	}

	msg.message = strings.Join(message, " ")

	return msg
}

func getMessagef(format string, args ...any) *zapMessage {
	var (
		arguments []interface{}
		msg       = &zapMessage{}
	)

	for n := range args {
		switch args[n].(type) {
		case fields.Field:
			if field, ok := args[n].(fields.Field); ok {
				msg.addFields(field)
			}
		case []fields.Field:
			if flds, ok := args[n].([]fields.Field); ok {
				msg.addFields(flds...)
			}
		default:
			arguments = append(arguments, args[n])
		}
	}

	msg.message = fmt.Sprintf(format, arguments...)

	return msg
}
