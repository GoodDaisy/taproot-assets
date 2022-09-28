package address

import (
	"context"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightninglabs/lndclient"
)

// Status denotes an address event's current status.
type Status uint8

const (
	// StatusTransactionDetected denotes that a transaction for an incoming
	// asset transfer was detected but the transaction hasn't been confirmed
	// yet.
	StatusTransactionDetected Status = 0

	// StatusTransactionConfirmed denotes that the transaction for an
	// incoming asset transfer was confirmed. The transfer now requires the
	// proof to be imported to proceed.
	StatusTransactionConfirmed Status = 1

	// StatusProofReceived denotes that the proof for an incoming asset
	// transfer was received and is now being validated and processed.
	StatusProofReceived Status = 2

	// StatusCompleted denotes that an incoming asset transfer was completed
	// successfully and the local node has taken over custody of the assets
	// that were transferred.
	StatusCompleted Status = 3
)

// EventQueryParams holds the set of query params for address events.
type EventQueryParams struct {
	// AddrTaprootOutputKey is the optional 32-byte x-only serialized
	// Taproot output key of the address to filter by. Must be set to nil
	// to return events for all addresses.
	AddrTaprootOutputKey []byte

	// StatusFrom is the smallest status to query for (inclusive). Can be
	// set to nil to return events of all states.
	StatusFrom *Status

	// StatusTo is the largest status to query for (inclusive). Can be
	// set to nil to return events of all states.
	StatusTo *Status
}

// Event represents a single incoming asset transfer that was initiated by
// sending an on-chain transaction to the Taproot output key generated by a Taro
// address. Each event represents a single on-chain UTXO that is being taken
// custody of and is being tracked/watched by the internal wallet. One Taro
// address can receive multiple times and therefore can have multiple events.
type Event struct {
	// ID is the database primary key ID of the address event.
	ID int32

	// CreationTime is the time the event was first created.
	CreationTime time.Time

	// Addr is the Taro address that was used to receive the assets.
	Addr *AddrWithKeyInfo

	// Status represents the current status of the incoming assets.
	Status Status

	// Outpoint is the on-chain transaction outpoint that contains the Taro
	// commitment for the incoming asset transfer.
	Outpoint wire.OutPoint

	// Amt is the amount of satoshis that were transferred in the Bitcoin
	// on-chain transaction. This is independent of the asset amount, which
	// can be looked up through the Addr field.
	Amt btcutil.Amount

	// InternalKey is the key used as the internal key for the on-chain
	// Taproot output. The internal key tweaked with the Taro commitment
	// (when NO tapscript sibling if present) is equal to the
	// TaprootOutputKey of the Addr.
	InternalKey *btcec.PublicKey

	// TapscriptSibling, if non-empty, signals that a Tapscript sibling leaf
	// was present in the on-chain Taproot output that sent the assets. This
	// can only be achieved by importing a proof and then scanning for the
	// Taproot output on chain retroactively, since at the time a Taro
	// address is created the sibling might not yet be known.
	//
	// NOTE: The functionality described above (importing the proof and re-
	// scanning the chain) is not yet implemented, so this should always be
	// empty or nil at the moment.
	TapscriptSibling []byte

	// ConfirmationHeight is the block height at which the incoming asset
	// transfer transaction was first confirmed.
	ConfirmationHeight uint32

	// HasProof indicates that a proof for this transfer was imported. We
	// don't keep a reference to it in memory as the proof itself can be
	// large. The proof can be fetched by the script key of the address.
	HasProof bool
}

// EventStorage is the interface that a component storing address events should
// implement.
type EventStorage interface {
	// GetOrCreateEvent creates a new address event for the given status,
	// address and transaction. If an event for that address and transaction
	// already exists, then the status and transaction information is
	// updated instead.
	GetOrCreateEvent(ctx context.Context, status Status,
		addr *AddrWithKeyInfo, walletTx *lndclient.Transaction,
		outputIdx uint32, tapscriptSibling *chainhash.Hash) (*Event,
		error)

	// QueryAddrEvents returns a list of event that match the given query
	// parameters.
	QueryAddrEvents(ctx context.Context, params EventQueryParams) ([]*Event,
		error)

	// CompleteEvent updates an address event as being complete and links it
	// with the proof and asset that was imported/created for it.
	CompleteEvent(ctx context.Context, event *Event, status Status,
		anchorPoint wire.OutPoint) error
}