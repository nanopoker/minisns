package gpool

import (
	"errors"
	"github.com/nanopoker/minisns/libs/rpc"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var (
	ErrInvalidConfig = errors.New("gpool : invalid pool config")
	ErrTimeout       = errors.New("gpool : get conn timeout")
	ErrPoolClosed    = errors.New("gpool : pool is closed")
	ErrPoolFulled    = errors.New("gpool : pool is full")
)

type Factory func() (*rpc.Client, error)

type Conn struct {
	Rpcclient *rpc.Client
	lastUsed  time.Time
}

type CliPool struct {
	factory  Factory
	conns    chan Conn
	init     uint32
	capacity uint32
	maxIdle  time.Duration
	rwl      sync.RWMutex
}

func NewPool(factory Factory, init, capacity uint32, maxIdle time.Duration) (*CliPool, error) {
	// check param
	if init < 0 || capacity <= 0 || capacity < init {
		return nil, ErrInvalidConfig
	}

	// create pool
	p := &CliPool{
		factory:  factory,
		conns:    make(chan Conn, capacity),
		init:     init,
		capacity: capacity,
		maxIdle:  maxIdle,
	}

	// create conns
	var idx uint32
	for idx = 0; idx < init; idx++ {
		c, err := factory()
		if err != nil {
			return nil, err
		}

		p.conns <- Conn{
			Rpcclient: c,
			lastUsed:  time.Now(),
		}
	}

	// fill empty
	rest := capacity - init
	for idx = 0; idx < rest; idx++ {
		p.conns <- Conn{}
	}

	return p, nil
}

func (p *CliPool) getConnsRLock() chan Conn {
	p.rwl.RLock()
	defer p.rwl.RUnlock()

	return p.conns
}

func (p *CliPool) Capacity() uint32 {
	conns := p.getConnsRLock()
	if conns == nil {
		return 0
	}
	return uint32(cap(conns))
}

func (p *CliPool) Available() uint32 {
	conns := p.getConnsRLock()
	if conns == nil {
		return 0
	}
	return uint32(len(conns))
}

func (p *CliPool) Get(ctx context.Context) (*Conn, error) {
	conns := p.getConnsRLock()
	if conns == nil {
		return nil, ErrPoolClosed
	}

	var conn Conn
	select {
	case conn = <-conns:
		// just take one
	case <-ctx.Done():
		return nil, ErrTimeout
	}

	if conn.Rpcclient != nil && p.maxIdle > 0 &&
		conn.lastUsed.Add(p.maxIdle).Before(time.Now()) {
		conn.Rpcclient.Close()
		conn.Rpcclient = nil
	}

	var err error
	if conn.Rpcclient == nil {
		conn.Rpcclient, err = p.factory()
		if err != nil {
			conns <- Conn{}
		}
	}
	return &conn, err
}

func (p *CliPool) Put(conn *Conn) error {
	newConn := Conn{
		Rpcclient: conn.Rpcclient,
		lastUsed:  time.Now(),
	}

	select {
	case p.conns <- newConn:
		// just put
	default:
		return ErrPoolFulled
	}
	return nil
}

func (p *CliPool) Close() {
	p.rwl.Lock()
	conns := p.conns
	p.conns = nil
	p.rwl.Unlock()

	if conns == nil {
		return
	}

	var idx uint32
	for idx = 0; idx < p.capacity; idx++ {
		conn := <-conns
		if conn.Rpcclient != nil {
			conn.Rpcclient.Close()
		}
	}
	close(conns)
}
