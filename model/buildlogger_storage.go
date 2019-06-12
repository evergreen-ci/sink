package model

import (
	"github.com/mongodb/anser/bsonutil"
	"github.com/pkg/errors"
)

// LogFormat is a type that describes the format of a log.
type LogFormat string

// Valid log formats.
const (
	LogFormatUnknown LogFormat = "unknown"
	LogFormatText    LogFormat = "text"
	LogFormatJSON    LogFormat = "json"
	LogFormatBSON    LogFormat = "bson"
)

// Validate the log format.
func (lf LogFormat) Validate() error {
	switch lf {
	case LogFormatUnknown, LogFormatText, LogFormatJSON, LogFormatBSON:
		return nil
	default:
		return errors.New("invalid log format type specified")
	}
}

// LogArtifact describes a bucket of logs stored in some kind of offline blob
// storage. It is the bridge between pail-backed offline log storage and the
// cedar-based log metadata storage. The prefix field indicates the name of the
// "sub-bucket". The top level bucket is accesible via the cedar.Environment
// interface.
type LogArtifactInfo struct {
	Type        PailType       `bson:"type"`
	Prefix      string         `bson:"prefix"`
	Permissions string         `bson:"permissions"`
	Version     int            `bson:"version"`
	Chunks      []LogChunkInfo `bson:"chunks,omitempty"`
}

var (
	logArtifactInfoTypeKey        = bsonutil.MustHaveTag(LogArtifactInfo{}, "Type")
	logArtifactInfoPrefixKey      = bsonutil.MustHaveTag(LogArtifactInfo{}, "Prefix")
	logArtifactInfoPermissionsKey = bsonutil.MustHaveTag(LogArtifactInfo{}, "Permissions")
	logArtifactInfoVersionKey     = bsonutil.MustHaveTag(LogArtifactInfo{}, "Version")
	logArtifactInfoChunksKey      = bsonutil.MustHaveTag(LogArtifactInfo{}, "Chunks")
)

// LogChunkInfo describes a chunk of log lines stored in pail-backed offline
// storage.
type LogChunkInfo struct {
	Key      string `bson:"key"`
	NumLines int    `bson:"num_lines"`
	Start    int64  `bson:"start"`
	End      int64  `bson:"end"`
}

var (
	logLogChunkInfoKeyKey      = bsonutil.MustHaveTag(LogChunkInfo{}, "Key")
	logLogChunkInfoNumLinesKey = bsonutil.MustHaveTag(LogChunkInfo{}, "NumLines")
	logLogChunkInfoStartKey    = bsonutil.MustHaveTag(LogChunkInfo{}, "Start")
	logLogChunkInfoEndKey      = bsonutil.MustHaveTag(LogChunkInfo{}, "End")
)
