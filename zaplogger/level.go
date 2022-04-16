package zaplogger

import (
	"go.uber.org/zap/zapcore"
)

// LogLevel уровень логирования
type LogLevel = zapcore.Level

// CountdownLevel - уровни в zap формируются от -1 (Debug) до 6 (Fatal).
// Данный метод вычисляет "обратное" значение. Используется для
// консольного способа указания уровня по счетчику, например при
// использовании ключа со счетчиком -v, -vvv, -vvvvv и т.д.
func CountdownLevel(l int) LogLevel {
	return zapcore.FatalLevel - zapcore.Level(l)
}
