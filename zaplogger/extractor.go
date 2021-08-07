package zaplogger

import (
	"go.uber.org/zap"
)

func (m *zapMessage) Bool(key string, value bool) {
	m.fields = append(m.fields, zap.Bool(key, value))
}

func (m *zapMessage) Int(key string, value int) {
	m.fields = append(m.fields, zap.Int(key, value))
}

func (m *zapMessage) Int8(key string, value int8) {
	m.fields = append(m.fields, zap.Int8(key, value))
}

func (m *zapMessage) Int16(key string, value int16) {
	m.fields = append(m.fields, zap.Int16(key, value))
}

func (m *zapMessage) Int32(key string, value int32) {
	m.fields = append(m.fields, zap.Int32(key, value))
}

func (m *zapMessage) Int64(key string, value int64) {
	m.fields = append(m.fields, zap.Int64(key, value))
}

func (m *zapMessage) Uint(key string, value uint) {
	m.fields = append(m.fields, zap.Uint(key, value))
}

func (m *zapMessage) Uint8(key string, value uint8) {
	m.fields = append(m.fields, zap.Uint8(key, value))
}

func (m *zapMessage) Uint16(key string, value uint16) {
	m.fields = append(m.fields, zap.Uint16(key, value))
}

func (m *zapMessage) Uint32(key string, value uint32) {
	m.fields = append(m.fields, zap.Uint32(key, value))
}

func (m *zapMessage) Uint64(key string, value uint64) {
	m.fields = append(m.fields, zap.Uint64(key, value))
}

func (m *zapMessage) Float32(key string, value float32) {
	m.fields = append(m.fields, zap.Float32(key, value))
}

func (m *zapMessage) Float64(key string, value float64) {
	m.fields = append(m.fields, zap.Float64(key, value))
}

func (m *zapMessage) Str(key string, value string) {
	m.fields = append(m.fields, zap.String(key, value))
}

func (m *zapMessage) Strings(key string, values []string) {
	m.fields = append(m.fields, zap.Strings(key, values))
}

func (m *zapMessage) Any(key string, value interface{}) {
	m.fields = append(m.fields, zap.Any(key, value))
}



