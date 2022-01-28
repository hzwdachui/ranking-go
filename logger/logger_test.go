package logger

import (
	"testing"
)

func TestMain(m *testing.M) {
	InitLogger("debug")
	m.Run()
}
func TestInitLogger(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1 正常加载",
			args: args{
				logLevel: "debug",
			},
		},
		{
			name: "case2 正常加载",
			args: args{
				logLevel: "warn",
			},
		},
		{
			name: "case3 正常加载",
			args: args{
				logLevel: "info",
			},
		},
		{
			name: "case4 正常加载",
			args: args{
				logLevel: "error",
			},
		},
		{
			name: "case4 正常加载",
			args: args{
				logLevel: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger(tt.args.logLevel)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				v: func() []interface{} {
					tmp := []string{"test"}
					v := make([]interface{}, len(tmp))
					for i, value := range tmp {
						v[i] = value
					}
					return v
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.v...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.format, tt.args.v...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.format, tt.args.v...)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				v: func() []interface{} {
					tmp := []string{"test"}
					v := make([]interface{}, len(tmp))
					for i, value := range tmp {
						v[i] = value
					}
					return v
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.v...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				v: func() []interface{} {
					tmp := []string{"test"}
					v := make([]interface{}, len(tmp))
					for i, value := range tmp {
						v[i] = value
					}
					return v
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.v...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				v: func() []interface{} {
					tmp := []string{"test"}
					v := make([]interface{}, len(tmp))
					for i, value := range tmp {
						v[i] = value
					}
					return v
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.v...)
		})
	}
}
