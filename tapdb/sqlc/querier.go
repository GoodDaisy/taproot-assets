// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

type Querier interface {
	AllAssets(ctx context.Context) ([]Asset, error)
	AllInternalKeys(ctx context.Context) ([]InternalKey, error)
	AllMintingBatches(ctx context.Context) ([]AllMintingBatchesRow, error)
	AnchorGenesisPoint(ctx context.Context, arg AnchorGenesisPointParams) error
	AnchorPendingAssets(ctx context.Context, arg AnchorPendingAssetsParams) error
	ApplyPendingOutput(ctx context.Context, arg ApplyPendingOutputParams) (int32, error)
	AssetsByGenesisPoint(ctx context.Context, prevOut []byte) ([]AssetsByGenesisPointRow, error)
	AssetsInBatch(ctx context.Context, rawKey []byte) ([]AssetsInBatchRow, error)
	BindMintingBatchWithTx(ctx context.Context, arg BindMintingBatchWithTxParams) error
	ConfirmChainAnchorTx(ctx context.Context, arg ConfirmChainAnchorTxParams) error
	ConfirmChainTx(ctx context.Context, arg ConfirmChainTxParams) error
	DeleteAllNodes(ctx context.Context, namespace string) (int64, error)
	DeleteAssetWitnesses(ctx context.Context, assetID int32) error
	DeleteManagedUTXO(ctx context.Context, outpoint []byte) error
	DeleteNode(ctx context.Context, arg DeleteNodeParams) (int64, error)
	DeleteRoot(ctx context.Context, namespace string) (int64, error)
	DeleteUniverseEvents(ctx context.Context, namespaceRoot string) error
	DeleteUniverseLeaves(ctx context.Context, namespace string) error
	DeleteUniverseRoot(ctx context.Context, namespaceRoot string) error
	DeleteUniverseServer(ctx context.Context, arg DeleteUniverseServerParams) error
	FetchAddrByTaprootOutputKey(ctx context.Context, taprootOutputKey []byte) (FetchAddrByTaprootOutputKeyRow, error)
	FetchAddrEvent(ctx context.Context, id int32) (FetchAddrEventRow, error)
	FetchAddrs(ctx context.Context, arg FetchAddrsParams) ([]FetchAddrsRow, error)
	FetchAllNodes(ctx context.Context) ([]MssmtNode, error)
	FetchAssetMeta(ctx context.Context, metaID int32) (FetchAssetMetaRow, error)
	FetchAssetMetaByHash(ctx context.Context, metaDataHash []byte) (FetchAssetMetaByHashRow, error)
	FetchAssetMetaForAsset(ctx context.Context, assetID []byte) (FetchAssetMetaForAssetRow, error)
	FetchAssetProof(ctx context.Context, tweakedScriptKey []byte) (FetchAssetProofRow, error)
	FetchAssetProofs(ctx context.Context) ([]FetchAssetProofsRow, error)
	FetchAssetWitnesses(ctx context.Context, assetID sql.NullInt32) ([]FetchAssetWitnessesRow, error)
	FetchAssetsByAnchorTx(ctx context.Context, anchorUtxoID sql.NullInt32) ([]Asset, error)
	// We use a LEFT JOIN here as not every asset has a group key, so this'll
	// generate rows that have NULL values for the faily key fields if an asset
	// doesn't have a group key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	FetchAssetsForBatch(ctx context.Context, rawKey []byte) ([]FetchAssetsForBatchRow, error)
	FetchChainTx(ctx context.Context, txid []byte) (ChainTxn, error)
	FetchChildren(ctx context.Context, arg FetchChildrenParams) ([]FetchChildrenRow, error)
	FetchChildrenSelfJoin(ctx context.Context, arg FetchChildrenSelfJoinParams) ([]FetchChildrenSelfJoinRow, error)
	FetchGenesisByAssetID(ctx context.Context, assetID []byte) (GenesisInfoView, error)
	FetchGenesisByID(ctx context.Context, genAssetID int32) (FetchGenesisByIDRow, error)
	FetchGenesisID(ctx context.Context, arg FetchGenesisIDParams) (int32, error)
	FetchGenesisPointByAnchorTx(ctx context.Context, anchorTxID sql.NullInt32) (GenesisPoint, error)
	FetchGroupByGenesis(ctx context.Context, genesisID int32) (FetchGroupByGenesisRow, error)
	// Sort and limit to return the genesis ID for initial genesis of the group.
	FetchGroupByGroupKey(ctx context.Context, groupKey []byte) (FetchGroupByGroupKeyRow, error)
	FetchGroupedAssets(ctx context.Context) ([]FetchGroupedAssetsRow, error)
	FetchManagedUTXO(ctx context.Context, arg FetchManagedUTXOParams) (FetchManagedUTXORow, error)
	FetchManagedUTXOs(ctx context.Context) ([]FetchManagedUTXOsRow, error)
	FetchMintingBatch(ctx context.Context, rawKey []byte) (FetchMintingBatchRow, error)
	FetchMintingBatchesByInverseState(ctx context.Context, batchState int16) ([]FetchMintingBatchesByInverseStateRow, error)
	FetchRootNode(ctx context.Context, namespace string) (MssmtNode, error)
	FetchScriptKeyByTweakedKey(ctx context.Context, tweakedScriptKey []byte) (FetchScriptKeyByTweakedKeyRow, error)
	FetchScriptKeyIDByTweakedKey(ctx context.Context, tweakedScriptKey []byte) (int32, error)
	FetchSeedlingByID(ctx context.Context, seedlingID int32) (AssetSeedling, error)
	FetchSeedlingID(ctx context.Context, arg FetchSeedlingIDParams) (int32, error)
	FetchSeedlingsForBatch(ctx context.Context, rawKey []byte) ([]FetchSeedlingsForBatchRow, error)
	FetchTransferInputs(ctx context.Context, transferID int32) ([]FetchTransferInputsRow, error)
	FetchTransferOutputs(ctx context.Context, transferID int32) ([]FetchTransferOutputsRow, error)
	FetchUniverseKeys(ctx context.Context, namespace string) ([]FetchUniverseKeysRow, error)
	FetchUniverseRoot(ctx context.Context, namespace string) (FetchUniverseRootRow, error)
	GenesisAssets(ctx context.Context) ([]GenesisAsset, error)
	GenesisPoints(ctx context.Context) ([]GenesisPoint, error)
	GetRootKey(ctx context.Context, id []byte) (Macaroon, error)
	InsertAddr(ctx context.Context, arg InsertAddrParams) (int32, error)
	InsertAssetSeedling(ctx context.Context, arg InsertAssetSeedlingParams) error
	InsertAssetSeedlingIntoBatch(ctx context.Context, arg InsertAssetSeedlingIntoBatchParams) error
	InsertAssetTransfer(ctx context.Context, arg InsertAssetTransferParams) (int32, error)
	InsertAssetTransferInput(ctx context.Context, arg InsertAssetTransferInputParams) error
	InsertAssetTransferOutput(ctx context.Context, arg InsertAssetTransferOutputParams) error
	InsertAssetWitness(ctx context.Context, arg InsertAssetWitnessParams) error
	InsertBranch(ctx context.Context, arg InsertBranchParams) error
	InsertCompactedLeaf(ctx context.Context, arg InsertCompactedLeafParams) error
	InsertLeaf(ctx context.Context, arg InsertLeafParams) error
	InsertNewAsset(ctx context.Context, arg InsertNewAssetParams) (int32, error)
	InsertNewProofEvent(ctx context.Context, arg InsertNewProofEventParams) error
	InsertNewSyncEvent(ctx context.Context, arg InsertNewSyncEventParams) error
	InsertPassiveAsset(ctx context.Context, arg InsertPassiveAssetParams) error
	InsertReceiverProofTransferAttempt(ctx context.Context, arg InsertReceiverProofTransferAttemptParams) error
	InsertRootKey(ctx context.Context, arg InsertRootKeyParams) error
	InsertUniverseLeaf(ctx context.Context, arg InsertUniverseLeafParams) error
	InsertUniverseServer(ctx context.Context, arg InsertUniverseServerParams) error
	ListUniverseServers(ctx context.Context) ([]UniverseServer, error)
	LogServerSync(ctx context.Context, arg LogServerSyncParams) error
	NewMintingBatch(ctx context.Context, arg NewMintingBatchParams) error
	// We use a LEFT JOIN here as not every asset has a group key, so this'll
	// generate rows that have NULL values for the group key fields if an asset
	// doesn't have a group key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	QueryAssetBalancesByAsset(ctx context.Context, assetIDFilter []byte) ([]QueryAssetBalancesByAssetRow, error)
	QueryAssetBalancesByGroup(ctx context.Context, keyGroupFilter []byte) ([]QueryAssetBalancesByGroupRow, error)
	// We'll use this clause to filter out for only transfers that are
	// unconfirmed. But only if the unconf_only field is set.
	// Here we have another optional query clause to select a given transfer
	// based on the anchor_tx_hash, but only if it's specified.
	QueryAssetTransfers(ctx context.Context, arg QueryAssetTransfersParams) ([]QueryAssetTransfersRow, error)
	// We use a LEFT JOIN here as not every asset has a group key, so this'll
	// generate rows that have NULL values for the group key fields if an asset
	// doesn't have a group key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	// This clause is used to select specific assets for a asset ID, general
	// channel balances, and also coin selection. We use the sqlc.narg feature to
	// make the entire statement evaluate to true, if none of these extra args are
	// specified.
	QueryAssets(ctx context.Context, arg QueryAssetsParams) ([]QueryAssetsRow, error)
	QueryEventIDs(ctx context.Context, arg QueryEventIDsParams) ([]QueryEventIDsRow, error)
	QueryPassiveAssets(ctx context.Context, transferID int32) ([]QueryPassiveAssetsRow, error)
	QueryReceiverProofTransferAttempt(ctx context.Context, proofLocatorHash []byte) ([]time.Time, error)
	// TODO(roasbeef): use the universe id instead for the grouping? so namespace
	// root, simplifies queries
	QueryUniverseAssetStats(ctx context.Context, arg QueryUniverseAssetStatsParams) ([]QueryUniverseAssetStatsRow, error)
	QueryUniverseLeaves(ctx context.Context, arg QueryUniverseLeavesParams) ([]QueryUniverseLeavesRow, error)
	QueryUniverseStats(ctx context.Context) (QueryUniverseStatsRow, error)
	ReAnchorPassiveAssets(ctx context.Context, arg ReAnchorPassiveAssetsParams) error
	SetAddrManaged(ctx context.Context, arg SetAddrManagedParams) error
	SetAssetSpent(ctx context.Context, arg SetAssetSpentParams) (int32, error)
	UniverseLeaves(ctx context.Context) ([]UniverseLeafe, error)
	UniverseRoots(ctx context.Context) ([]UniverseRootsRow, error)
	UpdateBatchGenesisTx(ctx context.Context, arg UpdateBatchGenesisTxParams) error
	UpdateMintingBatchState(ctx context.Context, arg UpdateMintingBatchStateParams) error
	UpsertAddrEvent(ctx context.Context, arg UpsertAddrEventParams) (int32, error)
	UpsertAssetGroupKey(ctx context.Context, arg UpsertAssetGroupKeyParams) (int32, error)
	UpsertAssetGroupSig(ctx context.Context, arg UpsertAssetGroupSigParams) (int32, error)
	UpsertAssetMeta(ctx context.Context, arg UpsertAssetMetaParams) (int32, error)
	UpsertAssetProof(ctx context.Context, arg UpsertAssetProofParams) error
	UpsertChainTx(ctx context.Context, arg UpsertChainTxParams) (int32, error)
	UpsertGenesisAsset(ctx context.Context, arg UpsertGenesisAssetParams) (int32, error)
	UpsertGenesisPoint(ctx context.Context, prevOut []byte) (int32, error)
	UpsertInternalKey(ctx context.Context, arg UpsertInternalKeyParams) (int32, error)
	UpsertManagedUTXO(ctx context.Context, arg UpsertManagedUTXOParams) (int32, error)
	UpsertRootNode(ctx context.Context, arg UpsertRootNodeParams) error
	UpsertScriptKey(ctx context.Context, arg UpsertScriptKeyParams) (int32, error)
	UpsertUniverseRoot(ctx context.Context, arg UpsertUniverseRootParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
