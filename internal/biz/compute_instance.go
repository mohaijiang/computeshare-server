package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ipfs/kubo/core"
	clientcomputev1 "github.com/mohaijiang/computeshare-client/api/compute/v1"
	pb "github.com/mohaijiang/computeshare-client/api/network/v1"
	"github.com/mohaijiang/computeshare-server/internal/global"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/samber/lo"
	"golang.org/x/exp/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ComputeSpec struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
}

type ComputeInstance struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	Port  string `json:"port,omitempty"`
	// ExpirationTime holds the value of the "expiration_time" field.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// 0: 启动中,1:运行中,2:连接中断, 3:过期
	Status int8 `json:"status,omitempty"`
	// 容器id
	ContainerID string `json:"container_id,omitempty"`
	// p2p agent Id
	PeerID  string `json:"peer_id,omitempty"`
	Command string `json:"command,omitempty"`
}

const (
	InstanceStatusStarting int8 = iota
	InstanceStatusRunning
	InstanceStatusTerminal
	InstanceStatusExpire
)

type ComputeInstanceCreate struct {
	SpecId   int32
	ImageId  int32
	Duration int32
	Name     string
}

type ComputeImage struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 显示名
	Name string `json:"name,omitempty"`
	// 镜像名
	Image string `json:"image,omitempty"`
	// 版本名
	Tag string `json:"tag,omitempty"`
	// 端口号
	Port    int32 `json:"port,omitempty"`
	Command string
}

type ComputeSpecRepo interface {
	List(ctx context.Context) ([]*ComputeSpec, error)
	Get(ctx context.Context, id int32) (*ComputeSpec, error)
}

type ComputeInstanceRepo interface {
	List(ctx context.Context, owner string) ([]*ComputeInstance, error)
	Create(ctx context.Context, instance *ComputeInstance) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, instance *ComputeInstance) error
	Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error)
}

type ComputeImageRepo interface {
	List(ctx context.Context) ([]*ComputeImage, error)
	Get(ctx context.Context, id int32) (*ComputeImage, error)
}

type ComputeInstanceUsercase struct {
	specRepo     ComputeSpecRepo
	instanceRepo ComputeInstanceRepo
	imageRepo    ComputeImageRepo
	agentRepo    AgentRepo
	p2pUsecase   *P2PUsecase
	log          *log.Helper
}

func NewComputeInstanceUsercase(
	specRepo ComputeSpecRepo,
	instanceRepo ComputeInstanceRepo,
	imageRepo ComputeImageRepo,
	ipfsNode *core.IpfsNode,
	agentRepo AgentRepo,
	p2pUsecase *P2PUsecase,
	logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:     specRepo,
		instanceRepo: instanceRepo,
		imageRepo:    imageRepo,
		p2pUsecase:   p2pUsecase,
		agentRepo:    agentRepo,
		log:          log.NewHelper(logger),
	}
}

func (uc *ComputeInstanceUsercase) ListComputeSpec(ctx context.Context) ([]*ComputeSpec, error) {
	return uc.specRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) ListComputeImage(ctx context.Context) ([]*ComputeImage, error) {
	return uc.imageRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) Create(ctx context.Context, cic *ComputeInstanceCreate) (*ComputeInstance, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("cannot get user ID")
	}

	computeSpec, err := uc.specRepo.Get(ctx, cic.SpecId)
	if err != nil {
		return nil, err
	}
	computeImage, err := uc.imageRepo.Get(ctx, cic.ImageId)
	if err != nil {
		return nil, err
	}

	instance := &ComputeInstance{
		Owner:          claim.UserID,
		Name:           cic.Name,
		Core:           computeSpec.Core,
		Memory:         computeSpec.Memory,
		Port:           fmt.Sprintf("%d", computeImage.Port),
		Image:          fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag),
		Command:        computeImage.Command,
		ExpirationTime: time.Now().AddDate(0, int(cic.Duration), 0),
		Status:         InstanceStatusStarting,
	}

	err = uc.instanceRepo.Create(ctx, instance)

	// 选择一个agent节点进行通信
	agent, err := uc.agentRepo.FindOneActiveAgent(ctx, instance.Core, instance.Memory)
	if err != nil {
		return nil, err
	}
	go uc.CreateInstanceOnAgent(agent.PeerId, instance)

	return instance, err
}

func (uc *ComputeInstanceUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.instanceRepo.Delete(ctx, id)
}

func (uc *ComputeInstanceUsercase) CreateInstanceOnAgent(peerId string, instance *ComputeInstance) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*20)

	vmClient, cleanup, err := uc.getVmClient(peerId)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
	defer cleanup()

	reply, err := vmClient.CreateVm(ctx, &clientcomputev1.CreateVmRequest{
		Image:   instance.Image,
		Port:    instance.Port,
		Command: strings.Fields(instance.Command),
	})
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
	fmt.Println("containerId:", reply.GetId())

	instance.Status = InstanceStatusRunning
	instance.PeerID = peerId
	instance.ContainerID = reply.GetId()

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
}

func (uc *ComputeInstanceUsercase) getVmClient(peerId string) (clientcomputev1.VmHTTPClient, func(), error) {
	ip, port, err := uc.createP2pForward(peerId)
	if err != nil {
		return nil, nil, err
	}

	time.Sleep(time.Second * 2)

	client, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(fmt.Sprintf("%s:%s", ip, port)),
		transhttp.WithTimeout(time.Second*10),
	)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return nil, nil, err
	}

	vmClient := clientcomputev1.NewVmHTTPClient(client)
	return vmClient, func() {
		_ = client.Close()
	}, nil
}

func (uc *ComputeInstanceUsercase) createP2pForward(peerId string) (string, string, error) {
	ctx := context.Background()
	pingOk := uc.p2pUsecase.Ping(ctx, peerId)

	fmt.Println("pingOk: ", pingOk)
	if !pingOk {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Errorf("无法与%s完成ping", peerId)
		return "", "", nil
	}

	list, err := uc.p2pUsecase.ListListen(ctx, nil)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error("查询p2p 列表失败")
		return "", "", nil
	}

	t, find := lo.Find(list.Result, func(item *pb.ListenReply) bool {
		if item == nil {
			return false
		}
		return item.TargetAddress == fmt.Sprintf("/p2p/%s", peerId)
	})

	if find {
		listenAddress := t.ListenAddress
		// 定义正则表达式模式，用于匹配IP地址和端口号
		pattern := `\/ip4\/([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)\/tcp\/([0-9]+)`

		// 编译正则表达式
		regex := regexp.MustCompile(pattern)

		// 使用正则表达式来提取IP地址和端口号
		matches := regex.FindStringSubmatch(listenAddress)
		if len(matches) >= 3 {
			ip := matches[1]   // 第一个匹配组为IP地址
			port := matches[2] // 第二个匹配组为端口号

			fmt.Printf("IP地址: %s\n", ip)
			fmt.Printf("端口号: %s\n", port)
			return ip, port, nil
		} else {
			fmt.Println("无法提取IP地址和端口号")
		}
	}

	listenIp := "127.0.0.1"
	listenPort := rand.Intn(9999) + 30000

	listenOpt := fmt.Sprintf("/ip4/%s/tcp/%d", listenIp, listenPort)
	listen, err := ma.NewMultiaddr(listenOpt)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return "", "", nil
	}
	targetOpt := fmt.Sprintf("/p2p/%s", peerId)
	proto := "/x/ssh"

	err = uc.p2pUsecase.CheckPort(listen)
	if err != nil {
		_ = uc.p2pUsecase.CloseListen(ctx, proto, listenOpt, targetOpt)
	}
	err = uc.p2pUsecase.CreateForward(ctx, proto, listenOpt, targetOpt)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return "", "", nil
	}
	return listenIp, strconv.Itoa(listenPort), nil
}

func (uc *ComputeInstanceUsercase) ListComputeInstance(ctx context.Context, owner string) ([]*ComputeInstance, error) {
	return uc.instanceRepo.List(ctx, owner)
}

func (uc *ComputeInstanceUsercase) Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error) {
	return uc.instanceRepo.Get(ctx, id)
}

func (uc *ComputeInstanceUsercase) Start(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	if instance.ContainerID == "" || instance.PeerID == "" {
		return fmt.Errorf("instance is not avaliable")
	}

	vmClient, cleanup, err := uc.getVmClient(instance.PeerID)
	if err != nil {
		return err
	}
	defer cleanup()

	_, err = vmClient.StartVm(ctx, &clientcomputev1.GetVmRequest{
		Id: instance.ContainerID,
	})

	return err

}

func (uc *ComputeInstanceUsercase) Stop(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	if instance.ContainerID == "" || instance.PeerID == "" {
		return fmt.Errorf("instance is not avaliable")
	}

	vmClient, cleanup, err := uc.getVmClient(instance.PeerID)
	if err != nil {
		return err
	}
	defer cleanup()

	_, err = vmClient.StopVm(ctx, &clientcomputev1.GetVmRequest{
		Id: instance.ContainerID,
	})
	return err
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (uc *ComputeInstanceUsercase) Terminal(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	instanceId, err := uuid.Parse(r.Form.Get("instanceId"))
	if err != nil {
		return
	}
	instance, err := uc.Get(context.Background(), instanceId)
	if err != nil {
		return
	}

	if instance.PeerID == "" {
		return
	}

	ip, port, err := uc.createP2pForward(instance.PeerID)
	if err != nil {
		return
	}

	// websocket握手
	// 建立与目标WebSocket服务器的连接
	targetConn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/v1/vm/%s/terminal?container=%s&workdir=/bin", ip, port, instanceId, instance.ContainerID), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer targetConn.Close()

	// 升级客户端连接为WebSocket
	clientConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer clientConn.Close()

	// 开始在两个WebSocket连接之间传递消息
	go func() {
		for {
			msgType, msg, err := clientConn.ReadMessage()
			if err != nil {
				return
			}
			if err := targetConn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	}()

	for {
		msgType, msg, err := targetConn.ReadMessage()
		if err != nil {
			return
		}
		if err := clientConn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
