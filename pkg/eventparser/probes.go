// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package eventparser

import (
	"fmt"

	"github.com/intelops/tarian-detector/pkg/err"
)

var probesErr = err.New("eventparser.probes")

// Param represents a parameter with its name, type, Linux type, and processing function
type Param struct {
	name      string                    // name of the parameter
	paramType TarianParamType           // type of the parameter
	linuxType string                    // Linux type of the parameter
	function  func(any) (string, error) // function to process the parameter
}

// TarianEvent represents an event with a name, syscall ID, event size, and parameters.
type TarianEvent struct {
	name      string  // name of the event
	syscallId int     // syscall ID
	eventSize uint32  // size of the event
	params    []Param // parameters of the event
}

// TarianEventMap is a map type that maps TarianEventsE to TarianEvent
type TarianEventMap map[TarianEventsE]TarianEvent

// AddTarianEvent adds a TarianEvent to the TarianEventMap at the specified index.
func (te TarianEventMap) AddTarianEvent(idx TarianEventsE, event TarianEvent) {
	te[idx] = event
}

// NewTarianEvent creates a new TarianEvent with the given id, name, size, and params.
func NewTarianEvent(id int, name string, size uint32, params ...Param) TarianEvent {
	return TarianEvent{
		name:      name,
		syscallId: id,
		eventSize: size,
		params:    params,
	}
}

// LoadTarianEvents loads the Tarian events into 'Events' variable by generating them using GenerateTarianEvents function
func LoadTarianEvents() {
	Events = GenerateTarianEvents()
}

// GenerateTarianEvents creates and returns a TarianEventMap
func GenerateTarianEvents() TarianEventMap {
	events := make(TarianEventMap)

	execve_e := NewTarianEvent(59, "sys_execve_entry", 16897,
		Param{name: "filename", paramType: TDT_STR, linuxType: "const char *"},
		Param{name: "argv", paramType: TDT_STR_ARR, linuxType: "const char **"},
		Param{name: "envp", paramType: TDT_STR_ARR, linuxType: "const char **"},
	)
	events.AddTarianEvent(TDE_SYSCALL_EXECVE_E, execve_e)

	execve_r := NewTarianEvent(59, "sys_execve_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_EXECVE_R, execve_r)

	execveat_e := NewTarianEvent(322, "sys_execveat_entry", 16905,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int", function: parseExecveatDird},
		Param{name: "filename", paramType: TDT_STR, linuxType: "const char *"},
		Param{name: "argv", paramType: TDT_STR_ARR, linuxType: "char const **"},
		Param{name: "envp", paramType: TDT_STR_ARR, linuxType: "char const **"},
		Param{name: "flags", paramType: TDT_S32, linuxType: "int", function: parseExecveatFlags},
	)
	events.AddTarianEvent(TDE_SYSCALL_EXECVEAT_E, execveat_e)

	execveat_r := NewTarianEvent(322, "sys_execveat_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_EXECVEAT_R, execveat_r)

	clone_e := NewTarianEvent(56, "sys_clone_entry", 8733,
		Param{name: "clone_flags", paramType: TDT_U64, linuxType: "unsigned long", function: parseCloneFlags},
		Param{name: "newsp", paramType: TDT_S64, linuxType: "unsigned long"},
		Param{name: "parent_tid", paramType: TDT_S32, linuxType: "int *"},
		Param{name: "child_tid", paramType: TDT_S32, linuxType: "int *"},
		Param{name: "tls", paramType: TDT_S64, linuxType: "unsigned long"},
	)
	events.AddTarianEvent(TDE_SYSCALL_CLONE_E, clone_e)

	clone_r := NewTarianEvent(56, "sys_clone_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_CLONE_R, clone_r)

	close_e := NewTarianEvent(3, "sys_close_entry", 8705,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_CLOSE_E, close_e)

	close_r := NewTarianEvent(3, "sys_close_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_CLOSE_R, close_r)

	read_e := NewTarianEvent(0, "sys_read_entry", 12807,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "buf", paramType: TDT_BYTE_ARR, linuxType: "char *"},
		Param{name: "count", paramType: TDT_U32, linuxType: "size_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_READ_E, read_e)

	read_r := NewTarianEvent(0, "sys_read_exit", 8709,
		Param{name: "return", paramType: TDT_S64, linuxType: "ssize_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_READ_R, read_r)

	write_e := NewTarianEvent(1, "sys_write_entry", 12807,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "buf", paramType: TDT_BYTE_ARR, linuxType: "const char *"},
		Param{name: "count", paramType: TDT_U32, linuxType: "size_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_WRITE_E, write_e)

	write_r := NewTarianEvent(1, "sys_write_exit", 8709,
		Param{name: "return", paramType: TDT_S64, linuxType: "ssize_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_WRITE_R, write_r)

	open_e := NewTarianEvent(2, "sys_open_entry", 12807,
		Param{name: "filename", paramType: TDT_STR, linuxType: "const char *"},
		Param{name: "flags", paramType: TDT_S32, linuxType: "int", function: parseOpenFlags},
		Param{name: "mode", paramType: TDT_U32, linuxType: "umode_t", function: parseOpenMode},
	)
	events.AddTarianEvent(TDE_SYSCALL_OPEN_E, open_e)

	open_r := NewTarianEvent(2, "sys_open_exit", 8705,
		Param{name: "return", paramType: TDT_U32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_OPEN_R, open_r)

	readv_e := NewTarianEvent(19, "sys_readv_entry", 12807,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "vec", paramType: TDT_BYTE_ARR, linuxType: "const struct iovec *"},
		Param{name: "vlen", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_READV_E, readv_e)

	readv_r := NewTarianEvent(19, "sys_readv_exit", 8709,
		Param{name: "return", paramType: TDT_S64, linuxType: "ssize_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_READV_R, readv_r)

	writev_e := NewTarianEvent(20, "sys_writev_entry", 12807,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "vec", paramType: TDT_BYTE_ARR, linuxType: "const struct iovec *"},
		Param{name: "vlen", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_WRITEV_E, writev_e)

	writev_r := NewTarianEvent(20, "sys_writev_exit", 8709,
		Param{name: "return", paramType: TDT_S64, linuxType: "ssize_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_WRITEV_R, writev_r)

	openat_e := NewTarianEvent(257, "sys_openat_entry", 12811,
		Param{name: "dfd", paramType: TDT_S32, linuxType: "int", function: parseExecveatDird},
		Param{name: "filename", paramType: TDT_STR, linuxType: "const char *"},
		Param{name: "flags", paramType: TDT_S32, linuxType: "int", function: parseOpenFlags},
		Param{name: "mode", paramType: TDT_U32, linuxType: "umode_t", function: parseOpenMode},
	)
	events.AddTarianEvent(TDE_SYSCALL_OPENAT_E, openat_e)

	openat_r := NewTarianEvent(257, "sys_openat_exit", 8705,
		Param{name: "return", paramType: TDT_U32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_OPENAT_R, openat_r)

	openat2_e := NewTarianEvent(437, "sys_openat2_entry", 12831,
		Param{name: "dfd", paramType: TDT_S32, linuxType: "int", function: parseExecveatDird},
		Param{name: "filename", paramType: TDT_STR, linuxType: "const char *"},
		Param{name: "flags", paramType: TDT_S64, linuxType: "unsigned long", function: parseOpenat2Flags},
		Param{name: "mode", paramType: TDT_S64, linuxType: "unsigned long", function: parseOpenat2Mode},
		Param{name: "resolve", paramType: TDT_S64, linuxType: "unsigned long", function: parseOpenat2Resolve},
		Param{name: "usize", paramType: TDT_S32, linuxType: "size_t"},
	)
	events.AddTarianEvent(TDE_SYSCALL_OPENAT2_E, openat2_e)

	openat2_r := NewTarianEvent(437, "sys_openat2_exit", 8709,
		Param{name: "return", paramType: TDT_S64, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_OPENAT2_R, openat2_r)

	listen_e := NewTarianEvent(50, "sys_listen_entry", 8709,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "backlog", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_LISTEN_E, listen_e)

	listen_r := NewTarianEvent(50, "sys_listen_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_LISTEN_R, listen_r)

	socket_e := NewTarianEvent(41, "sys_socket_entry", 8713,
		Param{name: "family", paramType: TDT_S32, linuxType: "int", function: parseSocketFamily},
		Param{name: "type", paramType: TDT_S32, linuxType: "int", function: parseSocketType},
		Param{name: "protocol", paramType: TDT_S32, linuxType: "int", function: parseSocketProtocol},
	)
	events.AddTarianEvent(TDE_SYSCALL_SOCKET_E, socket_e)

	socket_r := NewTarianEvent(41, "sys_socket_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_SOCKET_R, socket_r)

	accept_e := NewTarianEvent(43, "sys_accept_entry", 8820,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "upeer_sockaddr", paramType: TDT_SOCKADDR, linuxType: "struct sockaddr *"},
		Param{name: "upper_addrlen", paramType: TDT_S32, linuxType: "int *"},
	)
	events.AddTarianEvent(TDE_SYSCALL_ACCEPT_E, accept_e)

	accept_r := NewTarianEvent(43, "sys_accept_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_ACCEPT_R, accept_r)

	bind_e := NewTarianEvent(49, "sys_bind_entry", 8820,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "umyaddr", paramType: TDT_SOCKADDR, linuxType: "struct sockaddr *"},
		Param{name: "addrlen", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_BIND_E, bind_e)

	bind_r := NewTarianEvent(49, "sys_bind_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_BIND_R, bind_r)

	connect_e := NewTarianEvent(42, "sys_connect_entry", 8820,
		Param{name: "fd", paramType: TDT_S32, linuxType: "int"},
		Param{name: "uservaddr", paramType: TDT_SOCKADDR, linuxType: "struct sockaddr *"},
		Param{name: "addrlen", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_CONNECT_E, connect_e)

	connect_r := NewTarianEvent(42, "sys_connect_exit", 8705,
		Param{name: "return", paramType: TDT_S32, linuxType: "int"},
	)
	events.AddTarianEvent(TDE_SYSCALL_CONNECT_R, connect_r)

	return events
}

func GetTarianEvent(idx TarianEventsE) (TarianEvent, error) {
	event, noEvent := Events[idx]
	if !noEvent {
		return event, probesErr.Throwf("missing event from 'var Events TarianEventMap' for key: %v", idx)
	}

	return event, nil
}

// processValue processes the value and returns the argument and an error, if any.
func (p *Param) processValue(val interface{}) (Arg, error) {
	arg := Arg{}

	// If a function is provided, call it with the value and handle the parsed value and error
	if p.function != nil {
		parsedValue, err := p.function(val)
		if err != nil {
			return arg, probesErr.Throwf("%v", err)
		}
		arg.Value = parsedValue
	} else {
		arg.Value = fmt.Sprintf("%v", val)
	}

	arg.Name = p.name
	arg.LinuxType = p.linuxType
	arg.TarianType = p.paramType.String()

	return arg, nil
}
