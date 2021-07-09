package background

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type Background struct {
	ctx    context.Context
	wg     sync.WaitGroup
	cancel context.CancelFunc
	cmd    ports.Command
}

func NewBackground(ctx context.Context, cmd ports.Command) *Background {
	cCtx, cancel := context.WithCancel(ctx)
	return &Background{
		ctx:    cCtx,
		cancel: cancel,
		wg:     sync.WaitGroup{},
		cmd:    cmd,
	}
}

func (b *Background) StartUpdateProcess() {
	for {
		select {
		case <-b.ctx.Done():
			log.Println("parou parou")
			b.wg.Done()
			return
		case <-time.After(5 * time.Second):
			log.Println("ta processando a parada")
			if err := b.cmd.Execute(); err != nil {
				log.Println("deu ruim ai: ", err.Error())
			}
		}
	}
}

func (b *Background) AddWg(i int) {
	b.wg.Add(1)
}

func (b *Background) Stop() {
	b.cancel()
	b.wg.Wait()
}
