// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 Authors of Tarian & the Organization created Tarian
package network_accept

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"

	"fmt"
	"os"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags $BPF_CFLAGS -target $CURR_ARCH  -type event_data accept accept.bpf.c -- -I../../../../headers

// getEbpfObject loads the eBPF objects from the compiled code.
func getEbpfObject() (*acceptObjects, error) {
	var bpfObj acceptObjects
	// Load eBPF objects from the compiled  code into bpfObj.
	err := loadAcceptObjects(&bpfObj, nil)
	// Returns nil err if any error occurs.
	if err != nil {
		return nil, err
	}

	return &bpfObj, nil
}

// AcceptEventData is the exported data from the eBPF struct counterpart
// The intention is to use the proper Go string instead of byte arrays from C.
// It makes it simpler to use and can generate proper json.
type AcceptEventData struct {
	Args [3]uint64
}

// newAcceptEventDataFromEbpf converts an EBPF accept event to an AcceptEventData. 
//This is used to avoid having to copy the event data to a new EventData struct and to ensure that it is safe to modify the fields of the EventData struct before passing it to the event handler.
// @param e - the EBPF accept event to convert to an Accept
func newAcceptEventDataFromEbpf(e acceptEventData) *AcceptEventData {
	evt := &AcceptEventData{
		Args: [3]uint64{
			e.Args[0],
			e.Args[1],
			e.Args[2],
		},
	}
	return evt
}

// NetworkAcceptDetector is a structure to manage eBPF interaction
type NetworkAcceptDetector struct {
	ebpfLink   link.Link
	perfReader *perf.Reader
}

// NewNetworkAcceptDetector creates a new instance of NetworkAcceptDetector.
func NewNetworkAcceptDetector() *NetworkAcceptDetector {
	return &NetworkAcceptDetector{}
}

// Start starts the NetworkAcceptDetector and sets up the required eBPF hooks.


// Start initiates the NetworkAcceptDetector.
// It returns an error if the start-up process encounters any issues.
func (o *NetworkAcceptDetector) Start() error {
	// Load eBPF objects from the compiled C code.
	bpfObjs, err := getEbpfObject()
	// Returns the error if any.
	if err != nil {
		return err
	}

	// Attach a kprobe to the function "__x64_sys_accept" with the provided eBPF object.
	l, err := link.Kprobe("__x64_sys_accept", bpfObjs.KprobeAccept, nil)
	// Returns the error if any.
	if err != nil {
		return err
	}

	o.ebpfLink = l

	// Create a perf reader for the eBPF event.
	rd, err := perf.NewReader(bpfObjs.Event, os.Getpagesize())

	// Returns the error if any.
	if err != nil {
		return err
	}

	o.perfReader = rd
	return nil
}

// Close closes the eBPF link and perf reader. 
// @param o - The NetworkAcceptDetector to close. Must not be nil.
// @return An error if any is encountered or nil otherwise. If a non nil error is encountered it is returned
func (o *NetworkAcceptDetector) Close() error {
	err := o.ebpfLink.Close()
	// Returns the error if any.
	if err != nil {
		return err
	}

	return o.perfReader.Close()
}

// Read reads and returns the next eBPF event from the NetworkAcceptDetector.
// @param o - The NetworkAcceptDetector to read from. Must be non nil
func (o *NetworkAcceptDetector) Read() (*AcceptEventData, error) {
	var ebpfEvent acceptEventData
	record, err := o.perfReader.Read()
	// Returns the error if any.
	if err != nil {
		// Returns the error if any.
		if errors.Is(err, perf.ErrClosed) {
			return nil, err
		}
		return nil, err
	}

	// Read the raw sample from the record. 
	if err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &ebpfEvent); err != nil {
		return nil, err
	}
	exportedEvent := newAcceptEventDataFromEbpf(ebpfEvent)
	return exportedEvent, nil
}


func (o *NetworkAcceptDetector) ReadAsInterface() (any, error) {
	return o.Read()
}


