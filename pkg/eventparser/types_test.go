// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package eventparser

import (
	"testing"
)

// TestTarianMetaData_Event tests the Event function.
func TestTarianMetaData_Event(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   int32
	}{
		{
			name:   "default values",
			fields: TarianMetaData{},
			want:   0,
		},
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Event = 29
				return t
			}(),
			want: 29,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Event(); got != tt.want {
				t.Errorf("TarianMetaData.Event() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Nparams tests the Nparams function.
func TestTarianMetaData_Nparams(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint8
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Nparams = 4
				return t
			}(),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Nparams(); got != tt.want {
				t.Errorf("TarianMetaData.Nparams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Syscall tests the Syscall function.
func TestTarianMetaData_Syscall(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   int32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Syscall = 29
				return t
			}(),
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Syscall(); got != tt.want {
				t.Errorf("TarianMetaData.Syscall() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Ts tests the Ts function.
func TestTarianMetaData_Ts(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Ts = 2932984232342
				return t
			}(),
			want: 2932984232342,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Ts(); got != tt.want {
				t.Errorf("TarianMetaData.Ts() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Processor tests the Processor function.
func TestTarianMetaData_Processor(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint16
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Processor = 15
				return t
			}(),
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Processor(); got != tt.want {
				t.Errorf("TarianMetaData.Processor() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_StartTime tests the StartTime function.
func TestTarianMetaData_StartTime(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.StartTime = 7423975270240932424
				return t
			}(),
			want: 7423975270240932424,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.StartTime(); got != tt.want {
				t.Errorf("TarianMetaData.StartTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_HostPid tests the HostPid function.
func TestTarianMetaData_HostPid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.HostPid = 344532
				return t
			}(),
			want: 344532,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.HostPid(); got != tt.want {
				t.Errorf("TarianMetaData.HostPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_HostTgid tests the HostTgid function.
func TestTarianMetaData_HostTgid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.HostTgid = 32342
				return t
			}(),
			want: 32342,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.HostTgid(); got != tt.want {
				t.Errorf("TarianMetaData.HostTgid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_HostPpid tests the HostPpid function.
func TestTarianMetaData_HostPpid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.HostPpid = 54354345
				return t
			}(),
			want: 54354345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.HostPpid(); got != tt.want {
				t.Errorf("TarianMetaData.HostPpid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Pid tests the Pid function.
func TestTarianMetaData_Pid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.Pid = 232543534
				return t
			}(),
			want: 232543534,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Pid(); got != tt.want {
				t.Errorf("TarianMetaData.Pid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Tgid tests the Tgid function.
func TestTarianMetaData_Tgid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.Tgid = 345252
				return t
			}(),
			want: 345252,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Tgid(); got != tt.want {
				t.Errorf("TarianMetaData.Tgid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Ppid tests the Ppid function.
func TestTarianMetaData_Ppid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.Ppid = 257435
				return t
			}(),
			want: 257435,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Ppid(); got != tt.want {
				t.Errorf("TarianMetaData.Ppid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Uid tests the Uid function.
func TestTarianMetaData_Uid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.Uid = 12345
				return t
			}(),
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Uid(); got != tt.want {
				t.Errorf("TarianMetaData.Uid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Gid tests the Gid function.
func TestTarianMetaData_Gid(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint32
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.Gid = 9876543
				return t
			}(),
			want: 9876543,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Gid(); got != tt.want {
				t.Errorf("TarianMetaData.Gid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_CgroupId tests the CgroupId function.
func TestTarianMetaData_CgroupId(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.CgroupId = 9342795229
				return t
			}(),
			want: 9342795229,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.CgroupId(); got != tt.want {
				t.Errorf("TarianMetaData.CgroupId() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_MountNsId tests the MountNsId function.
func TestTarianMetaData_MountNsId(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.MountNsId = 432543253
				return t
			}(),
			want: 432543253,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.MountNsId(); got != tt.want {
				t.Errorf("TarianMetaData.MountNsId() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_PidNsId tests the PidNsId function.
func TestTarianMetaData_PidNsId(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.PidNsId = 7425928524
				return t
			}(),
			want: 7425928524,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.PidNsId(); got != tt.want {
				t.Errorf("TarianMetaData.PidNsId() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_ExecId tests the ExecId function.
func TestTarianMetaData_ExecId(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.ExecId = 12432543632652652632
				return t
			}(),
			want: 12432543632652652632,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.ExecId(); got != tt.want {
				t.Errorf("TarianMetaData.ExecId() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_ParentExecId tests the ParentExecId function.
func TestTarianMetaData_ParentExecId(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   uint64
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.ParentExecId = 4358723423092423121
				return t
			}(),
			want: 4358723423092423121,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.ParentExecId(); got != tt.want {
				t.Errorf("TarianMetaData.ParentExecId() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Comm tests the Comm function.
func TestTarianMetaData_Comm(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.MetaData.Task.Comm = [16]byte{99, 111, 109, 109, 0}
				return t
			}(),
			want: "comm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Comm(); got != tt.want {
				t.Errorf("TarianMetaData.Comm() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Sysname tests the Sysname function.
func TestTarianMetaData_Sysname(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.SystemInfo.Sysname = [65]byte{0, 0, 115, 121, 115, 0}
				return t
			}(),
			want: "sys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Sysname(); got != tt.want {
				t.Errorf("TarianMetaData.Sysname() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Nodename tests the Nodename function.
func TestTarianMetaData_Nodename(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.SystemInfo.Nodename = [65]byte{115, 121, 115, 0}
				return t
			}(),
			want: "sys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Nodename(); got != tt.want {
				t.Errorf("TarianMetaData.Nodename() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Release tests the Release function.
func TestTarianMetaData_Release(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.SystemInfo.Release = [65]byte{115, 121, 115}
				return t
			}(),
			want: "sys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Release(); got != tt.want {
				t.Errorf("TarianMetaData.Release() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Version tests the Version function.
func TestTarianMetaData_Version(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.SystemInfo.Version = [65]byte{115, 121, 115}
				return t
			}(),
			want: "sys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Version(); got != tt.want {
				t.Errorf("TarianMetaData.Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Machine tests the Machine function.
func TestTarianMetaData_Machine(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.SystemInfo.Machine = [65]byte{115, 121, 115}
				return t
			}(),
			want: "sys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Machine(); got != tt.want {
				t.Errorf("TarianMetaData.Machine() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTarianMetaData_Domainname tests the Domainname function.
func TestTarianMetaData_Domainname(t *testing.T) {
	tests := []struct {
		name   string
		fields TarianMetaData
		want   string
	}{
		{
			name: "valid values",
			fields: func() TarianMetaData {
				var t TarianMetaData
				t.SystemInfo.Domainname = [65]byte{115, 121, 115}
				return t
			}(),
			want: "sys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TarianMetaData{
				MetaData:   tt.fields.MetaData,
				SystemInfo: tt.fields.SystemInfo,
			}
			if got := tr.Domainname(); got != tt.want {
				t.Errorf("TarianMetaData.Domainname() = %v, want %v", got, tt.want)
			}
		})
	}
}
