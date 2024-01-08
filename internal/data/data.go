package data

import (
	"context"
	"fmt"

	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//go:generate go generate ./ent

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewAgentRepo,
	NewUserRepo,
	NewStorageRepo,
	NewComputeSpecRepo,
	NewComputeInstanceRepo,
	NewComputeImageRepo,
	NewScriptRepo,
	NewScriptExecutionRecordRepo,
	NewNetworkMappingRepo,
	NewGatewayRepo,
	NewGatewayPortRepo,
	NewTaskRepo,
	NewDomainBindingRepository,
	NewS3UserRepo,
	NewStorageProviderRepo,
	NewCycleRepo,
)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

func (d *Data) GetDB() *ent.Client {
	return d.db
}

func (d *Data) getUserClient(ctx context.Context) *ent.UserClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.User
	}
	return d.db.User
}

func (d *Data) getEmployeeClient(ctx context.Context) *ent.EmployeeClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Employee
	}
	return d.db.Employee
}

func (d *Data) getTaskClient(ctx context.Context) *ent.TaskClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Task
	}
	return d.db.Task
}

func (d *Data) getS3UserClient(ctx context.Context) *ent.S3UserClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.S3User
	}
	return d.db.S3User
}

func (d *Data) getS3BucketClient(ctx context.Context) *ent.S3BucketClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.S3Bucket
	}
	return d.db.S3Bucket
}

func (d *Data) getStorageProviderClient(ctx context.Context) *ent.StorageProviderClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.StorageProvider
	}
	return d.db.StorageProvider
}

func (d *Data) getStorage(ctx context.Context) *ent.StorageClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Storage
	}
	return d.db.Storage
}

func (d *Data) getScriptExecutionRecord(ctx context.Context) *ent.ScriptExecutionRecordClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.ScriptExecutionRecord
	}
	return d.db.ScriptExecutionRecord
}

func (d *Data) getScript(ctx context.Context) *ent.ScriptClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Script
	}
	return d.db.Script
}

func (d *Data) getNetworkMapping(ctx context.Context) *ent.NetworkMappingClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.NetworkMapping
	}
	return d.db.NetworkMapping
}

func (d *Data) getGatewayPort(ctx context.Context) *ent.GatewayPortClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.GatewayPort
	}
	return d.db.GatewayPort
}

func (d *Data) getGateway(ctx context.Context) *ent.GatewayClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Gateway
	}
	return d.db.Gateway
}

func (d *Data) getDomainBinding(ctx context.Context) *ent.DomainBindingClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.DomainBinding
	}
	return d.db.DomainBinding
}

func (d *Data) getComputeSpec(ctx context.Context) *ent.ComputeSpecClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.ComputeSpec
	}
	return d.db.ComputeSpec
}

func (d *Data) getComputeImage(ctx context.Context) *ent.ComputeImageClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.ComputeImage
	}
	return d.db.ComputeImage
}

func (d *Data) getComputeInstance(ctx context.Context) *ent.ComputeInstanceClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.ComputeInstance
	}
	return d.db.ComputeInstance
}

func (d *Data) getAgent(ctx context.Context) *ent.AgentClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Agent
	}
	return d.db.Agent
}

func (d *Data) getCycle(ctx context.Context) *ent.CycleClient {
	tx, ok := getTx(ctx)
	if ok {
		return tx.Cycle
	}
	return d.db.Cycle
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func getTx(ctx context.Context) (*ent.Tx, bool) {
	if tx, ok := ctx.Value("tx").(*ent.Tx); ok {
		return tx, ok
	} else {
		return nil, false
	}
}
