package metric

import (
	"context"
	"errors"
	"kafka-connect/internal/application/repo"
	"kafka-connect/internal/domain"
	"time"
)

type Processor struct {
	repo repo.Repo
}

func (p Processor) Process(l domain.List) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.processGauge(ctx, l.Alloc)
	if err != nil {
		return err
	}
	err = p.processGauge(ctx, l.FreeMemory)
	if err != nil {
		return err
	}
	err = p.processGauge(ctx, l.TotalMemory)
	if err != nil {
		return err
	}
	err = p.processCount(ctx, l.PollCount)
	if err != nil {
		return err
	}

	return
}

func (p Processor) processGauge(ctx context.Context, m domain.Metric) error {
	data, err := m.Bytes()
	if err != nil {
		return err
	}
	return p.repo.Save(ctx, m.Name, data)
}

func (p Processor) processCount(ctx context.Context, m domain.Metric) error {
	oldData, err := p.repo.Get(ctx, m.Name)
	if err != nil {
		if !errors.Is(err, repo.NotExistsKey) {
			return err
		}
	} else {
		oldM, err := domain.NewMetric(oldData)
		if err != nil {
			return err
		}
		m.Value += oldM.Value
	}
	return p.processGauge(ctx, m)
}

func NewProcess(repo repo.Repo) Processor {
	return Processor{repo}
}
