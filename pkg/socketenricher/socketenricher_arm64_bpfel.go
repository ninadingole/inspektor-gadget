// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package socketenricher

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type socketenricherBufT struct{ Buf [32768]uint8 }

type socketenricherSocketsKey struct {
	Netns  uint32
	Family uint16
	Proto  uint8
	_      [1]byte
	Port   uint16
	_      [2]byte
}

type socketenricherSocketsValue struct {
	Mntns             uint64
	PidTgid           uint64
	UidGid            uint64
	Task              [16]int8
	Ptask             [16]int8
	Sock              uint64
	DeletionTimestamp uint64
	Cwd               [512]int8
	Exepath           [512]int8
	Ppid              uint32
	Ipv6only          int8
	_                 [3]byte
}

// loadSocketenricher returns the embedded CollectionSpec for socketenricher.
func loadSocketenricher() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_SocketenricherBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load socketenricher: %w", err)
	}

	return spec, err
}

// loadSocketenricherObjects loads socketenricher and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*socketenricherObjects
//	*socketenricherPrograms
//	*socketenricherMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadSocketenricherObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadSocketenricher()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// socketenricherSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherSpecs struct {
	socketenricherProgramSpecs
	socketenricherMapSpecs
	socketenricherVariableSpecs
}

// socketenricherProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherProgramSpecs struct {
	IgBindIpv4E   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv4_e"`
	IgBindIpv4X   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv4_x"`
	IgBindIpv6E   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv6_e"`
	IgBindIpv6X   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv6_x"`
	IgFreeIpv4E   *ebpf.ProgramSpec `ebpf:"ig_free_ipv4_e"`
	IgFreeIpv6E   *ebpf.ProgramSpec `ebpf:"ig_free_ipv6_e"`
	IgTcpCoE      *ebpf.ProgramSpec `ebpf:"ig_tcp_co_e"`
	IgTcpCoX      *ebpf.ProgramSpec `ebpf:"ig_tcp_co_x"`
	IgUdp6Sendmsg *ebpf.ProgramSpec `ebpf:"ig_udp6_sendmsg"`
	IgUdpSendmsg  *ebpf.ProgramSpec `ebpf:"ig_udp_sendmsg"`
}

// socketenricherMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherMapSpecs struct {
	Bufs              *ebpf.MapSpec `ebpf:"bufs"`
	GadgetSockets     *ebpf.MapSpec `ebpf:"gadget_sockets"`
	IgTmpSocketsValue *ebpf.MapSpec `ebpf:"ig_tmp_sockets_value"`
	Start             *ebpf.MapSpec `ebpf:"start"`
}

// socketenricherVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherVariableSpecs struct {
	DisableBpfIterators *ebpf.VariableSpec `ebpf:"disable_bpf_iterators"`
}

// socketenricherObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherObjects struct {
	socketenricherPrograms
	socketenricherMaps
	socketenricherVariables
}

func (o *socketenricherObjects) Close() error {
	return _SocketenricherClose(
		&o.socketenricherPrograms,
		&o.socketenricherMaps,
	)
}

// socketenricherMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherMaps struct {
	Bufs              *ebpf.Map `ebpf:"bufs"`
	GadgetSockets     *ebpf.Map `ebpf:"gadget_sockets"`
	IgTmpSocketsValue *ebpf.Map `ebpf:"ig_tmp_sockets_value"`
	Start             *ebpf.Map `ebpf:"start"`
}

func (m *socketenricherMaps) Close() error {
	return _SocketenricherClose(
		m.Bufs,
		m.GadgetSockets,
		m.IgTmpSocketsValue,
		m.Start,
	)
}

// socketenricherVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherVariables struct {
	DisableBpfIterators *ebpf.Variable `ebpf:"disable_bpf_iterators"`
}

// socketenricherPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherPrograms struct {
	IgBindIpv4E   *ebpf.Program `ebpf:"ig_bind_ipv4_e"`
	IgBindIpv4X   *ebpf.Program `ebpf:"ig_bind_ipv4_x"`
	IgBindIpv6E   *ebpf.Program `ebpf:"ig_bind_ipv6_e"`
	IgBindIpv6X   *ebpf.Program `ebpf:"ig_bind_ipv6_x"`
	IgFreeIpv4E   *ebpf.Program `ebpf:"ig_free_ipv4_e"`
	IgFreeIpv6E   *ebpf.Program `ebpf:"ig_free_ipv6_e"`
	IgTcpCoE      *ebpf.Program `ebpf:"ig_tcp_co_e"`
	IgTcpCoX      *ebpf.Program `ebpf:"ig_tcp_co_x"`
	IgUdp6Sendmsg *ebpf.Program `ebpf:"ig_udp6_sendmsg"`
	IgUdpSendmsg  *ebpf.Program `ebpf:"ig_udp_sendmsg"`
}

func (p *socketenricherPrograms) Close() error {
	return _SocketenricherClose(
		p.IgBindIpv4E,
		p.IgBindIpv4X,
		p.IgBindIpv6E,
		p.IgBindIpv6X,
		p.IgFreeIpv4E,
		p.IgFreeIpv6E,
		p.IgTcpCoE,
		p.IgTcpCoX,
		p.IgUdp6Sendmsg,
		p.IgUdpSendmsg,
	)
}

func _SocketenricherClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed socketenricher_arm64_bpfel.o
var _SocketenricherBytes []byte
