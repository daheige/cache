package cache

// Cache cache interface
type Cache interface {
	Set(key string, entry []byte) error
	Get(key string) ([]byte, error)

	SetJson(key string, value interface{}) error
	GetJson(key string, bean interface{}) error

	Delete(key string) error
	Reset() error // clear all entries
	Len() int

	// Capacity returns amount of bytes store in the cache.
	Capacity() int
	Stats() Stats
	KeyMetadata(key string) Metadata
	GetWithInfo(key string) ([]byte, Response, error)
	Close() error
}

// RemoveReason is a value used to signal to the user why a particular key was removed in the OnRemove callback.
type RemoveReason uint32

const (
	// Expired means the key is past its LifeWindow.
	Expired RemoveReason = iota + 1
	// NoSpace means the key is the oldest and the cache size was at its maximum when Set was called, or the
	// entry exceeded the maximum shard size.
	NoSpace
	// Deleted means Delete was called and this key was removed as a result.
	Deleted
)

// Response will contain metadata about the entry for which GetWithInfo(key) was called
type Response struct {
	EntryStatus RemoveReason
}

// Metadata contains information of a specific entry
type Metadata struct {
	RequestCount uint32
}

// Stats stores cache statistics
type Stats struct {
	// Hits is a number of successfully found keys
	Hits int64 `json:"hits"`
	// Misses is a number of not found keys
	Misses int64 `json:"misses"`
	// DelHits is a number of successfully deleted keys
	DelHits int64 `json:"delete_hits"`
	// DelMisses is a number of not deleted keys
	DelMisses int64 `json:"delete_misses"`
	// Collisions is a number of happened key-collisions
	Collisions int64 `json:"collisions"`
}

// bytes to bit
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)
