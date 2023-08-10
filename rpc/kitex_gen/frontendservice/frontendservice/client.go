// Code generated by Kitex v0.4.4. DO NOT EDIT.

package frontendservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	frontendservice "github.com/selectdb/ccr_syncer/rpc/kitex_gen/frontendservice"
	masterservice "github.com/selectdb/ccr_syncer/rpc/kitex_gen/masterservice"
	status "github.com/selectdb/ccr_syncer/rpc/kitex_gen/status"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetDbNames(ctx context.Context, params *frontendservice.TGetDbsParams, callOptions ...callopt.Option) (r *frontendservice.TGetDbsResult_, err error)
	GetTableNames(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TGetTablesResult_, err error)
	DescribeTable(ctx context.Context, params *frontendservice.TDescribeTableParams, callOptions ...callopt.Option) (r *frontendservice.TDescribeTableResult_, err error)
	DescribeTables(ctx context.Context, params *frontendservice.TDescribeTablesParams, callOptions ...callopt.Option) (r *frontendservice.TDescribeTablesResult_, err error)
	ShowVariables(ctx context.Context, params *frontendservice.TShowVariableRequest, callOptions ...callopt.Option) (r *frontendservice.TShowVariableResult_, err error)
	ReportExecStatus(ctx context.Context, params *frontendservice.TReportExecStatusParams, callOptions ...callopt.Option) (r *frontendservice.TReportExecStatusResult_, err error)
	FinishTask(ctx context.Context, request *masterservice.TFinishTaskRequest, callOptions ...callopt.Option) (r *masterservice.TMasterResult_, err error)
	Report(ctx context.Context, request *masterservice.TReportRequest, callOptions ...callopt.Option) (r *masterservice.TMasterResult_, err error)
	FetchResource(ctx context.Context, callOptions ...callopt.Option) (r *masterservice.TFetchResourceResult_, err error)
	Forward(ctx context.Context, params *frontendservice.TMasterOpRequest, callOptions ...callopt.Option) (r *frontendservice.TMasterOpResult_, err error)
	ListTableStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListTableStatusResult_, err error)
	ListTablePrivilegeStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListPrivilegesResult_, err error)
	ListSchemaPrivilegeStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListPrivilegesResult_, err error)
	ListUserPrivilegeStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListPrivilegesResult_, err error)
	UpdateExportTaskStatus(ctx context.Context, request *frontendservice.TUpdateExportTaskStatusRequest, callOptions ...callopt.Option) (r *frontendservice.TFeResult_, err error)
	LoadTxnBegin(ctx context.Context, request *frontendservice.TLoadTxnBeginRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnBeginResult_, err error)
	LoadTxnPreCommit(ctx context.Context, request *frontendservice.TLoadTxnCommitRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnCommitResult_, err error)
	LoadTxn2PC(ctx context.Context, request *frontendservice.TLoadTxn2PCRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxn2PCResult_, err error)
	LoadTxnCommit(ctx context.Context, request *frontendservice.TLoadTxnCommitRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnCommitResult_, err error)
	LoadTxnRollback(ctx context.Context, request *frontendservice.TLoadTxnRollbackRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnRollbackResult_, err error)
	BeginTxn(ctx context.Context, request *frontendservice.TBeginTxnRequest, callOptions ...callopt.Option) (r *frontendservice.TBeginTxnResult_, err error)
	CommitTxn(ctx context.Context, request *frontendservice.TCommitTxnRequest, callOptions ...callopt.Option) (r *frontendservice.TCommitTxnResult_, err error)
	RollbackTxn(ctx context.Context, request *frontendservice.TRollbackTxnRequest, callOptions ...callopt.Option) (r *frontendservice.TRollbackTxnResult_, err error)
	GetBinlog(ctx context.Context, request *frontendservice.TGetBinlogRequest, callOptions ...callopt.Option) (r *frontendservice.TGetBinlogResult_, err error)
	GetSnapshot(ctx context.Context, request *frontendservice.TGetSnapshotRequest, callOptions ...callopt.Option) (r *frontendservice.TGetSnapshotResult_, err error)
	RestoreSnapshot(ctx context.Context, request *frontendservice.TRestoreSnapshotRequest, callOptions ...callopt.Option) (r *frontendservice.TRestoreSnapshotResult_, err error)
	WaitingTxnStatus(ctx context.Context, request *frontendservice.TWaitingTxnStatusRequest, callOptions ...callopt.Option) (r *frontendservice.TWaitingTxnStatusResult_, err error)
	StreamLoadPut(ctx context.Context, request *frontendservice.TStreamLoadPutRequest, callOptions ...callopt.Option) (r *frontendservice.TStreamLoadPutResult_, err error)
	StreamLoadWithLoadStatus(ctx context.Context, request *frontendservice.TStreamLoadWithLoadStatusRequest, callOptions ...callopt.Option) (r *frontendservice.TStreamLoadWithLoadStatusResult_, err error)
	StreamLoadMultiTablePut(ctx context.Context, request *frontendservice.TStreamLoadPutRequest, callOptions ...callopt.Option) (r *frontendservice.TStreamLoadMultiTablePutResult_, err error)
	SnapshotLoaderReport(ctx context.Context, request *frontendservice.TSnapshotLoaderReportRequest, callOptions ...callopt.Option) (r *status.TStatus, err error)
	Ping(ctx context.Context, request *frontendservice.TFrontendPingFrontendRequest, callOptions ...callopt.Option) (r *frontendservice.TFrontendPingFrontendResult_, err error)
	AddColumns(ctx context.Context, request *frontendservice.TAddColumnsRequest, callOptions ...callopt.Option) (r *frontendservice.TAddColumnsResult_, err error)
	InitExternalCtlMeta(ctx context.Context, request *frontendservice.TInitExternalCtlMetaRequest, callOptions ...callopt.Option) (r *frontendservice.TInitExternalCtlMetaResult_, err error)
	FetchSchemaTableData(ctx context.Context, request *frontendservice.TFetchSchemaTableDataRequest, callOptions ...callopt.Option) (r *frontendservice.TFetchSchemaTableDataResult_, err error)
	AcquireToken(ctx context.Context, callOptions ...callopt.Option) (r *frontendservice.TMySqlLoadAcquireTokenResult_, err error)
	ConfirmUnusedRemoteFiles(ctx context.Context, request *frontendservice.TConfirmUnusedRemoteFilesRequest, callOptions ...callopt.Option) (r *frontendservice.TConfirmUnusedRemoteFilesResult_, err error)
	CheckAuth(ctx context.Context, request *frontendservice.TCheckAuthRequest, callOptions ...callopt.Option) (r *frontendservice.TCheckAuthResult_, err error)
	GetQueryStats(ctx context.Context, request *frontendservice.TGetQueryStatsRequest, callOptions ...callopt.Option) (r *frontendservice.TQueryStatsResult_, err error)
	GetTabletReplicaInfos(ctx context.Context, request *frontendservice.TGetTabletReplicaInfosRequest, callOptions ...callopt.Option) (r *frontendservice.TGetTabletReplicaInfosResult_, err error)
	GetMasterToken(ctx context.Context, request *frontendservice.TGetMasterTokenRequest, callOptions ...callopt.Option) (r *frontendservice.TGetMasterTokenResult_, err error)
	GetBinlogLag(ctx context.Context, request *frontendservice.TGetBinlogLagRequest, callOptions ...callopt.Option) (r *frontendservice.TGetBinlogLagResult_, err error)
	UpdateStatsCache(ctx context.Context, request *frontendservice.TUpdateFollowerStatsCacheRequest, callOptions ...callopt.Option) (r *status.TStatus, err error)
	GetAutoIncrementRange(ctx context.Context, request *frontendservice.TAutoIncrementRangeRequest, callOptions ...callopt.Option) (r *frontendservice.TAutoIncrementRangeResult_, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFrontendServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFrontendServiceClient struct {
	*kClient
}

func (p *kFrontendServiceClient) GetDbNames(ctx context.Context, params *frontendservice.TGetDbsParams, callOptions ...callopt.Option) (r *frontendservice.TGetDbsResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetDbNames(ctx, params)
}

func (p *kFrontendServiceClient) GetTableNames(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TGetTablesResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTableNames(ctx, params)
}

func (p *kFrontendServiceClient) DescribeTable(ctx context.Context, params *frontendservice.TDescribeTableParams, callOptions ...callopt.Option) (r *frontendservice.TDescribeTableResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescribeTable(ctx, params)
}

func (p *kFrontendServiceClient) DescribeTables(ctx context.Context, params *frontendservice.TDescribeTablesParams, callOptions ...callopt.Option) (r *frontendservice.TDescribeTablesResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescribeTables(ctx, params)
}

func (p *kFrontendServiceClient) ShowVariables(ctx context.Context, params *frontendservice.TShowVariableRequest, callOptions ...callopt.Option) (r *frontendservice.TShowVariableResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowVariables(ctx, params)
}

func (p *kFrontendServiceClient) ReportExecStatus(ctx context.Context, params *frontendservice.TReportExecStatusParams, callOptions ...callopt.Option) (r *frontendservice.TReportExecStatusResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReportExecStatus(ctx, params)
}

func (p *kFrontendServiceClient) FinishTask(ctx context.Context, request *masterservice.TFinishTaskRequest, callOptions ...callopt.Option) (r *masterservice.TMasterResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FinishTask(ctx, request)
}

func (p *kFrontendServiceClient) Report(ctx context.Context, request *masterservice.TReportRequest, callOptions ...callopt.Option) (r *masterservice.TMasterResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Report(ctx, request)
}

func (p *kFrontendServiceClient) FetchResource(ctx context.Context, callOptions ...callopt.Option) (r *masterservice.TFetchResourceResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FetchResource(ctx)
}

func (p *kFrontendServiceClient) Forward(ctx context.Context, params *frontendservice.TMasterOpRequest, callOptions ...callopt.Option) (r *frontendservice.TMasterOpResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Forward(ctx, params)
}

func (p *kFrontendServiceClient) ListTableStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListTableStatusResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListTableStatus(ctx, params)
}

func (p *kFrontendServiceClient) ListTablePrivilegeStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListPrivilegesResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListTablePrivilegeStatus(ctx, params)
}

func (p *kFrontendServiceClient) ListSchemaPrivilegeStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListPrivilegesResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListSchemaPrivilegeStatus(ctx, params)
}

func (p *kFrontendServiceClient) ListUserPrivilegeStatus(ctx context.Context, params *frontendservice.TGetTablesParams, callOptions ...callopt.Option) (r *frontendservice.TListPrivilegesResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListUserPrivilegeStatus(ctx, params)
}

func (p *kFrontendServiceClient) UpdateExportTaskStatus(ctx context.Context, request *frontendservice.TUpdateExportTaskStatusRequest, callOptions ...callopt.Option) (r *frontendservice.TFeResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateExportTaskStatus(ctx, request)
}

func (p *kFrontendServiceClient) LoadTxnBegin(ctx context.Context, request *frontendservice.TLoadTxnBeginRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnBeginResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoadTxnBegin(ctx, request)
}

func (p *kFrontendServiceClient) LoadTxnPreCommit(ctx context.Context, request *frontendservice.TLoadTxnCommitRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnCommitResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoadTxnPreCommit(ctx, request)
}

func (p *kFrontendServiceClient) LoadTxn2PC(ctx context.Context, request *frontendservice.TLoadTxn2PCRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxn2PCResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoadTxn2PC(ctx, request)
}

func (p *kFrontendServiceClient) LoadTxnCommit(ctx context.Context, request *frontendservice.TLoadTxnCommitRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnCommitResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoadTxnCommit(ctx, request)
}

func (p *kFrontendServiceClient) LoadTxnRollback(ctx context.Context, request *frontendservice.TLoadTxnRollbackRequest, callOptions ...callopt.Option) (r *frontendservice.TLoadTxnRollbackResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoadTxnRollback(ctx, request)
}

func (p *kFrontendServiceClient) BeginTxn(ctx context.Context, request *frontendservice.TBeginTxnRequest, callOptions ...callopt.Option) (r *frontendservice.TBeginTxnResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BeginTxn(ctx, request)
}

func (p *kFrontendServiceClient) CommitTxn(ctx context.Context, request *frontendservice.TCommitTxnRequest, callOptions ...callopt.Option) (r *frontendservice.TCommitTxnResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommitTxn(ctx, request)
}

func (p *kFrontendServiceClient) RollbackTxn(ctx context.Context, request *frontendservice.TRollbackTxnRequest, callOptions ...callopt.Option) (r *frontendservice.TRollbackTxnResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RollbackTxn(ctx, request)
}

func (p *kFrontendServiceClient) GetBinlog(ctx context.Context, request *frontendservice.TGetBinlogRequest, callOptions ...callopt.Option) (r *frontendservice.TGetBinlogResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBinlog(ctx, request)
}

func (p *kFrontendServiceClient) GetSnapshot(ctx context.Context, request *frontendservice.TGetSnapshotRequest, callOptions ...callopt.Option) (r *frontendservice.TGetSnapshotResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSnapshot(ctx, request)
}

func (p *kFrontendServiceClient) RestoreSnapshot(ctx context.Context, request *frontendservice.TRestoreSnapshotRequest, callOptions ...callopt.Option) (r *frontendservice.TRestoreSnapshotResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RestoreSnapshot(ctx, request)
}

func (p *kFrontendServiceClient) WaitingTxnStatus(ctx context.Context, request *frontendservice.TWaitingTxnStatusRequest, callOptions ...callopt.Option) (r *frontendservice.TWaitingTxnStatusResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.WaitingTxnStatus(ctx, request)
}

func (p *kFrontendServiceClient) StreamLoadPut(ctx context.Context, request *frontendservice.TStreamLoadPutRequest, callOptions ...callopt.Option) (r *frontendservice.TStreamLoadPutResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StreamLoadPut(ctx, request)
}

func (p *kFrontendServiceClient) StreamLoadWithLoadStatus(ctx context.Context, request *frontendservice.TStreamLoadWithLoadStatusRequest, callOptions ...callopt.Option) (r *frontendservice.TStreamLoadWithLoadStatusResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StreamLoadWithLoadStatus(ctx, request)
}

func (p *kFrontendServiceClient) StreamLoadMultiTablePut(ctx context.Context, request *frontendservice.TStreamLoadPutRequest, callOptions ...callopt.Option) (r *frontendservice.TStreamLoadMultiTablePutResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StreamLoadMultiTablePut(ctx, request)
}

func (p *kFrontendServiceClient) SnapshotLoaderReport(ctx context.Context, request *frontendservice.TSnapshotLoaderReportRequest, callOptions ...callopt.Option) (r *status.TStatus, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SnapshotLoaderReport(ctx, request)
}

func (p *kFrontendServiceClient) Ping(ctx context.Context, request *frontendservice.TFrontendPingFrontendRequest, callOptions ...callopt.Option) (r *frontendservice.TFrontendPingFrontendResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Ping(ctx, request)
}

func (p *kFrontendServiceClient) AddColumns(ctx context.Context, request *frontendservice.TAddColumnsRequest, callOptions ...callopt.Option) (r *frontendservice.TAddColumnsResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddColumns(ctx, request)
}

func (p *kFrontendServiceClient) InitExternalCtlMeta(ctx context.Context, request *frontendservice.TInitExternalCtlMetaRequest, callOptions ...callopt.Option) (r *frontendservice.TInitExternalCtlMetaResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InitExternalCtlMeta(ctx, request)
}

func (p *kFrontendServiceClient) FetchSchemaTableData(ctx context.Context, request *frontendservice.TFetchSchemaTableDataRequest, callOptions ...callopt.Option) (r *frontendservice.TFetchSchemaTableDataResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FetchSchemaTableData(ctx, request)
}

func (p *kFrontendServiceClient) AcquireToken(ctx context.Context, callOptions ...callopt.Option) (r *frontendservice.TMySqlLoadAcquireTokenResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AcquireToken(ctx)
}

func (p *kFrontendServiceClient) ConfirmUnusedRemoteFiles(ctx context.Context, request *frontendservice.TConfirmUnusedRemoteFilesRequest, callOptions ...callopt.Option) (r *frontendservice.TConfirmUnusedRemoteFilesResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ConfirmUnusedRemoteFiles(ctx, request)
}

func (p *kFrontendServiceClient) CheckAuth(ctx context.Context, request *frontendservice.TCheckAuthRequest, callOptions ...callopt.Option) (r *frontendservice.TCheckAuthResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckAuth(ctx, request)
}

func (p *kFrontendServiceClient) GetQueryStats(ctx context.Context, request *frontendservice.TGetQueryStatsRequest, callOptions ...callopt.Option) (r *frontendservice.TQueryStatsResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetQueryStats(ctx, request)
}

func (p *kFrontendServiceClient) GetTabletReplicaInfos(ctx context.Context, request *frontendservice.TGetTabletReplicaInfosRequest, callOptions ...callopt.Option) (r *frontendservice.TGetTabletReplicaInfosResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTabletReplicaInfos(ctx, request)
}

func (p *kFrontendServiceClient) GetMasterToken(ctx context.Context, request *frontendservice.TGetMasterTokenRequest, callOptions ...callopt.Option) (r *frontendservice.TGetMasterTokenResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMasterToken(ctx, request)
}

func (p *kFrontendServiceClient) GetBinlogLag(ctx context.Context, request *frontendservice.TGetBinlogLagRequest, callOptions ...callopt.Option) (r *frontendservice.TGetBinlogLagResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBinlogLag(ctx, request)
}

func (p *kFrontendServiceClient) UpdateStatsCache(ctx context.Context, request *frontendservice.TUpdateFollowerStatsCacheRequest, callOptions ...callopt.Option) (r *status.TStatus, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateStatsCache(ctx, request)
}

func (p *kFrontendServiceClient) GetAutoIncrementRange(ctx context.Context, request *frontendservice.TAutoIncrementRangeRequest, callOptions ...callopt.Option) (r *frontendservice.TAutoIncrementRangeResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAutoIncrementRange(ctx, request)
}
