package serverbox

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/ramdrjn/serverbox/pkgs/statistics/pkgs/sb_stats_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Statistics struct {
	conn       *grpc.ClientConn
	statistics pb.StatisticsClient
	enabled    bool
}

func getStatsClient(conn *grpc.ClientConn) (pb.StatisticsClient, error) {
	if conn == nil {
		Log.Error("statistics server connection not valid")
		return nil, errors.New("connection not valid")
	}
	cli := pb.NewStatisticsClient(conn)
	return cli, nil
}

func InitializeStatistics(stats *Statistics) (err error) {
	var opts []grpc.DialOption
	if sbc.Conf.Statistics.Enabled == false {
		sbc.Log.Debugln("init: statistics not configured")
		return nil
	}

	host := fmt.Sprintf("%s:%d", sbc.Conf.Statistics.Host,
		sbc.Conf.Statistics.Port)

	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	sbc.Log.Debug("dialling: ", host)
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		sbc.Log.Error("server connect fail: ", err)

		//TODO - Retry connection.

		return err
	}

	cli, _ := getStatsClient(conn)

	sbc.Stats.conn = conn
	sbc.Stats.statistics = cli

	return err
}

func ShutDownStatistics(stats *Statistics) (err error) {
	if sbc.Conf.Statistics.Enabled == false {
		sbc.Log.Debugln("shut: statistics not configured")
		return nil
	}

	sbc.Stats.conn.Close()

	return nil
}

func (s *Statistics) RegisterForStats(uuid string, t string) error {
	var regType pb.RegisterReq_Type
	switch {
	case t == "server":
		regType = pb.RegisterReq_SERVER
	case t == "state":
		regType = pb.RegisterReq_STATE
	default:
		Log.Error("invalid type for registration")
		return errors.New("invalid type")
	}
	req := &pb.RegisterReq{Uuid: uuid, Type: regType}
	ctx := context.TODO()
	res, err := s.statistics.RegisterForStats(ctx, req)
	if err != nil {
		Log.Error("registration failed for type: ", t)
	}
	if err == nil && res.Enrolled == false {
		Log.Error("registration not enrolled for type: ", t)
		err = errors.New("registration not enrolled")
	}
	return err
}
