// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package eventparser

import (
	"fmt"

	"github.com/intelops/tarian-detector/pkg/utils"
)

// Arg represents an argument with its name, value, and types for Tarian and Linux.
type Arg struct {
	Name       string // Name is the name of the argument
	Value      string // Value is the value of the argument
	TarianType string // TarianType is the Tarian type of the argument
	LinuxType  string // LinuxType is the Linux type of the argumen
}

// String returns a string representation of the Arg struct.
func (arg Arg) String() string {
	str := ""

	str += fmt.Sprintf("Name: %v\n", arg.Name)
	str += fmt.Sprintf("Value: %v\n", arg.Value)
	str += fmt.Sprintf("TarianType: %v\n", arg.TarianType)
	str += fmt.Sprintf("LinuxType: %v\n", arg.LinuxType)

	return str
}

// HostDetails represents information about the host system.
type HostDetails struct {
	// Sysname is the name of the operating system.
	Sysname string `json:"sysname"`
	// Hostname is the name of the host system.
	Hostname string `json:"hostname"`
	// Release is the release level of the operating system.
	Release string `json:"release"`
	// KernelVersion is the version of the kernel.
	KernelVersion string `json:"kernelversion"`
	// Machine is the machine architecture.
	Machine string `json:"machine"`
	// Domainname is the domain name of the host.
	Domainname string `json:"domainname"`
}

// String returns a string representation of the HostDetails struct.
func (h HostDetails) String() string {
	str := ""

	str += fmt.Sprintf("Sysname: %v\n", h.Sysname)
	str += fmt.Sprintf("Hostname: %v\n", h.Hostname)
	str += fmt.Sprintf("Release: %v\n", h.Release)
	str += fmt.Sprintf("KernelVersion: %v\n", h.KernelVersion)
	str += fmt.Sprintf("Machine: %v\n", h.Machine)
	str += fmt.Sprintf("Domainname: %v\n", h.Domainname)

	return str
}

// Kubernetes represents information about a Kubernetes pod and container.
type Kubernetes struct {
	// pod information
	// PodUid is the unique ID for the pod.
	PodUid string `json:"pod_uid"`
	// PodName is the name of the pod.
	PodName string `json:"pod_name"`
	// PodGeneratedName is the generated name of the pod.
	PodGeneratedName string `json:"pod_generated_name"`
	// PodKind is the kind of the pod (e.g. Pod).
	PodKind string `json:"pod_kind"`
	// PodAPIVersion is the API version of the pod.
	PodAPIVersion string `json:"pod_api_version"`
	// PodLabels is a map of labels applied to the pod.
	PodLabels map[string]string `json:"pod_labels"`
	// PodAnnotations is a map of annotations applied to the pod.
	PodAnnotations map[string]string `json:"pod_annotations"`

	// ContainerID is the ID of the container.
	ContainerID string `json:"container_id"`

	// Namespace is the name of the namespace that the pod is in.
	Namespace string `json:"namespace"`
}

// String returns a string representation of the Kubernetes struct.
func (k Kubernetes) String() string {
	str := ""

	str += fmt.Sprintf("PodUid: %v\n", k.PodUid)
	str += fmt.Sprintf("PodName: %v\n", k.PodName)
	str += fmt.Sprintf("PodGeneratedName: %v\n", k.PodGeneratedName)
	str += fmt.Sprintf("PodKind: %v\n", k.PodKind)
	str += fmt.Sprintf("PodAPIVersion: %v\n", k.PodAPIVersion)
	str += fmt.Sprintf("PodLabels: %v\n", k.PodLabels)
	str += fmt.Sprintf("PodAnnotations: %v\n", k.PodAnnotations)
	str += fmt.Sprintf("ContainerID: %v\n", k.ContainerID)
	str += fmt.Sprintf("Namespace: %v\n", k.Namespace)

	return str
}

// TarianDetectorEvent represents an event that has been received from the
// kernel and parsed.
type TarianDetectorEvent struct {
	// EventId is a unique identifier represents what kind of event it is.
	EventId string `json:"event_id"`
	// SyscallId is the identifier of the system call that event is generated from.
	SyscallId int32 `json:"syscall_id"`
	// Timestamp is the timestamp of when the event started by kernel.
	Timestamp uint64 `json:"timestamp"`
	// ProcessorId is the identifier of the processor that recorded the event.
	ProcessorId uint16 `json:"processor_id"`
	// ThreadStartTime is the timestamp of when the thread that recorded the
	// event started.
	ThreadStartTime uint64 `json:"thread_start_time"`
	// HostProcessId is the ID of the process in the host's context.
	HostProcessId uint32 `json:"host_process_id"`
	// HostThreadId is the ID of the thread in the host's context.
	HostThreadId uint32 `json:"host_thread_id"`
	// HostParentProcessId is the ID of the parent process in the host's
	// context.
	HostParentProcessId uint32 `json:"host_parent_process_id"`
	// ProcessId is the ID of the process in the container's context.
	ProcessId uint32 `json:"process_id"`
	// ThreadId is the ID of the thread in the container's context.
	ThreadId uint32 `json:"thread_id"`
	// ParentProcessId is the ID of the parent process in the container's
	// context.
	ParentProcessId uint32 `json:"parent_process_id"`
	// UserId is the user ID of the user running the process.
	UserId uint32 `json:"user_id"`
	// GroupId is the group ID of the group that the user belongs to.
	GroupId uint32 `json:"group_id"`
	// CgroupId is the control group ID of the process.
	CgroupId uint64 `json:"cgroup_id"`
	// MountNamespaceId is the ID of the mount namespace of the process.
	MountNamespaceId uint64 `json:"mount_namespace_id"`
	// PidNamespaceId is the ID of the PID namespace of the process.
	PidNamespaceId uint64 `json:"pid_namespace_id"`
	// ExecId is the execution ID of the process.
	ExecId uint64 `json:"exec_id"`
	// ParentExecId is the execution ID of the parent process.
	ParentExecId uint64 `json:"parent_exec_id"`
	// ProcessName is the name of the process.
	ProcessName string `json:"process_name"`
	// Directory is the working directory of the process.
	Directory string `json:"directory"`
	// Executable is the executable path of the process.
	Executable string `json:"executable"`
	// HostDetails contains information about the host system.
	HostDetails HostDetails `json:"host_details"`
	// Kubernetes contains information about the Kubernetes pod and
	// container.
	Kubernetes Kubernetes `json:"kubernetes,omitempty"`
	// Context is the list of arguments that were passed to the system call.
	Context []Arg `json:"context"`
}

// String returns a string representation of the TarianDetectorEvent.
func (t TarianDetectorEvent) String() string {
	str := ""

	str += fmt.Sprintf("EventId: %v\n", t.EventId)
	str += fmt.Sprintf("SyscallId: %v\n", t.SyscallId)
	str += fmt.Sprintf("Timestamp: %v\n", t.Timestamp)
	str += fmt.Sprintf("ProcessorId: %v\n", t.ProcessorId)
	str += fmt.Sprintf("ThreadStartTime: %v\n", t.ThreadStartTime)
	str += fmt.Sprintf("HostProcessId: %v\n", t.HostProcessId)
	str += fmt.Sprintf("HostThreadId: %v\n", t.HostThreadId)
	str += fmt.Sprintf("HostParentProcessId: %v\n", t.HostParentProcessId)
	str += fmt.Sprintf("ProcessId: %v\n", t.ProcessId)
	str += fmt.Sprintf("ThreadId: %v\n", t.ThreadId)
	str += fmt.Sprintf("ParentProcessId: %v\n", t.ParentProcessId)
	str += fmt.Sprintf("UserId: %v\n", t.UserId)
	str += fmt.Sprintf("GroupId: %v\n", t.GroupId)
	str += fmt.Sprintf("CgroupId: %v\n", t.CgroupId)
	str += fmt.Sprintf("MountNamespaceId: %v\n", t.MountNamespaceId)
	str += fmt.Sprintf("PidNamespaceId: %v\n", t.PidNamespaceId)
	str += fmt.Sprintf("ExecId: %v\n", t.ExecId)
	str += fmt.Sprintf("ParentExecId: %v\n", t.ParentExecId)
	str += fmt.Sprintf("ProcessName: %v\n", t.ProcessName)
	str += fmt.Sprintf("Directory: %v\n", t.Directory)
	str += fmt.Sprintf("Executable: %v\n", t.Executable)
	str += fmt.Sprintf("HostDetails: {\n\n%s\n}\n", t.HostDetails)
	str += fmt.Sprintf("Kubernetes: {\n\n%s\n}\n", t.Kubernetes)
	str += fmt.Sprintf("Context: {\n\n%s\n}\n", t.Context)

	return str
}

// TarianMetaData represents the metadata associated with a event being received from the kernel.
// The first 755 bytes received from kernel in form of []byte are of type TarianMetaData
type TarianMetaData struct {
	MetaData struct {
		Event     int32    // Event identifier
		Nparams   uint8    // Number of parameters
		Syscall   int32    // System call id
		Ts        uint64   // Timestamp
		Processor uint16   // Processor number
		Task      struct { // Task information
			StartTime    uint64    // Start time
			HostPid      uint32    // Host process ID
			HostTgid     uint32    // Host thread group ID
			HostPpid     uint32    // Host parent process ID
			Pid          uint32    // Process ID of a namespace
			Tgid         uint32    // Thread group ID of a namespace
			Ppid         uint32    // Parent process ID of a namespace
			Uid          uint32    // User ID
			Gid          uint32    // Group ID
			CgroupId     uint64    // Cgroup ID
			MountNsId    uint64    // Mount namespace ID
			PidNsId      uint64    // Process namespace ID
			ExecId       uint64    // Execution ID
			ParentExecId uint64    // Parent execution ID
			Comm         [16]uint8 // Command
		}
	}
	SystemInfo struct {
		Sysname    [65]uint8 // System name
		Nodename   [65]uint8 // Node name
		Release    [65]uint8 // Release
		Version    [65]uint8 // Version
		Machine    [65]uint8 // Machine
		Domainname [65]uint8 // Domain name
	}
}

// Event returns the Event field of TarianMetaData
func (t TarianMetaData) Event() int32 {
	return t.MetaData.Event
}

// Nparams returns the Nparams field of TarianMetaData
func (t TarianMetaData) Nparams() uint8 {
	return t.MetaData.Nparams
}

// Syscall returns the Syscall field of TarianMetaData
func (t TarianMetaData) Syscall() int32 {
	return t.MetaData.Syscall
}

func (t *TarianMetaData) SetSyscall(syscallId int32) {
	t.MetaData.Syscall = syscallId
}

// Ts returns the Ts field of TarianMetaData
func (t TarianMetaData) Ts() uint64 {
	return t.MetaData.Ts
}

// Processor returns the Processor field of TarianMetaData
func (t TarianMetaData) Processor() uint16 {
	return t.MetaData.Processor
}

// StartTime returns the StartTime field of TarianMetaData
func (t TarianMetaData) StartTime() uint64 {
	return t.MetaData.Task.StartTime
}

// Hostpid returns the Hostpid field of TarianMetaData
func (t TarianMetaData) HostPid() uint32 {
	return t.MetaData.Task.HostPid
}

// HostTgid returns the HostTgid field of TarianMetaData
func (t TarianMetaData) HostTgid() uint32 {
	return t.MetaData.Task.HostTgid
}

// HostPpid returns the HostPpid field of TarianMetaData
func (t TarianMetaData) HostPpid() uint32 {
	return t.MetaData.Task.HostPpid
}

// Pid returns the Pid field of TarianMetaData
func (t TarianMetaData) Pid() uint32 {
	return t.MetaData.Task.Pid
}

// Tgid returns the Tgid field of TarianMetaData
func (t TarianMetaData) Tgid() uint32 {
	return t.MetaData.Task.Tgid
}

// Ppid returns the Ppid field of TarianMetaData
func (t TarianMetaData) Ppid() uint32 {
	return t.MetaData.Task.Ppid
}

// Uid returns the Uid field of TarianMetaData
func (t TarianMetaData) Uid() uint32 {
	return t.MetaData.Task.Uid
}

// Gid returns the Gid field of TarianMetaData
func (t TarianMetaData) Gid() uint32 {
	return t.MetaData.Task.Gid
}

// CgroupId returns the CgroupId field of TarianMetaData
func (t TarianMetaData) CgroupId() uint64 {
	return t.MetaData.Task.CgroupId
}

// MountNsId returns the MountNsId field of TarianMetaData
func (t TarianMetaData) MountNsId() uint64 {
	return t.MetaData.Task.MountNsId
}

// PidNsId returns the PidNsId field of TarianMetaData
func (t TarianMetaData) PidNsId() uint64 {
	return t.MetaData.Task.PidNsId
}

// ExecId returns the ExecId field of TarianMetaData
func (t TarianMetaData) ExecId() uint64 {
	return t.MetaData.Task.ExecId
}

// ParentExecId returns the ParentExecId field of TarianMetaData
func (t TarianMetaData) ParentExecId() uint64 {
	return t.MetaData.Task.ParentExecId
}

// Comm returns the Comm field of TarianMetaData
func (t TarianMetaData) Comm() string {
	return utils.ToString(t.MetaData.Task.Comm[:], 0, len(t.MetaData.Task.Comm))
}

// Sysname returns the Sysname field of TarianMetaData
func (t TarianMetaData) Sysname() string {
	return utils.ToString(t.SystemInfo.Sysname[:], 0, len(t.SystemInfo.Sysname))
}

// Nodename returns the Nodename field of TarianMetaData
func (t TarianMetaData) Nodename() string {
	return utils.ToString(t.SystemInfo.Nodename[:], 0, len(t.SystemInfo.Nodename))
}

// Release returns the Release field of TarianMetaData
func (t TarianMetaData) Release() string {
	return utils.ToString(t.SystemInfo.Release[:], 0, len(t.SystemInfo.Release))
}

// Version returns the Version field of TarianMetaData
func (t TarianMetaData) Version() string {
	return utils.ToString(t.SystemInfo.Version[:], 0, len(t.SystemInfo.Version))
}

// Machine returns the Machine field of TarianMetaData
func (t TarianMetaData) Machine() string {
	return utils.ToString(t.SystemInfo.Machine[:], 0, len(t.SystemInfo.Machine))
}

// Domainname returns the Domainname field of TarianMetaData
func (t TarianMetaData) Domainname() string {
	return utils.ToString(t.SystemInfo.Domainname[:], 0, len(t.SystemInfo.Domainname))
}

// TarianParamType represents the type of Tarian parameter
type TarianParamType uint32

const (
	TDT_NONE      TarianParamType = 0  // TDT_NONE represents the absence of a Tarian parameter
	TDT_U8        TarianParamType = 1  // TDT_U8 represents an 8-bit unsigned integer Tarian parameter
	TDT_U16       TarianParamType = 2  // TDT_U16 represents a 16-bit unsigned integer Tarian parameter
	TDT_U32       TarianParamType = 3  // TDT_U32 represents a 32-bit unsigned integer Tarian parameter
	TDT_U64       TarianParamType = 4  // TDT_U64 represents a 64-bit unsigned integer Tarian parameter
	TDT_S8        TarianParamType = 5  // TDT_S8 represents an 8-bit signed integer Tarian parameter
	TDT_S16       TarianParamType = 6  // TDT_S16 represents a 16-bit signed integer Tarian parameter
	TDT_S32       TarianParamType = 7  // TDT_S32 represents a 32-bit signed integer Tarian parameter
	TDT_S64       TarianParamType = 8  // TDT_S64 represents a 64-bit signed integer Tarian parameter
	TDT_IPV6      TarianParamType = 9  // TDT_IPV6 represents an IPv6 Tarian parameter
	TDT_STR       TarianParamType = 10 // TDT_STR represents a string Tarian parameter
	TDT_STR_ARR   TarianParamType = 11 // TDT_STR_ARR represents an array of strings Tarian parameter
	TDT_BYTE_ARR  TarianParamType = 12 // TDT_BYTE_ARR represents an array of bytes Tarian parameter
	TDT_IOVEC_ARR TarianParamType = 15 // TDT_IOVEC_ARR represents an array of I/O vectors Tarian parameter
	TDT_SOCKADDR  TarianParamType = 14 // TDT_SOCKADDR represents a socket address Tarian parameter
)

// String returns the string representation of the TarianParamType
func (t TarianParamType) String() string {
	switch t {
	case TDT_NONE:
		return "TDT_NONE"
	case TDT_U8:
		return "TDT_U8"
	case TDT_U16:
		return "TDT_U16"
	case TDT_U32:
		return "TDT_U32"
	case TDT_U64:
		return "TDT_U64"
	case TDT_S8:
		return "TDT_S8"
	case TDT_S16:
		return "TDT_S16"
	case TDT_S32:
		return "TDT_S32"
	case TDT_S64:
		return "TDT_S64"
	case TDT_IPV6:
		return "TDT_IPV6"
	case TDT_STR:
		return "TDT_STR"
	case TDT_STR_ARR:
		return "TDT_STR_ARR"
	case TDT_BYTE_ARR:
		return "TDT_BYTE_ARR"
	case TDT_IOVEC_ARR:
		return "TDT_IOVEC_ARR"
	case TDT_SOCKADDR:
		return "TDT_SOCKADDR"
	default:
		return fmt.Sprintf("unknown TarianParamType(%d)", int(t))
	}
}

// TarianEventsE represents the type for Tarian events enumeration.
type TarianEventsE int

const (
	TDE_SYSCALL_EXECVE_E TarianEventsE = 2 // TDE_SYSCALL_EXECVE_E represents the start of an execve syscall
	TDE_SYSCALL_EXECVE_R TarianEventsE = 3 // TDE_SYSCALL_EXECVE_R represents the return of an execve syscall

	TDE_SYSCALL_EXECVEAT_E TarianEventsE = 4 // TDE_SYSCALL_EXECVEAT_E represents the start of an execveat syscall
	TDE_SYSCALL_EXECVEAT_R TarianEventsE = 5 // TDE_SYSCALL_EXECVEAT_R represents the return of an execveat syscall

	TDE_SYSCALL_CLONE_E TarianEventsE = 6 // TDE_SYSCALL_CLONE_E represents the start of a clone syscall
	TDE_SYSCALL_CLONE_R TarianEventsE = 7 // TDE_SYSCALL_CLONE_R represents the return of a clone syscall

	TDE_SYSCALL_CLOSE_E TarianEventsE = 8 // TDE_SYSCALL_CLOSE_E represents the start of a close syscall
	TDE_SYSCALL_CLOSE_R TarianEventsE = 9 // TDE_SYSCALL_CLOSE_R represents the return of a close syscall

	TDE_SYSCALL_READ_E TarianEventsE = 10 // TDE_SYSCALL_READ_E represents the start of a read syscall
	TDE_SYSCALL_READ_R TarianEventsE = 11 // TDE_SYSCALL_READ_R represents the return of a read syscall

	TDE_SYSCALL_WRITE_E TarianEventsE = 12 // TDE_SYSCALL_WRITE_E represents the start of a write syscall
	TDE_SYSCALL_WRITE_R TarianEventsE = 13 // TDE_SYSCALL_WRITE_R represents the return of a write syscall

	TDE_SYSCALL_OPEN_E TarianEventsE = 14 // TDE_SYSCALL_OPEN_E represents the start of an open syscall
	TDE_SYSCALL_OPEN_R TarianEventsE = 15 // TDE_SYSCALL_OPEN_R represents the return of an open syscall

	TDE_SYSCALL_READV_E TarianEventsE = 16 // TDE_SYSCALL_READV_E represents the start of a readv syscall
	TDE_SYSCALL_READV_R TarianEventsE = 17 // TDE_SYSCALL_READV_R represents the return of a readv syscall

	TDE_SYSCALL_WRITEV_E TarianEventsE = 18 // TDE_SYSCALL_WRITEV_E represents the start of a writev syscall
	TDE_SYSCALL_WRITEV_R TarianEventsE = 19 // TDE_SYSCALL_WRITEV_R represents the return of a writev syscall

	TDE_SYSCALL_OPENAT_E TarianEventsE = 20 // TDE_SYSCALL_OPENAT_E represents the start of an openat syscall
	TDE_SYSCALL_OPENAT_R TarianEventsE = 21 // TDE_SYSCALL_OPENAT_R represents the return of an openat syscall

	TDE_SYSCALL_OPENAT2_E TarianEventsE = 22 // TDE_SYSCALL_OPENAT2_E represents the start of an openat2 syscall
	TDE_SYSCALL_OPENAT2_R TarianEventsE = 23 // TDE_SYSCALL_OPENAT2_R represents the return of an openat2 syscall

	TDE_SYSCALL_LISTEN_E TarianEventsE = 24 // TDE_SYSCALL_LISTEN_E represents the start of a listen syscall
	TDE_SYSCALL_LISTEN_R TarianEventsE = 25 // TDE_SYSCALL_LISTEN_R represents the return of a listen syscall

	TDE_SYSCALL_SOCKET_E TarianEventsE = 26 // TDE_SYSCALL_SOCKET_E represents the start of a socket syscall
	TDE_SYSCALL_SOCKET_R TarianEventsE = 27 // TDE_SYSCALL_SOCKET_R represents the return of a socket syscall

	TDE_SYSCALL_ACCEPT_E TarianEventsE = 28 // TDE_SYSCALL_ACCEPT_E represents the start of an accept syscall
	TDE_SYSCALL_ACCEPT_R TarianEventsE = 29 // TDE_SYSCALL_ACCEPT_R represents the return of an accept syscall

	TDE_SYSCALL_BIND_E TarianEventsE = 30 // TDE_SYSCALL_BIND_E represents the start of a bind syscall
	TDE_SYSCALL_BIND_R TarianEventsE = 31 // TDE_SYSCALL_BIND_R represents the return of a bind syscall

	TDE_SYSCALL_CONNECT_E TarianEventsE = 32 // TDE_SYSCALL_CONNECT_E represents the start of a connect syscall
	TDE_SYSCALL_CONNECT_R TarianEventsE = 33 // TDE_SYSCALL_CONNECT_R represents the return of a connect syscall
)

// String returns the string representation of TarianEventsE
func (t TarianEventsE) String() string {
	switch t {
	case TDE_SYSCALL_EXECVE_E:
		return "TDE_SYSCALL_EXECVE_E"
	case TDE_SYSCALL_EXECVE_R:
		return "TDE_SYSCALL_EXECVE_R"
	case TDE_SYSCALL_CLONE_E:
		return "TDE_SYSCALL_CLONE_E"
	case TDE_SYSCALL_CLONE_R:
		return "TDE_SYSCALL_CLONE_R"
	case TDE_SYSCALL_CLOSE_E:
		return "TDE_SYSCALL_CLOSE_E"
	case TDE_SYSCALL_CLOSE_R:
		return "TDE_SYSCALL_CLOSE_R"
	case TDE_SYSCALL_READ_E:
		return "TDE_SYSCALL_READ_E"
	case TDE_SYSCALL_READ_R:
		return "TDE_SYSCALL_READ_R"
	case TDE_SYSCALL_WRITE_E:
		return "TDE_SYSCALL_WRITE_E"
	case TDE_SYSCALL_WRITE_R:
		return "TDE_SYSCALL_WRITE_R"
	case TDE_SYSCALL_OPEN_E:
		return "TDE_SYSCALL_OPEN_E"
	case TDE_SYSCALL_OPEN_R:
		return "TDE_SYSCALL_OPEN_R"
	case TDE_SYSCALL_READV_E:
		return "TDE_SYSCALL_READV_E"
	case TDE_SYSCALL_READV_R:
		return "TDE_SYSCALL_READV_R"
	case TDE_SYSCALL_WRITEV_E:
		return "TDE_SYSCALL_WRITEV_E"
	case TDE_SYSCALL_WRITEV_R:
		return "TDE_SYSCALL_WRITEV_R"
	case TDE_SYSCALL_OPENAT_E:
		return "TDE_SYSCALL_OPENAT_E"
	case TDE_SYSCALL_OPENAT_R:
		return "TDE_SYSCALL_OPENAT_R"
	case TDE_SYSCALL_OPENAT2_E:
		return "TDE_SYSCALL_OPENAT2_E"
	case TDE_SYSCALL_OPENAT2_R:
		return "TDE_SYSCALL_OPENAT2_R"
	case TDE_SYSCALL_LISTEN_E:
		return "TDE_SYSCALL_LISTEN_E"
	case TDE_SYSCALL_LISTEN_R:
		return "TDE_SYSCALL_LISTEN_R"
	case TDE_SYSCALL_SOCKET_E:
		return "TDE_SYSCALL_SOCKET_E"
	case TDE_SYSCALL_SOCKET_R:
		return "TDE_SYSCALL_SOCKET_R"
	case TDE_SYSCALL_ACCEPT_E:
		return "TDE_SYSCALL_ACCEPT_E"
	case TDE_SYSCALL_ACCEPT_R:
		return "TDE_SYSCALL_ACCEPT_R"
	case TDE_SYSCALL_BIND_E:
		return "TDE_SYSCALL_BIND_E"
	case TDE_SYSCALL_BIND_R:
		return "TDE_SYSCALL_BIND_R"
	case TDE_SYSCALL_CONNECT_E:
		return "TDE_SYSCALL_CONNECT_E"
	case TDE_SYSCALL_CONNECT_R:
		return "TDE_SYSCALL_CONNECT_R"
	default:
		return fmt.Sprintf("unknown TarianEventsE(%d)", int(t))
	}
}
