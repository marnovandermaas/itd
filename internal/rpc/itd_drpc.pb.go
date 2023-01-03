// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: itd.proto

package rpc

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_itd_proto struct{}

func (drpcEncoding_File_itd_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_itd_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_itd_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_itd_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_itd_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCITDClient interface {
	DRPCConn() drpc.Conn

	HeartRate(ctx context.Context, in *Empty) (*IntResponse, error)
	WatchHeartRate(ctx context.Context, in *Empty) (DRPCITD_WatchHeartRateClient, error)
	BatteryLevel(ctx context.Context, in *Empty) (*IntResponse, error)
	WatchBatteryLevel(ctx context.Context, in *Empty) (DRPCITD_WatchBatteryLevelClient, error)
	Motion(ctx context.Context, in *Empty) (*MotionResponse, error)
	WatchMotion(ctx context.Context, in *Empty) (DRPCITD_WatchMotionClient, error)
	StepCount(ctx context.Context, in *Empty) (*IntResponse, error)
	WatchStepCount(ctx context.Context, in *Empty) (DRPCITD_WatchStepCountClient, error)
	Version(ctx context.Context, in *Empty) (*StringResponse, error)
	Address(ctx context.Context, in *Empty) (*StringResponse, error)
	Notify(ctx context.Context, in *NotifyRequest) (*Empty, error)
	SetTime(ctx context.Context, in *SetTimeRequest) (*Empty, error)
	WeatherUpdate(ctx context.Context, in *Empty) (*Empty, error)
	FirmwareUpgrade(ctx context.Context, in *FirmwareUpgradeRequest) (DRPCITD_FirmwareUpgradeClient, error)
}

type drpcITDClient struct {
	cc drpc.Conn
}

func NewDRPCITDClient(cc drpc.Conn) DRPCITDClient {
	return &drpcITDClient{cc}
}

func (c *drpcITDClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcITDClient) HeartRate(ctx context.Context, in *Empty) (*IntResponse, error) {
	out := new(IntResponse)
	err := c.cc.Invoke(ctx, "/rpc.ITD/HeartRate", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) WatchHeartRate(ctx context.Context, in *Empty) (DRPCITD_WatchHeartRateClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.ITD/WatchHeartRate", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcITD_WatchHeartRateClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCITD_WatchHeartRateClient interface {
	drpc.Stream
	Recv() (*IntResponse, error)
}

type drpcITD_WatchHeartRateClient struct {
	drpc.Stream
}

func (x *drpcITD_WatchHeartRateClient) Recv() (*IntResponse, error) {
	m := new(IntResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcITD_WatchHeartRateClient) RecvMsg(m *IntResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

func (c *drpcITDClient) BatteryLevel(ctx context.Context, in *Empty) (*IntResponse, error) {
	out := new(IntResponse)
	err := c.cc.Invoke(ctx, "/rpc.ITD/BatteryLevel", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) WatchBatteryLevel(ctx context.Context, in *Empty) (DRPCITD_WatchBatteryLevelClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.ITD/WatchBatteryLevel", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcITD_WatchBatteryLevelClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCITD_WatchBatteryLevelClient interface {
	drpc.Stream
	Recv() (*IntResponse, error)
}

type drpcITD_WatchBatteryLevelClient struct {
	drpc.Stream
}

func (x *drpcITD_WatchBatteryLevelClient) Recv() (*IntResponse, error) {
	m := new(IntResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcITD_WatchBatteryLevelClient) RecvMsg(m *IntResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

func (c *drpcITDClient) Motion(ctx context.Context, in *Empty) (*MotionResponse, error) {
	out := new(MotionResponse)
	err := c.cc.Invoke(ctx, "/rpc.ITD/Motion", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) WatchMotion(ctx context.Context, in *Empty) (DRPCITD_WatchMotionClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.ITD/WatchMotion", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcITD_WatchMotionClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCITD_WatchMotionClient interface {
	drpc.Stream
	Recv() (*MotionResponse, error)
}

type drpcITD_WatchMotionClient struct {
	drpc.Stream
}

func (x *drpcITD_WatchMotionClient) Recv() (*MotionResponse, error) {
	m := new(MotionResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcITD_WatchMotionClient) RecvMsg(m *MotionResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

func (c *drpcITDClient) StepCount(ctx context.Context, in *Empty) (*IntResponse, error) {
	out := new(IntResponse)
	err := c.cc.Invoke(ctx, "/rpc.ITD/StepCount", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) WatchStepCount(ctx context.Context, in *Empty) (DRPCITD_WatchStepCountClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.ITD/WatchStepCount", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcITD_WatchStepCountClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCITD_WatchStepCountClient interface {
	drpc.Stream
	Recv() (*IntResponse, error)
}

type drpcITD_WatchStepCountClient struct {
	drpc.Stream
}

func (x *drpcITD_WatchStepCountClient) Recv() (*IntResponse, error) {
	m := new(IntResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcITD_WatchStepCountClient) RecvMsg(m *IntResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

func (c *drpcITDClient) Version(ctx context.Context, in *Empty) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/rpc.ITD/Version", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) Address(ctx context.Context, in *Empty) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/rpc.ITD/Address", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) Notify(ctx context.Context, in *NotifyRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.ITD/Notify", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) SetTime(ctx context.Context, in *SetTimeRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.ITD/SetTime", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) WeatherUpdate(ctx context.Context, in *Empty) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.ITD/WeatherUpdate", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcITDClient) FirmwareUpgrade(ctx context.Context, in *FirmwareUpgradeRequest) (DRPCITD_FirmwareUpgradeClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.ITD/FirmwareUpgrade", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcITD_FirmwareUpgradeClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCITD_FirmwareUpgradeClient interface {
	drpc.Stream
	Recv() (*DFUProgress, error)
}

type drpcITD_FirmwareUpgradeClient struct {
	drpc.Stream
}

func (x *drpcITD_FirmwareUpgradeClient) Recv() (*DFUProgress, error) {
	m := new(DFUProgress)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcITD_FirmwareUpgradeClient) RecvMsg(m *DFUProgress) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

type DRPCITDServer interface {
	HeartRate(context.Context, *Empty) (*IntResponse, error)
	WatchHeartRate(*Empty, DRPCITD_WatchHeartRateStream) error
	BatteryLevel(context.Context, *Empty) (*IntResponse, error)
	WatchBatteryLevel(*Empty, DRPCITD_WatchBatteryLevelStream) error
	Motion(context.Context, *Empty) (*MotionResponse, error)
	WatchMotion(*Empty, DRPCITD_WatchMotionStream) error
	StepCount(context.Context, *Empty) (*IntResponse, error)
	WatchStepCount(*Empty, DRPCITD_WatchStepCountStream) error
	Version(context.Context, *Empty) (*StringResponse, error)
	Address(context.Context, *Empty) (*StringResponse, error)
	Notify(context.Context, *NotifyRequest) (*Empty, error)
	SetTime(context.Context, *SetTimeRequest) (*Empty, error)
	WeatherUpdate(context.Context, *Empty) (*Empty, error)
	FirmwareUpgrade(*FirmwareUpgradeRequest, DRPCITD_FirmwareUpgradeStream) error
}

type DRPCITDUnimplementedServer struct{}

func (s *DRPCITDUnimplementedServer) HeartRate(context.Context, *Empty) (*IntResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) WatchHeartRate(*Empty, DRPCITD_WatchHeartRateStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) BatteryLevel(context.Context, *Empty) (*IntResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) WatchBatteryLevel(*Empty, DRPCITD_WatchBatteryLevelStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) Motion(context.Context, *Empty) (*MotionResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) WatchMotion(*Empty, DRPCITD_WatchMotionStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) StepCount(context.Context, *Empty) (*IntResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) WatchStepCount(*Empty, DRPCITD_WatchStepCountStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) Version(context.Context, *Empty) (*StringResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) Address(context.Context, *Empty) (*StringResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) Notify(context.Context, *NotifyRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) SetTime(context.Context, *SetTimeRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) WeatherUpdate(context.Context, *Empty) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCITDUnimplementedServer) FirmwareUpgrade(*FirmwareUpgradeRequest, DRPCITD_FirmwareUpgradeStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCITDDescription struct{}

func (DRPCITDDescription) NumMethods() int { return 14 }

func (DRPCITDDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/rpc.ITD/HeartRate", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					HeartRate(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.HeartRate, true
	case 1:
		return "/rpc.ITD/WatchHeartRate", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCITDServer).
					WatchHeartRate(
						in1.(*Empty),
						&drpcITD_WatchHeartRateStream{in2.(drpc.Stream)},
					)
			}, DRPCITDServer.WatchHeartRate, true
	case 2:
		return "/rpc.ITD/BatteryLevel", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					BatteryLevel(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.BatteryLevel, true
	case 3:
		return "/rpc.ITD/WatchBatteryLevel", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCITDServer).
					WatchBatteryLevel(
						in1.(*Empty),
						&drpcITD_WatchBatteryLevelStream{in2.(drpc.Stream)},
					)
			}, DRPCITDServer.WatchBatteryLevel, true
	case 4:
		return "/rpc.ITD/Motion", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					Motion(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.Motion, true
	case 5:
		return "/rpc.ITD/WatchMotion", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCITDServer).
					WatchMotion(
						in1.(*Empty),
						&drpcITD_WatchMotionStream{in2.(drpc.Stream)},
					)
			}, DRPCITDServer.WatchMotion, true
	case 6:
		return "/rpc.ITD/StepCount", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					StepCount(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.StepCount, true
	case 7:
		return "/rpc.ITD/WatchStepCount", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCITDServer).
					WatchStepCount(
						in1.(*Empty),
						&drpcITD_WatchStepCountStream{in2.(drpc.Stream)},
					)
			}, DRPCITDServer.WatchStepCount, true
	case 8:
		return "/rpc.ITD/Version", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					Version(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.Version, true
	case 9:
		return "/rpc.ITD/Address", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					Address(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.Address, true
	case 10:
		return "/rpc.ITD/Notify", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					Notify(
						ctx,
						in1.(*NotifyRequest),
					)
			}, DRPCITDServer.Notify, true
	case 11:
		return "/rpc.ITD/SetTime", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					SetTime(
						ctx,
						in1.(*SetTimeRequest),
					)
			}, DRPCITDServer.SetTime, true
	case 12:
		return "/rpc.ITD/WeatherUpdate", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCITDServer).
					WeatherUpdate(
						ctx,
						in1.(*Empty),
					)
			}, DRPCITDServer.WeatherUpdate, true
	case 13:
		return "/rpc.ITD/FirmwareUpgrade", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCITDServer).
					FirmwareUpgrade(
						in1.(*FirmwareUpgradeRequest),
						&drpcITD_FirmwareUpgradeStream{in2.(drpc.Stream)},
					)
			}, DRPCITDServer.FirmwareUpgrade, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterITD(mux drpc.Mux, impl DRPCITDServer) error {
	return mux.Register(impl, DRPCITDDescription{})
}

type DRPCITD_HeartRateStream interface {
	drpc.Stream
	SendAndClose(*IntResponse) error
}

type drpcITD_HeartRateStream struct {
	drpc.Stream
}

func (x *drpcITD_HeartRateStream) SendAndClose(m *IntResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_WatchHeartRateStream interface {
	drpc.Stream
	Send(*IntResponse) error
}

type drpcITD_WatchHeartRateStream struct {
	drpc.Stream
}

func (x *drpcITD_WatchHeartRateStream) Send(m *IntResponse) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCITD_BatteryLevelStream interface {
	drpc.Stream
	SendAndClose(*IntResponse) error
}

type drpcITD_BatteryLevelStream struct {
	drpc.Stream
}

func (x *drpcITD_BatteryLevelStream) SendAndClose(m *IntResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_WatchBatteryLevelStream interface {
	drpc.Stream
	Send(*IntResponse) error
}

type drpcITD_WatchBatteryLevelStream struct {
	drpc.Stream
}

func (x *drpcITD_WatchBatteryLevelStream) Send(m *IntResponse) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCITD_MotionStream interface {
	drpc.Stream
	SendAndClose(*MotionResponse) error
}

type drpcITD_MotionStream struct {
	drpc.Stream
}

func (x *drpcITD_MotionStream) SendAndClose(m *MotionResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_WatchMotionStream interface {
	drpc.Stream
	Send(*MotionResponse) error
}

type drpcITD_WatchMotionStream struct {
	drpc.Stream
}

func (x *drpcITD_WatchMotionStream) Send(m *MotionResponse) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCITD_StepCountStream interface {
	drpc.Stream
	SendAndClose(*IntResponse) error
}

type drpcITD_StepCountStream struct {
	drpc.Stream
}

func (x *drpcITD_StepCountStream) SendAndClose(m *IntResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_WatchStepCountStream interface {
	drpc.Stream
	Send(*IntResponse) error
}

type drpcITD_WatchStepCountStream struct {
	drpc.Stream
}

func (x *drpcITD_WatchStepCountStream) Send(m *IntResponse) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCITD_VersionStream interface {
	drpc.Stream
	SendAndClose(*StringResponse) error
}

type drpcITD_VersionStream struct {
	drpc.Stream
}

func (x *drpcITD_VersionStream) SendAndClose(m *StringResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_AddressStream interface {
	drpc.Stream
	SendAndClose(*StringResponse) error
}

type drpcITD_AddressStream struct {
	drpc.Stream
}

func (x *drpcITD_AddressStream) SendAndClose(m *StringResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_NotifyStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcITD_NotifyStream struct {
	drpc.Stream
}

func (x *drpcITD_NotifyStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_SetTimeStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcITD_SetTimeStream struct {
	drpc.Stream
}

func (x *drpcITD_SetTimeStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_WeatherUpdateStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcITD_WeatherUpdateStream struct {
	drpc.Stream
}

func (x *drpcITD_WeatherUpdateStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCITD_FirmwareUpgradeStream interface {
	drpc.Stream
	Send(*DFUProgress) error
}

type drpcITD_FirmwareUpgradeStream struct {
	drpc.Stream
}

func (x *drpcITD_FirmwareUpgradeStream) Send(m *DFUProgress) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCFSClient interface {
	DRPCConn() drpc.Conn

	RemoveAll(ctx context.Context, in *PathsRequest) (*Empty, error)
	Remove(ctx context.Context, in *PathsRequest) (*Empty, error)
	Rename(ctx context.Context, in *RenameRequest) (*Empty, error)
	MkdirAll(ctx context.Context, in *PathsRequest) (*Empty, error)
	Mkdir(ctx context.Context, in *PathsRequest) (*Empty, error)
	ReadDir(ctx context.Context, in *PathRequest) (*DirResponse, error)
	Upload(ctx context.Context, in *TransferRequest) (DRPCFS_UploadClient, error)
	Download(ctx context.Context, in *TransferRequest) (DRPCFS_DownloadClient, error)
	LoadResources(ctx context.Context, in *PathRequest) (DRPCFS_LoadResourcesClient, error)
}

type drpcFSClient struct {
	cc drpc.Conn
}

func NewDRPCFSClient(cc drpc.Conn) DRPCFSClient {
	return &drpcFSClient{cc}
}

func (c *drpcFSClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcFSClient) RemoveAll(ctx context.Context, in *PathsRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.FS/RemoveAll", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFSClient) Remove(ctx context.Context, in *PathsRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.FS/Remove", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFSClient) Rename(ctx context.Context, in *RenameRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.FS/Rename", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFSClient) MkdirAll(ctx context.Context, in *PathsRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.FS/MkdirAll", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFSClient) Mkdir(ctx context.Context, in *PathsRequest) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.FS/Mkdir", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFSClient) ReadDir(ctx context.Context, in *PathRequest) (*DirResponse, error) {
	out := new(DirResponse)
	err := c.cc.Invoke(ctx, "/rpc.FS/ReadDir", drpcEncoding_File_itd_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFSClient) Upload(ctx context.Context, in *TransferRequest) (DRPCFS_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.FS/Upload", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcFS_UploadClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCFS_UploadClient interface {
	drpc.Stream
	Recv() (*TransferProgress, error)
}

type drpcFS_UploadClient struct {
	drpc.Stream
}

func (x *drpcFS_UploadClient) Recv() (*TransferProgress, error) {
	m := new(TransferProgress)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcFS_UploadClient) RecvMsg(m *TransferProgress) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

func (c *drpcFSClient) Download(ctx context.Context, in *TransferRequest) (DRPCFS_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.FS/Download", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcFS_DownloadClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCFS_DownloadClient interface {
	drpc.Stream
	Recv() (*TransferProgress, error)
}

type drpcFS_DownloadClient struct {
	drpc.Stream
}

func (x *drpcFS_DownloadClient) Recv() (*TransferProgress, error) {
	m := new(TransferProgress)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcFS_DownloadClient) RecvMsg(m *TransferProgress) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

func (c *drpcFSClient) LoadResources(ctx context.Context, in *PathRequest) (DRPCFS_LoadResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, "/rpc.FS/LoadResources", drpcEncoding_File_itd_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcFS_LoadResourcesClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCFS_LoadResourcesClient interface {
	drpc.Stream
	Recv() (*ResourceLoadProgress, error)
}

type drpcFS_LoadResourcesClient struct {
	drpc.Stream
}

func (x *drpcFS_LoadResourcesClient) Recv() (*ResourceLoadProgress, error) {
	m := new(ResourceLoadProgress)
	if err := x.MsgRecv(m, drpcEncoding_File_itd_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcFS_LoadResourcesClient) RecvMsg(m *ResourceLoadProgress) error {
	return x.MsgRecv(m, drpcEncoding_File_itd_proto{})
}

type DRPCFSServer interface {
	RemoveAll(context.Context, *PathsRequest) (*Empty, error)
	Remove(context.Context, *PathsRequest) (*Empty, error)
	Rename(context.Context, *RenameRequest) (*Empty, error)
	MkdirAll(context.Context, *PathsRequest) (*Empty, error)
	Mkdir(context.Context, *PathsRequest) (*Empty, error)
	ReadDir(context.Context, *PathRequest) (*DirResponse, error)
	Upload(*TransferRequest, DRPCFS_UploadStream) error
	Download(*TransferRequest, DRPCFS_DownloadStream) error
	LoadResources(*PathRequest, DRPCFS_LoadResourcesStream) error
}

type DRPCFSUnimplementedServer struct{}

func (s *DRPCFSUnimplementedServer) RemoveAll(context.Context, *PathsRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) Remove(context.Context, *PathsRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) Rename(context.Context, *RenameRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) MkdirAll(context.Context, *PathsRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) Mkdir(context.Context, *PathsRequest) (*Empty, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) ReadDir(context.Context, *PathRequest) (*DirResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) Upload(*TransferRequest, DRPCFS_UploadStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) Download(*TransferRequest, DRPCFS_DownloadStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFSUnimplementedServer) LoadResources(*PathRequest, DRPCFS_LoadResourcesStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCFSDescription struct{}

func (DRPCFSDescription) NumMethods() int { return 9 }

func (DRPCFSDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/rpc.FS/RemoveAll", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFSServer).
					RemoveAll(
						ctx,
						in1.(*PathsRequest),
					)
			}, DRPCFSServer.RemoveAll, true
	case 1:
		return "/rpc.FS/Remove", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFSServer).
					Remove(
						ctx,
						in1.(*PathsRequest),
					)
			}, DRPCFSServer.Remove, true
	case 2:
		return "/rpc.FS/Rename", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFSServer).
					Rename(
						ctx,
						in1.(*RenameRequest),
					)
			}, DRPCFSServer.Rename, true
	case 3:
		return "/rpc.FS/MkdirAll", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFSServer).
					MkdirAll(
						ctx,
						in1.(*PathsRequest),
					)
			}, DRPCFSServer.MkdirAll, true
	case 4:
		return "/rpc.FS/Mkdir", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFSServer).
					Mkdir(
						ctx,
						in1.(*PathsRequest),
					)
			}, DRPCFSServer.Mkdir, true
	case 5:
		return "/rpc.FS/ReadDir", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFSServer).
					ReadDir(
						ctx,
						in1.(*PathRequest),
					)
			}, DRPCFSServer.ReadDir, true
	case 6:
		return "/rpc.FS/Upload", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCFSServer).
					Upload(
						in1.(*TransferRequest),
						&drpcFS_UploadStream{in2.(drpc.Stream)},
					)
			}, DRPCFSServer.Upload, true
	case 7:
		return "/rpc.FS/Download", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCFSServer).
					Download(
						in1.(*TransferRequest),
						&drpcFS_DownloadStream{in2.(drpc.Stream)},
					)
			}, DRPCFSServer.Download, true
	case 8:
		return "/rpc.FS/LoadResources", drpcEncoding_File_itd_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCFSServer).
					LoadResources(
						in1.(*PathRequest),
						&drpcFS_LoadResourcesStream{in2.(drpc.Stream)},
					)
			}, DRPCFSServer.LoadResources, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterFS(mux drpc.Mux, impl DRPCFSServer) error {
	return mux.Register(impl, DRPCFSDescription{})
}

type DRPCFS_RemoveAllStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcFS_RemoveAllStream struct {
	drpc.Stream
}

func (x *drpcFS_RemoveAllStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFS_RemoveStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcFS_RemoveStream struct {
	drpc.Stream
}

func (x *drpcFS_RemoveStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFS_RenameStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcFS_RenameStream struct {
	drpc.Stream
}

func (x *drpcFS_RenameStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFS_MkdirAllStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcFS_MkdirAllStream struct {
	drpc.Stream
}

func (x *drpcFS_MkdirAllStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFS_MkdirStream interface {
	drpc.Stream
	SendAndClose(*Empty) error
}

type drpcFS_MkdirStream struct {
	drpc.Stream
}

func (x *drpcFS_MkdirStream) SendAndClose(m *Empty) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFS_ReadDirStream interface {
	drpc.Stream
	SendAndClose(*DirResponse) error
}

type drpcFS_ReadDirStream struct {
	drpc.Stream
}

func (x *drpcFS_ReadDirStream) SendAndClose(m *DirResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_itd_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFS_UploadStream interface {
	drpc.Stream
	Send(*TransferProgress) error
}

type drpcFS_UploadStream struct {
	drpc.Stream
}

func (x *drpcFS_UploadStream) Send(m *TransferProgress) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCFS_DownloadStream interface {
	drpc.Stream
	Send(*TransferProgress) error
}

type drpcFS_DownloadStream struct {
	drpc.Stream
}

func (x *drpcFS_DownloadStream) Send(m *TransferProgress) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}

type DRPCFS_LoadResourcesStream interface {
	drpc.Stream
	Send(*ResourceLoadProgress) error
}

type drpcFS_LoadResourcesStream struct {
	drpc.Stream
}

func (x *drpcFS_LoadResourcesStream) Send(m *ResourceLoadProgress) error {
	return x.MsgSend(m, drpcEncoding_File_itd_proto{})
}
