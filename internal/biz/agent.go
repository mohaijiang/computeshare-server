package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	go_ipfs_p2p "github.com/mohaijiang/go-ipfs-p2p"
	"time"
)

type Agent struct {
	ID             uuid.UUID
	PeerId         string
	Active         bool
	LastUpdateTime time.Time
}

type AgentRepo interface {
	//db
	ListAgent(ctx context.Context) ([]*Agent, error)
	GetAgent(ctx context.Context, id uuid.UUID) (*Agent, error)
	CreateAgent(ctx context.Context, agent *Agent) error
	UpdateAgent(ctx context.Context, id uuid.UUID, agent *Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	FindByPeerId(ctx context.Context, peerId string) (*Agent, error)
	FindOneActiveAgent(ctx context.Context, cpu string, memory string) (*Agent, error)
}

type AgentUsecase struct {
	repo         AgentRepo
	log          *log.Helper
	p2pClient    *go_ipfs_p2p.P2pClient
	instanceRepo ComputeInstanceRepo
}

func NewAgentUsecase(repo AgentRepo, p2pClient *go_ipfs_p2p.P2pClient, instanceRepo ComputeInstanceRepo, logger log.Logger) *AgentUsecase {
	return &AgentUsecase{
		repo:         repo,
		p2pClient:    p2pClient,
		instanceRepo: instanceRepo,
		log:          log.NewHelper(logger),
	}
}

func (uc *AgentUsecase) List(ctx context.Context) (ps []*Agent, err error) {
	ps, err = uc.repo.ListAgent(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *AgentUsecase) Get(ctx context.Context, id uuid.UUID) (p *Agent, err error) {
	p, err = uc.repo.GetAgent(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *AgentUsecase) Create(ctx context.Context, agent *Agent) error {
	entity, err := uc.repo.FindByPeerId(ctx, agent.PeerId)
	if err != nil {
		return uc.repo.CreateAgent(ctx, agent)
	} else {
		return uc.repo.UpdateAgent(ctx, entity.ID, agent)
	}
}

func (uc *AgentUsecase) Update(ctx context.Context, id uuid.UUID, agent *Agent) error {
	return uc.repo.UpdateAgent(ctx, id, agent)
}

func (uc *AgentUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteAgent(ctx, id)
}

func (uc *AgentUsecase) FindOneActiveAgent(ctx context.Context, cpu string, memory string) (*Agent, error) {
	return uc.repo.FindOneActiveAgent(ctx, cpu, memory)
}

func (s *AgentUsecase) SyncAgentStatus() {
	ctx := context.Background()
	list, err := s.List(ctx)
	if err != nil {
		s.log.Error(err)
		return
	}

	for _, ag := range list {
		err := s.p2pClient.CheckForwardHealth("/x/ssh", ag.PeerId)
		if err != nil {
			s.log.Warnf("agent %s cannot connect.", ag.PeerId)
			ag.Active = false
			_ = s.Update(ctx, ag.ID, ag)
		} else {
			ag.Active = true
			log.Infof("agent %s check connect success.", ag.PeerId)
			_ = s.Update(ctx, ag.ID, ag)
		}

	}
}

func (uc *AgentUsecase) ListAgentInstance(ctx context.Context, peerId string) ([]*ComputeInstance, error) {
	return uc.instanceRepo.ListByPeerId(ctx, peerId)
}

func (uc *AgentUsecase) ReportInstanceStatus(ctx context.Context, instance *ComputeInstance) error {
	return uc.instanceRepo.Update(ctx, instance.ID, instance)
}
