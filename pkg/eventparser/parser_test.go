// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package eventparser

import (
	"reflect"
	"testing"
)

// TestNewByteStream tests the NewByteStream function.
func TestNewByteStream(t *testing.T) {
	type args struct {
		inputData []byte
		n         uint8
	}
	tests := []struct {
		name string
		args args
		want *ByteStream
	}{
		{
			name: "valid values",
			args: args{
				inputData: []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:         8,
			},
			want: &ByteStream{
				data:     []byte{1, 2, 3, 4, 5, 6, 7, 8},
				position: 0,
				nparams:  8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewByteStream(tt.args.inputData, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewByteStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestParseByteArray tests the ParseByteArray function.
func TestParseByteArray(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name       string
		args       args
		loadEvents bool
		want       TarianDetectorEvent
		wantErr    bool
	}{
		{
			name: "short data buffer",
			args: args{
				data: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			},
			loadEvents: false,
			want:       TarianDetectorEvent{},
			wantErr:    true,
		},
		{
			name: "invalid event id",
			args: args{
				data: func() []byte {
					data := make([]byte, 1000)
					data[0] = 1
					data[1] = 2
					data[2] = 3
					data[3] = 4

					return data
				}(),
			},
			loadEvents: false,
			want:       TarianDetectorEvent{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		if tt.loadEvents {
			LoadTarianEvents()
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseByteArray(nil, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseByteArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestByteStream_parseParams tests the parseParams function.
func TestByteStream_parseParams(t *testing.T) {
	type args struct {
		event TarianEvent
	}
	tests := []struct {
		name    string
		fields  ByteStream
		args    args
		want    []Arg
		wantErr bool
	}{
		{
			name:    "empty tarian event",
			wantErr: true,
		}, {
			name: "break bs.position >= len(bs.data)",
			fields: ByteStream{
				position: 0,
				nparams:  2,
			},
			args: args{
				event: TarianEvent{
					name:      "test",
					syscallId: 0,
					eventSize: 783,
					params: []Param{
						{}, {},
					},
				},
			},
			want:    []Arg{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &ByteStream{
				data:     tt.fields.data,
				position: tt.fields.position,
				nparams:  tt.fields.nparams,
			}
			got, err := bs.parseParams(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteStream.parseParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteStream.parseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestByteStream_parseParam tests the parseParam function.
func TestByteStream_parseParam(t *testing.T) {
	type args struct {
		p Param
	}
	tests := []struct {
		name    string
		fields  ByteStream
		args    args
		want    Arg
		wantErr bool
	}{
		{
			name: "call Uint8",
			fields: ByteStream{
				data:     []byte{1, 2, 3},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_U8,
					linuxType: "uint8_t",
				},
			},
			want: Arg{
				Name:       "test",
				Value:      "1",
				TarianType: "TDT_U8",
				LinuxType:  "uint8_t",
			},
			wantErr: false,
		},
		{
			name: "call Uint16",
			fields: ByteStream{
				data:     []byte{1, 2, 3},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_U16,
					linuxType: "uint16_t",
				},
			},
			want: Arg{
				Name:       "test",
				Value:      "513",
				TarianType: "TDT_U16",
				LinuxType:  "uint16_t",
			},
			wantErr: false,
		},
		{
			name: "call Uint64",
			fields: ByteStream{
				data:     []byte{1, 2, 3, 4, 5, 6, 7, 8},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_U64,
					linuxType: "uint64_t",
				},
			},
			want: Arg{
				Name:       "test",
				Value:      "578437695752307201",
				TarianType: "TDT_U64",
				LinuxType:  "uint64_t",
			},
			wantErr: false,
		},
		{
			name: "call Int8",
			fields: ByteStream{
				data:     []byte{1, 2, 3},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_S8,
					linuxType: "int8_t",
				},
			},
			want:    Arg{"test", "1", "TDT_S8", "int8_t"},
			wantErr: false,
		},
		{
			name: "call Int16",
			fields: ByteStream{
				data:     []byte{1, 2, 3},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_S16,
					linuxType: "int16_t",
				},
			},
			want:    Arg{"test", "513", "TDT_S16", "int16_t"},
			wantErr: false,
		},
		{
			name: "call Int64",
			fields: ByteStream{
				data:     []byte{1, 2, 3, 4, 5, 6, 7, 8},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_S64,
					linuxType: "int64_t",
				},
			},
			want:    Arg{"test", "578437695752307201", "TDT_S64", "int64_t"},
			wantErr: false,
		},
		{
			name: "call String",
			fields: ByteStream{
				data:     []byte{4, 0, 65, 66, 67, 68},
				position: 0,
				nparams:  2,
			},
			args: args{
				p: Param{
					name:      "test",
					paramType: TDT_STR,
					linuxType: "string",
				},
			},
			want:    Arg{"test", "ABCD", "TDT_STR", "string"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &ByteStream{
				data:     tt.fields.data,
				position: tt.fields.position,
				nparams:  tt.fields.nparams,
			}
			got, err := bs.parseParam(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteStream.parseParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteStream.parseParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestByteStream_parseString tests the parseString function.
func TestByteStream_parseString(t *testing.T) {
	type fields struct {
		data     []byte
		position int
		nparams  uint8
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "invalid values",
			fields: fields{
				data:     []byte{1},
				position: 0,
				nparams:  2,
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &ByteStream{
				data:     tt.fields.data,
				position: tt.fields.position,
				nparams:  tt.fields.nparams,
			}
			got, err := bs.parseString()
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteStream.parseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ByteStream.parseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestByteStream_parseRawArray tests the parseRawArray function.
func TestByteStream_parseRawArray(t *testing.T) {
	type fields struct {
		data     []byte
		position int
		nparams  uint8
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "invalid values",
			fields: fields{
				data:     []byte{1},
				position: 0,
				nparams:  2,
			},
			want:    []byte{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &ByteStream{
				data:     tt.fields.data,
				position: tt.fields.position,
				nparams:  tt.fields.nparams,
			}
			got, err := bs.parseRawArray()
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteStream.parseRawArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteStream.parseRawArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestByteStream_parseSocketAddress tests the parseSocketAddress function.
func TestByteStream_parseSocketAddress(t *testing.T) {
	type fields struct {
		data     []byte
		position int
		nparams  uint8
	}
	tests := []struct {
		name    string
		fields  fields
		want    any
		wantErr bool
	}{
		{
			name: "invalid values",
			fields: fields{
				data:     []byte{},
				position: 0,
				nparams:  2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unhandled case",
			fields: fields{
				data:     []byte{0},
				position: 0,
				nparams:  2,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "AF_INET Uint16 error",
			fields: fields{
				data:     []byte{2, 8, 9, 0, 3, 1},
				position: 0,
				nparams:  2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "AF_INET6 Uint16 error",
			fields: fields{
				data:     []byte{10, 8, 9, 0, 3, 1},
				position: 0,
				nparams:  2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "AF_UNIX Uint16 error",
			fields: fields{
				data:     []byte{1, 8},
				position: 0,
				nparams:  2,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &ByteStream{
				data:     tt.fields.data,
				position: tt.fields.position,
				nparams:  tt.fields.nparams,
			}
			got, err := bs.parseSocketAddress()
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteStream.parseSocketAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteStream.parseSocketAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
