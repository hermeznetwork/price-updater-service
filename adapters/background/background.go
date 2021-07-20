package background

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type Background struct {
	ctx    context.Context
	wg     sync.WaitGroup
	cancel context.CancelFunc
	cmd    ports.Command
	cfg    config.Main
}

func NewBackground(ctx context.Context, cmd ports.Command, cfg config.Main) *Background {
	cCtx, cancel := context.WithCancel(ctx)
	return &Background{
		ctx:    cCtx,
		cancel: cancel,
		wg:     sync.WaitGroup{},
		cmd:    cmd,
		cfg:    cfg,
	}
}

func (b *Background) StartUpdateProcess() {
	for {
		select {
		case <-b.ctx.Done():
			log.Println("graceful shutdown...")
			b.wg.Done()
			return
		case <-time.After(b.cfg.TimeToUpdate):
			log.Println("Executing...")
			if err := b.cmd.Execute(); err != nil {
				log.Println("error while try executing: ", err.Error())
			}
			log.Println("...done.")
		}
	}
}

func (b *Background) AddWg(i int) {
	b.wg.Add(i)
}

func (b *Background) Stop() {
	b.cancel()
	b.wg.Wait()
}
