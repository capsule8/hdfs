package hadoop_hdfs

import (
	"sync"

	proto "github.com/golang/protobuf/proto"
)

var loadOnce sync.Once

func LoadAPI() {
	loadOnce.Do(func() {
		proto.RegisterType((*GetReplicaVisibleLengthRequestProto)(nil), "hadoop.hdfs.GetReplicaVisibleLengthRequestProto")
		proto.RegisterType((*GetReplicaVisibleLengthResponseProto)(nil), "hadoop.hdfs.GetReplicaVisibleLengthResponseProto")
		proto.RegisterType((*RefreshNamenodesRequestProto)(nil), "hadoop.hdfs.RefreshNamenodesRequestProto")
		proto.RegisterType((*RefreshNamenodesResponseProto)(nil), "hadoop.hdfs.RefreshNamenodesResponseProto")
		proto.RegisterType((*DeleteBlockPoolRequestProto)(nil), "hadoop.hdfs.DeleteBlockPoolRequestProto")
		proto.RegisterType((*DeleteBlockPoolResponseProto)(nil), "hadoop.hdfs.DeleteBlockPoolResponseProto")
		proto.RegisterType((*GetBlockLocalPathInfoRequestProto)(nil), "hadoop.hdfs.GetBlockLocalPathInfoRequestProto")
		proto.RegisterType((*GetBlockLocalPathInfoResponseProto)(nil), "hadoop.hdfs.GetBlockLocalPathInfoResponseProto")
		proto.RegisterType((*ShutdownDatanodeRequestProto)(nil), "hadoop.hdfs.ShutdownDatanodeRequestProto")
		proto.RegisterType((*ShutdownDatanodeResponseProto)(nil), "hadoop.hdfs.ShutdownDatanodeResponseProto")
		proto.RegisterType((*EvictWritersRequestProto)(nil), "hadoop.hdfs.EvictWritersRequestProto")
		proto.RegisterType((*EvictWritersResponseProto)(nil), "hadoop.hdfs.EvictWritersResponseProto")
		proto.RegisterType((*GetDatanodeInfoRequestProto)(nil), "hadoop.hdfs.GetDatanodeInfoRequestProto")
		proto.RegisterType((*GetDatanodeInfoResponseProto)(nil), "hadoop.hdfs.GetDatanodeInfoResponseProto")
		proto.RegisterType((*TriggerBlockReportRequestProto)(nil), "hadoop.hdfs.TriggerBlockReportRequestProto")
		proto.RegisterType((*TriggerBlockReportResponseProto)(nil), "hadoop.hdfs.TriggerBlockReportResponseProto")
		proto.RegisterType((*GetBalancerBandwidthRequestProto)(nil), "hadoop.hdfs.GetBalancerBandwidthRequestProto")
		proto.RegisterType((*GetBalancerBandwidthResponseProto)(nil), "hadoop.hdfs.GetBalancerBandwidthResponseProto")
		proto.RegisterFile("ClientDatanodeProtocol.proto", fileDescriptor6)
		proto.RegisterType((*GetBlockLocationsRequestProto)(nil), "hadoop.hdfs.GetBlockLocationsRequestProto")
		proto.RegisterType((*GetBlockLocationsResponseProto)(nil), "hadoop.hdfs.GetBlockLocationsResponseProto")
		proto.RegisterType((*GetServerDefaultsRequestProto)(nil), "hadoop.hdfs.GetServerDefaultsRequestProto")
		proto.RegisterType((*GetServerDefaultsResponseProto)(nil), "hadoop.hdfs.GetServerDefaultsResponseProto")
		proto.RegisterType((*CreateRequestProto)(nil), "hadoop.hdfs.CreateRequestProto")
		proto.RegisterType((*CreateResponseProto)(nil), "hadoop.hdfs.CreateResponseProto")
		proto.RegisterType((*AppendRequestProto)(nil), "hadoop.hdfs.AppendRequestProto")
		proto.RegisterType((*AppendResponseProto)(nil), "hadoop.hdfs.AppendResponseProto")
		proto.RegisterType((*SetReplicationRequestProto)(nil), "hadoop.hdfs.SetReplicationRequestProto")
		proto.RegisterType((*SetReplicationResponseProto)(nil), "hadoop.hdfs.SetReplicationResponseProto")
		proto.RegisterType((*SetStoragePolicyRequestProto)(nil), "hadoop.hdfs.SetStoragePolicyRequestProto")
		proto.RegisterType((*SetStoragePolicyResponseProto)(nil), "hadoop.hdfs.SetStoragePolicyResponseProto")
		proto.RegisterType((*UnsetStoragePolicyRequestProto)(nil), "hadoop.hdfs.UnsetStoragePolicyRequestProto")
		proto.RegisterType((*UnsetStoragePolicyResponseProto)(nil), "hadoop.hdfs.UnsetStoragePolicyResponseProto")
		proto.RegisterType((*GetStoragePolicyRequestProto)(nil), "hadoop.hdfs.GetStoragePolicyRequestProto")
		proto.RegisterType((*GetStoragePolicyResponseProto)(nil), "hadoop.hdfs.GetStoragePolicyResponseProto")
		proto.RegisterType((*GetStoragePoliciesRequestProto)(nil), "hadoop.hdfs.GetStoragePoliciesRequestProto")
		proto.RegisterType((*GetStoragePoliciesResponseProto)(nil), "hadoop.hdfs.GetStoragePoliciesResponseProto")
		proto.RegisterType((*SetPermissionRequestProto)(nil), "hadoop.hdfs.SetPermissionRequestProto")
		proto.RegisterType((*SetPermissionResponseProto)(nil), "hadoop.hdfs.SetPermissionResponseProto")
		proto.RegisterType((*SetOwnerRequestProto)(nil), "hadoop.hdfs.SetOwnerRequestProto")
		proto.RegisterType((*SetOwnerResponseProto)(nil), "hadoop.hdfs.SetOwnerResponseProto")
		proto.RegisterType((*AbandonBlockRequestProto)(nil), "hadoop.hdfs.AbandonBlockRequestProto")
		proto.RegisterType((*AbandonBlockResponseProto)(nil), "hadoop.hdfs.AbandonBlockResponseProto")
		proto.RegisterType((*AddBlockRequestProto)(nil), "hadoop.hdfs.AddBlockRequestProto")
		proto.RegisterType((*AddBlockResponseProto)(nil), "hadoop.hdfs.AddBlockResponseProto")
		proto.RegisterType((*GetAdditionalDatanodeRequestProto)(nil), "hadoop.hdfs.GetAdditionalDatanodeRequestProto")
		proto.RegisterType((*GetAdditionalDatanodeResponseProto)(nil), "hadoop.hdfs.GetAdditionalDatanodeResponseProto")
		proto.RegisterType((*CompleteRequestProto)(nil), "hadoop.hdfs.CompleteRequestProto")
		proto.RegisterType((*CompleteResponseProto)(nil), "hadoop.hdfs.CompleteResponseProto")
		proto.RegisterType((*ReportBadBlocksRequestProto)(nil), "hadoop.hdfs.ReportBadBlocksRequestProto")
		proto.RegisterType((*ReportBadBlocksResponseProto)(nil), "hadoop.hdfs.ReportBadBlocksResponseProto")
		proto.RegisterType((*ConcatRequestProto)(nil), "hadoop.hdfs.ConcatRequestProto")
		proto.RegisterType((*ConcatResponseProto)(nil), "hadoop.hdfs.ConcatResponseProto")
		proto.RegisterType((*TruncateRequestProto)(nil), "hadoop.hdfs.TruncateRequestProto")
		proto.RegisterType((*TruncateResponseProto)(nil), "hadoop.hdfs.TruncateResponseProto")
		proto.RegisterType((*RenameRequestProto)(nil), "hadoop.hdfs.RenameRequestProto")
		proto.RegisterType((*RenameResponseProto)(nil), "hadoop.hdfs.RenameResponseProto")
		proto.RegisterType((*Rename2RequestProto)(nil), "hadoop.hdfs.Rename2RequestProto")
		proto.RegisterType((*Rename2ResponseProto)(nil), "hadoop.hdfs.Rename2ResponseProto")
		proto.RegisterType((*DeleteRequestProto)(nil), "hadoop.hdfs.DeleteRequestProto")
		proto.RegisterType((*DeleteResponseProto)(nil), "hadoop.hdfs.DeleteResponseProto")
		proto.RegisterType((*MkdirsRequestProto)(nil), "hadoop.hdfs.MkdirsRequestProto")
		proto.RegisterType((*MkdirsResponseProto)(nil), "hadoop.hdfs.MkdirsResponseProto")
		proto.RegisterType((*GetListingRequestProto)(nil), "hadoop.hdfs.GetListingRequestProto")
		proto.RegisterType((*GetListingResponseProto)(nil), "hadoop.hdfs.GetListingResponseProto")
		proto.RegisterType((*GetSnapshottableDirListingRequestProto)(nil), "hadoop.hdfs.GetSnapshottableDirListingRequestProto")
		proto.RegisterType((*GetSnapshottableDirListingResponseProto)(nil), "hadoop.hdfs.GetSnapshottableDirListingResponseProto")
		proto.RegisterType((*GetSnapshotDiffReportRequestProto)(nil), "hadoop.hdfs.GetSnapshotDiffReportRequestProto")
		proto.RegisterType((*GetSnapshotDiffReportResponseProto)(nil), "hadoop.hdfs.GetSnapshotDiffReportResponseProto")
		proto.RegisterType((*RenewLeaseRequestProto)(nil), "hadoop.hdfs.RenewLeaseRequestProto")
		proto.RegisterType((*RenewLeaseResponseProto)(nil), "hadoop.hdfs.RenewLeaseResponseProto")
		proto.RegisterType((*RecoverLeaseRequestProto)(nil), "hadoop.hdfs.RecoverLeaseRequestProto")
		proto.RegisterType((*RecoverLeaseResponseProto)(nil), "hadoop.hdfs.RecoverLeaseResponseProto")
		proto.RegisterType((*GetFsStatusRequestProto)(nil), "hadoop.hdfs.GetFsStatusRequestProto")
		proto.RegisterType((*GetFsStatsResponseProto)(nil), "hadoop.hdfs.GetFsStatsResponseProto")
		proto.RegisterType((*GetDatanodeReportRequestProto)(nil), "hadoop.hdfs.GetDatanodeReportRequestProto")
		proto.RegisterType((*GetDatanodeReportResponseProto)(nil), "hadoop.hdfs.GetDatanodeReportResponseProto")
		proto.RegisterType((*GetDatanodeStorageReportRequestProto)(nil), "hadoop.hdfs.GetDatanodeStorageReportRequestProto")
		proto.RegisterType((*DatanodeStorageReportProto)(nil), "hadoop.hdfs.DatanodeStorageReportProto")
		proto.RegisterType((*GetDatanodeStorageReportResponseProto)(nil), "hadoop.hdfs.GetDatanodeStorageReportResponseProto")
		proto.RegisterType((*GetPreferredBlockSizeRequestProto)(nil), "hadoop.hdfs.GetPreferredBlockSizeRequestProto")
		proto.RegisterType((*GetPreferredBlockSizeResponseProto)(nil), "hadoop.hdfs.GetPreferredBlockSizeResponseProto")
		proto.RegisterType((*SetSafeModeRequestProto)(nil), "hadoop.hdfs.SetSafeModeRequestProto")
		proto.RegisterType((*SetSafeModeResponseProto)(nil), "hadoop.hdfs.SetSafeModeResponseProto")
		proto.RegisterType((*SaveNamespaceRequestProto)(nil), "hadoop.hdfs.SaveNamespaceRequestProto")
		proto.RegisterType((*SaveNamespaceResponseProto)(nil), "hadoop.hdfs.SaveNamespaceResponseProto")
		proto.RegisterType((*RollEditsRequestProto)(nil), "hadoop.hdfs.RollEditsRequestProto")
		proto.RegisterType((*RollEditsResponseProto)(nil), "hadoop.hdfs.RollEditsResponseProto")
		proto.RegisterType((*RestoreFailedStorageRequestProto)(nil), "hadoop.hdfs.RestoreFailedStorageRequestProto")
		proto.RegisterType((*RestoreFailedStorageResponseProto)(nil), "hadoop.hdfs.RestoreFailedStorageResponseProto")
		proto.RegisterType((*RefreshNodesRequestProto)(nil), "hadoop.hdfs.RefreshNodesRequestProto")
		proto.RegisterType((*RefreshNodesResponseProto)(nil), "hadoop.hdfs.RefreshNodesResponseProto")
		proto.RegisterType((*FinalizeUpgradeRequestProto)(nil), "hadoop.hdfs.FinalizeUpgradeRequestProto")
		proto.RegisterType((*FinalizeUpgradeResponseProto)(nil), "hadoop.hdfs.FinalizeUpgradeResponseProto")
		proto.RegisterType((*RollingUpgradeRequestProto)(nil), "hadoop.hdfs.RollingUpgradeRequestProto")
		proto.RegisterType((*RollingUpgradeInfoProto)(nil), "hadoop.hdfs.RollingUpgradeInfoProto")
		proto.RegisterType((*RollingUpgradeResponseProto)(nil), "hadoop.hdfs.RollingUpgradeResponseProto")
		proto.RegisterType((*ListCorruptFileBlocksRequestProto)(nil), "hadoop.hdfs.ListCorruptFileBlocksRequestProto")
		proto.RegisterType((*ListCorruptFileBlocksResponseProto)(nil), "hadoop.hdfs.ListCorruptFileBlocksResponseProto")
		proto.RegisterType((*MetaSaveRequestProto)(nil), "hadoop.hdfs.MetaSaveRequestProto")
		proto.RegisterType((*MetaSaveResponseProto)(nil), "hadoop.hdfs.MetaSaveResponseProto")
		proto.RegisterType((*GetFileInfoRequestProto)(nil), "hadoop.hdfs.GetFileInfoRequestProto")
		proto.RegisterType((*GetFileInfoResponseProto)(nil), "hadoop.hdfs.GetFileInfoResponseProto")
		proto.RegisterType((*IsFileClosedRequestProto)(nil), "hadoop.hdfs.IsFileClosedRequestProto")
		proto.RegisterType((*IsFileClosedResponseProto)(nil), "hadoop.hdfs.IsFileClosedResponseProto")
		proto.RegisterType((*CacheDirectiveInfoProto)(nil), "hadoop.hdfs.CacheDirectiveInfoProto")
		proto.RegisterType((*CacheDirectiveInfoExpirationProto)(nil), "hadoop.hdfs.CacheDirectiveInfoExpirationProto")
		proto.RegisterType((*CacheDirectiveStatsProto)(nil), "hadoop.hdfs.CacheDirectiveStatsProto")
		proto.RegisterType((*AddCacheDirectiveRequestProto)(nil), "hadoop.hdfs.AddCacheDirectiveRequestProto")
		proto.RegisterType((*AddCacheDirectiveResponseProto)(nil), "hadoop.hdfs.AddCacheDirectiveResponseProto")
		proto.RegisterType((*ModifyCacheDirectiveRequestProto)(nil), "hadoop.hdfs.ModifyCacheDirectiveRequestProto")
		proto.RegisterType((*ModifyCacheDirectiveResponseProto)(nil), "hadoop.hdfs.ModifyCacheDirectiveResponseProto")
		proto.RegisterType((*RemoveCacheDirectiveRequestProto)(nil), "hadoop.hdfs.RemoveCacheDirectiveRequestProto")
		proto.RegisterType((*RemoveCacheDirectiveResponseProto)(nil), "hadoop.hdfs.RemoveCacheDirectiveResponseProto")
		proto.RegisterType((*ListCacheDirectivesRequestProto)(nil), "hadoop.hdfs.ListCacheDirectivesRequestProto")
		proto.RegisterType((*CacheDirectiveEntryProto)(nil), "hadoop.hdfs.CacheDirectiveEntryProto")
		proto.RegisterType((*ListCacheDirectivesResponseProto)(nil), "hadoop.hdfs.ListCacheDirectivesResponseProto")
		proto.RegisterType((*CachePoolInfoProto)(nil), "hadoop.hdfs.CachePoolInfoProto")
		proto.RegisterType((*CachePoolStatsProto)(nil), "hadoop.hdfs.CachePoolStatsProto")
		proto.RegisterType((*AddCachePoolRequestProto)(nil), "hadoop.hdfs.AddCachePoolRequestProto")
		proto.RegisterType((*AddCachePoolResponseProto)(nil), "hadoop.hdfs.AddCachePoolResponseProto")
		proto.RegisterType((*ModifyCachePoolRequestProto)(nil), "hadoop.hdfs.ModifyCachePoolRequestProto")
		proto.RegisterType((*ModifyCachePoolResponseProto)(nil), "hadoop.hdfs.ModifyCachePoolResponseProto")
		proto.RegisterType((*RemoveCachePoolRequestProto)(nil), "hadoop.hdfs.RemoveCachePoolRequestProto")
		proto.RegisterType((*RemoveCachePoolResponseProto)(nil), "hadoop.hdfs.RemoveCachePoolResponseProto")
		proto.RegisterType((*ListCachePoolsRequestProto)(nil), "hadoop.hdfs.ListCachePoolsRequestProto")
		proto.RegisterType((*ListCachePoolsResponseProto)(nil), "hadoop.hdfs.ListCachePoolsResponseProto")
		proto.RegisterType((*CachePoolEntryProto)(nil), "hadoop.hdfs.CachePoolEntryProto")
		proto.RegisterType((*GetFileLinkInfoRequestProto)(nil), "hadoop.hdfs.GetFileLinkInfoRequestProto")
		proto.RegisterType((*GetFileLinkInfoResponseProto)(nil), "hadoop.hdfs.GetFileLinkInfoResponseProto")
		proto.RegisterType((*GetContentSummaryRequestProto)(nil), "hadoop.hdfs.GetContentSummaryRequestProto")
		proto.RegisterType((*GetContentSummaryResponseProto)(nil), "hadoop.hdfs.GetContentSummaryResponseProto")
		proto.RegisterType((*GetQuotaUsageRequestProto)(nil), "hadoop.hdfs.GetQuotaUsageRequestProto")
		proto.RegisterType((*GetQuotaUsageResponseProto)(nil), "hadoop.hdfs.GetQuotaUsageResponseProto")
		proto.RegisterType((*SetQuotaRequestProto)(nil), "hadoop.hdfs.SetQuotaRequestProto")
		proto.RegisterType((*SetQuotaResponseProto)(nil), "hadoop.hdfs.SetQuotaResponseProto")
		proto.RegisterType((*FsyncRequestProto)(nil), "hadoop.hdfs.FsyncRequestProto")
		proto.RegisterType((*FsyncResponseProto)(nil), "hadoop.hdfs.FsyncResponseProto")
		proto.RegisterType((*SetTimesRequestProto)(nil), "hadoop.hdfs.SetTimesRequestProto")
		proto.RegisterType((*SetTimesResponseProto)(nil), "hadoop.hdfs.SetTimesResponseProto")
		proto.RegisterType((*CreateSymlinkRequestProto)(nil), "hadoop.hdfs.CreateSymlinkRequestProto")
		proto.RegisterType((*CreateSymlinkResponseProto)(nil), "hadoop.hdfs.CreateSymlinkResponseProto")
		proto.RegisterType((*GetLinkTargetRequestProto)(nil), "hadoop.hdfs.GetLinkTargetRequestProto")
		proto.RegisterType((*GetLinkTargetResponseProto)(nil), "hadoop.hdfs.GetLinkTargetResponseProto")
		proto.RegisterType((*UpdateBlockForPipelineRequestProto)(nil), "hadoop.hdfs.UpdateBlockForPipelineRequestProto")
		proto.RegisterType((*UpdateBlockForPipelineResponseProto)(nil), "hadoop.hdfs.UpdateBlockForPipelineResponseProto")
		proto.RegisterType((*UpdatePipelineRequestProto)(nil), "hadoop.hdfs.UpdatePipelineRequestProto")
		proto.RegisterType((*UpdatePipelineResponseProto)(nil), "hadoop.hdfs.UpdatePipelineResponseProto")
		proto.RegisterType((*SetBalancerBandwidthRequestProto)(nil), "hadoop.hdfs.SetBalancerBandwidthRequestProto")
		proto.RegisterType((*SetBalancerBandwidthResponseProto)(nil), "hadoop.hdfs.SetBalancerBandwidthResponseProto")
		proto.RegisterType((*GetDataEncryptionKeyRequestProto)(nil), "hadoop.hdfs.GetDataEncryptionKeyRequestProto")
		proto.RegisterType((*GetDataEncryptionKeyResponseProto)(nil), "hadoop.hdfs.GetDataEncryptionKeyResponseProto")
		proto.RegisterType((*CreateSnapshotRequestProto)(nil), "hadoop.hdfs.CreateSnapshotRequestProto")
		proto.RegisterType((*CreateSnapshotResponseProto)(nil), "hadoop.hdfs.CreateSnapshotResponseProto")
		proto.RegisterType((*RenameSnapshotRequestProto)(nil), "hadoop.hdfs.RenameSnapshotRequestProto")
		proto.RegisterType((*RenameSnapshotResponseProto)(nil), "hadoop.hdfs.RenameSnapshotResponseProto")
		proto.RegisterType((*AllowSnapshotRequestProto)(nil), "hadoop.hdfs.AllowSnapshotRequestProto")
		proto.RegisterType((*AllowSnapshotResponseProto)(nil), "hadoop.hdfs.AllowSnapshotResponseProto")
		proto.RegisterType((*DisallowSnapshotRequestProto)(nil), "hadoop.hdfs.DisallowSnapshotRequestProto")
		proto.RegisterType((*DisallowSnapshotResponseProto)(nil), "hadoop.hdfs.DisallowSnapshotResponseProto")
		proto.RegisterType((*DeleteSnapshotRequestProto)(nil), "hadoop.hdfs.DeleteSnapshotRequestProto")
		proto.RegisterType((*DeleteSnapshotResponseProto)(nil), "hadoop.hdfs.DeleteSnapshotResponseProto")
		proto.RegisterType((*CheckAccessRequestProto)(nil), "hadoop.hdfs.CheckAccessRequestProto")
		proto.RegisterType((*CheckAccessResponseProto)(nil), "hadoop.hdfs.CheckAccessResponseProto")
		proto.RegisterType((*GetCurrentEditLogTxidRequestProto)(nil), "hadoop.hdfs.GetCurrentEditLogTxidRequestProto")
		proto.RegisterType((*GetCurrentEditLogTxidResponseProto)(nil), "hadoop.hdfs.GetCurrentEditLogTxidResponseProto")
		proto.RegisterType((*GetEditsFromTxidRequestProto)(nil), "hadoop.hdfs.GetEditsFromTxidRequestProto")
		proto.RegisterType((*GetEditsFromTxidResponseProto)(nil), "hadoop.hdfs.GetEditsFromTxidResponseProto")
		proto.RegisterEnum("hadoop.hdfs.CreateFlagProto", CreateFlagProto_name, CreateFlagProto_value)
		proto.RegisterEnum("hadoop.hdfs.DatanodeReportTypeProto", DatanodeReportTypeProto_name, DatanodeReportTypeProto_value)
		proto.RegisterEnum("hadoop.hdfs.SafeModeActionProto", SafeModeActionProto_name, SafeModeActionProto_value)
		proto.RegisterEnum("hadoop.hdfs.RollingUpgradeActionProto", RollingUpgradeActionProto_name, RollingUpgradeActionProto_value)
		proto.RegisterEnum("hadoop.hdfs.CacheFlagProto", CacheFlagProto_name, CacheFlagProto_value)
		proto.RegisterFile("ClientNamenodeProtocol.proto", fileDescriptor4)
		proto.RegisterType((*StartReconfigurationRequestProto)(nil), "hadoop.hdfs.StartReconfigurationRequestProto")
		proto.RegisterType((*StartReconfigurationResponseProto)(nil), "hadoop.hdfs.StartReconfigurationResponseProto")
		proto.RegisterType((*GetReconfigurationStatusRequestProto)(nil), "hadoop.hdfs.GetReconfigurationStatusRequestProto")
		proto.RegisterType((*GetReconfigurationStatusConfigChangeProto)(nil), "hadoop.hdfs.GetReconfigurationStatusConfigChangeProto")
		proto.RegisterType((*GetReconfigurationStatusResponseProto)(nil), "hadoop.hdfs.GetReconfigurationStatusResponseProto")
		proto.RegisterType((*ListReconfigurablePropertiesRequestProto)(nil), "hadoop.hdfs.ListReconfigurablePropertiesRequestProto")
		proto.RegisterType((*ListReconfigurablePropertiesResponseProto)(nil), "hadoop.hdfs.ListReconfigurablePropertiesResponseProto")
		proto.RegisterFile("ReconfigurationProtocol.proto", fileDescriptor0)
		proto.RegisterType((*AclEntryProto)(nil), "hadoop.hdfs.AclEntryProto")
		proto.RegisterType((*AclStatusProto)(nil), "hadoop.hdfs.AclStatusProto")
		proto.RegisterType((*ModifyAclEntriesRequestProto)(nil), "hadoop.hdfs.ModifyAclEntriesRequestProto")
		proto.RegisterType((*ModifyAclEntriesResponseProto)(nil), "hadoop.hdfs.ModifyAclEntriesResponseProto")
		proto.RegisterType((*RemoveAclRequestProto)(nil), "hadoop.hdfs.RemoveAclRequestProto")
		proto.RegisterType((*RemoveAclResponseProto)(nil), "hadoop.hdfs.RemoveAclResponseProto")
		proto.RegisterType((*RemoveAclEntriesRequestProto)(nil), "hadoop.hdfs.RemoveAclEntriesRequestProto")
		proto.RegisterType((*RemoveAclEntriesResponseProto)(nil), "hadoop.hdfs.RemoveAclEntriesResponseProto")
		proto.RegisterType((*RemoveDefaultAclRequestProto)(nil), "hadoop.hdfs.RemoveDefaultAclRequestProto")
		proto.RegisterType((*RemoveDefaultAclResponseProto)(nil), "hadoop.hdfs.RemoveDefaultAclResponseProto")
		proto.RegisterType((*SetAclRequestProto)(nil), "hadoop.hdfs.SetAclRequestProto")
		proto.RegisterType((*SetAclResponseProto)(nil), "hadoop.hdfs.SetAclResponseProto")
		proto.RegisterType((*GetAclStatusRequestProto)(nil), "hadoop.hdfs.GetAclStatusRequestProto")
		proto.RegisterType((*GetAclStatusResponseProto)(nil), "hadoop.hdfs.GetAclStatusResponseProto")
		proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryScopeProto", AclEntryProto_AclEntryScopeProto_name, AclEntryProto_AclEntryScopeProto_value)
		proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryTypeProto", AclEntryProto_AclEntryTypeProto_name, AclEntryProto_AclEntryTypeProto_value)
		proto.RegisterEnum("hadoop.hdfs.AclEntryProto_FsActionProto", AclEntryProto_FsActionProto_name, AclEntryProto_FsActionProto_value)
		proto.RegisterFile("acl.proto", fileDescriptor9)
		proto.RegisterType((*DataTransferEncryptorMessageProto)(nil), "hadoop.hdfs.DataTransferEncryptorMessageProto")
		proto.RegisterType((*BaseHeaderProto)(nil), "hadoop.hdfs.BaseHeaderProto")
		proto.RegisterType((*DataTransferTraceInfoProto)(nil), "hadoop.hdfs.DataTransferTraceInfoProto")
		proto.RegisterType((*ClientOperationHeaderProto)(nil), "hadoop.hdfs.ClientOperationHeaderProto")
		proto.RegisterType((*CachingStrategyProto)(nil), "hadoop.hdfs.CachingStrategyProto")
		proto.RegisterType((*OpReadBlockProto)(nil), "hadoop.hdfs.OpReadBlockProto")
		proto.RegisterType((*ChecksumProto)(nil), "hadoop.hdfs.ChecksumProto")
		proto.RegisterType((*OpWriteBlockProto)(nil), "hadoop.hdfs.OpWriteBlockProto")
		proto.RegisterType((*OpTransferBlockProto)(nil), "hadoop.hdfs.OpTransferBlockProto")
		proto.RegisterType((*OpReplaceBlockProto)(nil), "hadoop.hdfs.OpReplaceBlockProto")
		proto.RegisterType((*OpCopyBlockProto)(nil), "hadoop.hdfs.OpCopyBlockProto")
		proto.RegisterType((*OpBlockChecksumProto)(nil), "hadoop.hdfs.OpBlockChecksumProto")
		proto.RegisterType((*OpBlockGroupChecksumProto)(nil), "hadoop.hdfs.OpBlockGroupChecksumProto")
		proto.RegisterType((*ShortCircuitShmIdProto)(nil), "hadoop.hdfs.ShortCircuitShmIdProto")
		proto.RegisterType((*ShortCircuitShmSlotProto)(nil), "hadoop.hdfs.ShortCircuitShmSlotProto")
		proto.RegisterType((*OpRequestShortCircuitAccessProto)(nil), "hadoop.hdfs.OpRequestShortCircuitAccessProto")
		proto.RegisterType((*ReleaseShortCircuitAccessRequestProto)(nil), "hadoop.hdfs.ReleaseShortCircuitAccessRequestProto")
		proto.RegisterType((*ReleaseShortCircuitAccessResponseProto)(nil), "hadoop.hdfs.ReleaseShortCircuitAccessResponseProto")
		proto.RegisterType((*ShortCircuitShmRequestProto)(nil), "hadoop.hdfs.ShortCircuitShmRequestProto")
		proto.RegisterType((*ShortCircuitShmResponseProto)(nil), "hadoop.hdfs.ShortCircuitShmResponseProto")
		proto.RegisterType((*PacketHeaderProto)(nil), "hadoop.hdfs.PacketHeaderProto")
		proto.RegisterType((*PipelineAckProto)(nil), "hadoop.hdfs.PipelineAckProto")
		proto.RegisterType((*ReadOpChecksumInfoProto)(nil), "hadoop.hdfs.ReadOpChecksumInfoProto")
		proto.RegisterType((*BlockOpResponseProto)(nil), "hadoop.hdfs.BlockOpResponseProto")
		proto.RegisterType((*ClientReadStatusProto)(nil), "hadoop.hdfs.ClientReadStatusProto")
		proto.RegisterType((*DNTransferAckProto)(nil), "hadoop.hdfs.DNTransferAckProto")
		proto.RegisterType((*OpBlockChecksumResponseProto)(nil), "hadoop.hdfs.OpBlockChecksumResponseProto")
		proto.RegisterType((*OpCustomProto)(nil), "hadoop.hdfs.OpCustomProto")
		proto.RegisterEnum("hadoop.hdfs.Status", Status_name, Status_value)
		proto.RegisterEnum("hadoop.hdfs.ShortCircuitFdResponse", ShortCircuitFdResponse_name, ShortCircuitFdResponse_value)
		proto.RegisterEnum("hadoop.hdfs.DataTransferEncryptorMessageProto_DataTransferEncryptorStatus", DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_name, DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_value)
		proto.RegisterEnum("hadoop.hdfs.OpWriteBlockProto_BlockConstructionStage", OpWriteBlockProto_BlockConstructionStage_name, OpWriteBlockProto_BlockConstructionStage_value)
		proto.RegisterFile("datatransfer.proto", fileDescriptor5)
		proto.RegisterType((*CreateEncryptionZoneRequestProto)(nil), "hadoop.hdfs.CreateEncryptionZoneRequestProto")
		proto.RegisterType((*CreateEncryptionZoneResponseProto)(nil), "hadoop.hdfs.CreateEncryptionZoneResponseProto")
		proto.RegisterType((*ListEncryptionZonesRequestProto)(nil), "hadoop.hdfs.ListEncryptionZonesRequestProto")
		proto.RegisterType((*EncryptionZoneProto)(nil), "hadoop.hdfs.EncryptionZoneProto")
		proto.RegisterType((*ListEncryptionZonesResponseProto)(nil), "hadoop.hdfs.ListEncryptionZonesResponseProto")
		proto.RegisterType((*GetEZForPathRequestProto)(nil), "hadoop.hdfs.GetEZForPathRequestProto")
		proto.RegisterType((*GetEZForPathResponseProto)(nil), "hadoop.hdfs.GetEZForPathResponseProto")
		proto.RegisterFile("encryption.proto", fileDescriptor2)
		proto.RegisterType((*SetErasureCodingPolicyRequestProto)(nil), "hadoop.hdfs.SetErasureCodingPolicyRequestProto")
		proto.RegisterType((*SetErasureCodingPolicyResponseProto)(nil), "hadoop.hdfs.SetErasureCodingPolicyResponseProto")
		proto.RegisterType((*GetErasureCodingPoliciesRequestProto)(nil), "hadoop.hdfs.GetErasureCodingPoliciesRequestProto")
		proto.RegisterType((*GetErasureCodingPoliciesResponseProto)(nil), "hadoop.hdfs.GetErasureCodingPoliciesResponseProto")
		proto.RegisterType((*GetErasureCodingPolicyRequestProto)(nil), "hadoop.hdfs.GetErasureCodingPolicyRequestProto")
		proto.RegisterType((*GetErasureCodingPolicyResponseProto)(nil), "hadoop.hdfs.GetErasureCodingPolicyResponseProto")
		proto.RegisterType((*BlockECReconstructionInfoProto)(nil), "hadoop.hdfs.BlockECReconstructionInfoProto")
		proto.RegisterFile("erasurecoding.proto", fileDescriptor3)
		proto.RegisterType((*ExtendedBlockProto)(nil), "hadoop.hdfs.ExtendedBlockProto")
		proto.RegisterType((*DatanodeIDProto)(nil), "hadoop.hdfs.DatanodeIDProto")
		proto.RegisterType((*DatanodeLocalInfoProto)(nil), "hadoop.hdfs.DatanodeLocalInfoProto")
		proto.RegisterType((*DatanodeInfosProto)(nil), "hadoop.hdfs.DatanodeInfosProto")
		proto.RegisterType((*DatanodeInfoProto)(nil), "hadoop.hdfs.DatanodeInfoProto")
		proto.RegisterType((*DatanodeStorageProto)(nil), "hadoop.hdfs.DatanodeStorageProto")
		proto.RegisterType((*StorageReportProto)(nil), "hadoop.hdfs.StorageReportProto")
		proto.RegisterType((*ContentSummaryProto)(nil), "hadoop.hdfs.ContentSummaryProto")
		proto.RegisterType((*QuotaUsageProto)(nil), "hadoop.hdfs.QuotaUsageProto")
		proto.RegisterType((*StorageTypeQuotaInfosProto)(nil), "hadoop.hdfs.StorageTypeQuotaInfosProto")
		proto.RegisterType((*StorageTypeQuotaInfoProto)(nil), "hadoop.hdfs.StorageTypeQuotaInfoProto")
		proto.RegisterType((*CorruptFileBlocksProto)(nil), "hadoop.hdfs.CorruptFileBlocksProto")
		proto.RegisterType((*FsPermissionProto)(nil), "hadoop.hdfs.FsPermissionProto")
		proto.RegisterType((*StorageTypesProto)(nil), "hadoop.hdfs.StorageTypesProto")
		proto.RegisterType((*BlockStoragePolicyProto)(nil), "hadoop.hdfs.BlockStoragePolicyProto")
		proto.RegisterType((*LocatedBlockProto)(nil), "hadoop.hdfs.LocatedBlockProto")
		proto.RegisterType((*DataEncryptionKeyProto)(nil), "hadoop.hdfs.DataEncryptionKeyProto")
		proto.RegisterType((*FileEncryptionInfoProto)(nil), "hadoop.hdfs.FileEncryptionInfoProto")
		proto.RegisterType((*PerFileEncryptionInfoProto)(nil), "hadoop.hdfs.PerFileEncryptionInfoProto")
		proto.RegisterType((*ZoneEncryptionInfoProto)(nil), "hadoop.hdfs.ZoneEncryptionInfoProto")
		proto.RegisterType((*CipherOptionProto)(nil), "hadoop.hdfs.CipherOptionProto")
		proto.RegisterType((*LocatedBlocksProto)(nil), "hadoop.hdfs.LocatedBlocksProto")
		proto.RegisterType((*ECSchemaOptionEntryProto)(nil), "hadoop.hdfs.ECSchemaOptionEntryProto")
		proto.RegisterType((*ECSchemaProto)(nil), "hadoop.hdfs.ECSchemaProto")
		proto.RegisterType((*ErasureCodingPolicyProto)(nil), "hadoop.hdfs.ErasureCodingPolicyProto")
		proto.RegisterType((*HdfsFileStatusProto)(nil), "hadoop.hdfs.HdfsFileStatusProto")
		proto.RegisterType((*FsServerDefaultsProto)(nil), "hadoop.hdfs.FsServerDefaultsProto")
		proto.RegisterType((*DirectoryListingProto)(nil), "hadoop.hdfs.DirectoryListingProto")
		proto.RegisterType((*SnapshottableDirectoryStatusProto)(nil), "hadoop.hdfs.SnapshottableDirectoryStatusProto")
		proto.RegisterType((*SnapshottableDirectoryListingProto)(nil), "hadoop.hdfs.SnapshottableDirectoryListingProto")
		proto.RegisterType((*SnapshotDiffReportEntryProto)(nil), "hadoop.hdfs.SnapshotDiffReportEntryProto")
		proto.RegisterType((*SnapshotDiffReportProto)(nil), "hadoop.hdfs.SnapshotDiffReportProto")
		proto.RegisterType((*BlockProto)(nil), "hadoop.hdfs.BlockProto")
		proto.RegisterType((*SnapshotInfoProto)(nil), "hadoop.hdfs.SnapshotInfoProto")
		proto.RegisterType((*RollingUpgradeStatusProto)(nil), "hadoop.hdfs.RollingUpgradeStatusProto")
		proto.RegisterType((*StorageUuidsProto)(nil), "hadoop.hdfs.StorageUuidsProto")
		proto.RegisterEnum("hadoop.hdfs.StorageTypeProto", StorageTypeProto_name, StorageTypeProto_value)
		proto.RegisterEnum("hadoop.hdfs.CipherSuiteProto", CipherSuiteProto_name, CipherSuiteProto_value)
		proto.RegisterEnum("hadoop.hdfs.CryptoProtocolVersionProto", CryptoProtocolVersionProto_name, CryptoProtocolVersionProto_value)
		proto.RegisterEnum("hadoop.hdfs.ChecksumTypeProto", ChecksumTypeProto_name, ChecksumTypeProto_value)
		proto.RegisterEnum("hadoop.hdfs.DatanodeInfoProto_AdminState", DatanodeInfoProto_AdminState_name, DatanodeInfoProto_AdminState_value)
		proto.RegisterEnum("hadoop.hdfs.DatanodeStorageProto_StorageState", DatanodeStorageProto_StorageState_name, DatanodeStorageProto_StorageState_value)
		proto.RegisterEnum("hadoop.hdfs.HdfsFileStatusProto_FileType", HdfsFileStatusProto_FileType_name, HdfsFileStatusProto_FileType_value)
		proto.RegisterFile("hdfs.proto", fileDescriptor8)
		proto.RegisterType((*EventProto)(nil), "hadoop.hdfs.EventProto")
		proto.RegisterType((*EventBatchProto)(nil), "hadoop.hdfs.EventBatchProto")
		proto.RegisterType((*CreateEventProto)(nil), "hadoop.hdfs.CreateEventProto")
		proto.RegisterType((*CloseEventProto)(nil), "hadoop.hdfs.CloseEventProto")
		proto.RegisterType((*TruncateEventProto)(nil), "hadoop.hdfs.TruncateEventProto")
		proto.RegisterType((*AppendEventProto)(nil), "hadoop.hdfs.AppendEventProto")
		proto.RegisterType((*RenameEventProto)(nil), "hadoop.hdfs.RenameEventProto")
		proto.RegisterType((*MetadataUpdateEventProto)(nil), "hadoop.hdfs.MetadataUpdateEventProto")
		proto.RegisterType((*UnlinkEventProto)(nil), "hadoop.hdfs.UnlinkEventProto")
		proto.RegisterType((*EventsListProto)(nil), "hadoop.hdfs.EventsListProto")
		proto.RegisterEnum("hadoop.hdfs.EventType", EventType_name, EventType_value)
		proto.RegisterEnum("hadoop.hdfs.INodeType", INodeType_name, INodeType_value)
		proto.RegisterEnum("hadoop.hdfs.MetadataUpdateType", MetadataUpdateType_name, MetadataUpdateType_value)
		proto.RegisterFile("inotify.proto", fileDescriptor7)
		proto.RegisterType((*XAttrProto)(nil), "hadoop.hdfs.XAttrProto")
		proto.RegisterType((*SetXAttrRequestProto)(nil), "hadoop.hdfs.SetXAttrRequestProto")
		proto.RegisterType((*SetXAttrResponseProto)(nil), "hadoop.hdfs.SetXAttrResponseProto")
		proto.RegisterType((*GetXAttrsRequestProto)(nil), "hadoop.hdfs.GetXAttrsRequestProto")
		proto.RegisterType((*GetXAttrsResponseProto)(nil), "hadoop.hdfs.GetXAttrsResponseProto")
		proto.RegisterType((*ListXAttrsRequestProto)(nil), "hadoop.hdfs.ListXAttrsRequestProto")
		proto.RegisterType((*ListXAttrsResponseProto)(nil), "hadoop.hdfs.ListXAttrsResponseProto")
		proto.RegisterType((*RemoveXAttrRequestProto)(nil), "hadoop.hdfs.RemoveXAttrRequestProto")
		proto.RegisterType((*RemoveXAttrResponseProto)(nil), "hadoop.hdfs.RemoveXAttrResponseProto")
		proto.RegisterEnum("hadoop.hdfs.XAttrSetFlagProto", XAttrSetFlagProto_name, XAttrSetFlagProto_value)
		proto.RegisterEnum("hadoop.hdfs.XAttrProto_XAttrNamespaceProto", XAttrProto_XAttrNamespaceProto_name, XAttrProto_XAttrNamespaceProto_value)
		proto.RegisterFile("xattr.proto", fileDescriptor1)
	})
}
