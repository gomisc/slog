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

func (m *zapMessage) addField(f fields.Field) {
	f.Value().Extract(f.Key(), m)
}

func (l *loggerImpl) getMessagef(format string, args ...interface{}) *zapMessage {
	var (
		arguments []interface{}
		msg     = &zapMessage{}
	)

	for n := range args {
		if field, ok := args[n].(fields.Field); ok {
			msg.addField(field)
		} else {
			arguments = append(arguments, args[n])
		}
	}

	msg.message = fmt.Sprintf(format, arguments...)

	return msg
}

func (l *loggerImpl) getMessage(args ...interface{}) *zapMessage {
	var (
		message []string
		msg     = &zapMessage{}
	)

	for n := range args {
		if field, ok := args[n].(fields.Field); ok {
			msg.addField(field)
		} else {
			message = append(message, fmt.Sprint(args[n]))
		}
	}

	msg.message = strings.Join(message, " ")

	return msg
}
