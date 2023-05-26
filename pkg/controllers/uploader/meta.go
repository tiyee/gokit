package uploader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tiyee/gokit/pkg/component/cache"

	"math"
	"sort"
	"strconv"
	"time"
)

const (
	InitializeStatus int8 = iota
	PendingStatus
	FulfilledStatus
	RejectStatus
)
const (
	OSS int8 = 1
	COS int8 = 2
)

type Uploader struct {
	RequestID string
	Object    string
	Size      int
	Host      string
	c         context.Context
}
type Meta struct {
	ctx         context.Context
	UploadID    string
	Object      string
	Size        int
	ChunkSize   int
	Host        string
	Digest      string
	Status      int8
	ContentType string
	uploaderSrc int8
	chunks      []*ChunkETag
	uploader    IUploader
	lastModify  int64
	cache       cache.ICache
}
type MetaCacheValue struct {
	UploadID   string       `json:"u"`
	Object     string       `json:"ob"`
	Size       int          `json:"s"`
	ChunkSize  int          `json:"c"`
	Chunks     []*ChunkETag `json:"ch"`
	Digest     string       `json:"d"`
	Status     int8         `json:"st"`
	LastModify int64        `json:"l"`
	Uploader   int8         `json:"ul"`
}
type OptionFunc func(m *Meta)
type ChunkETag struct {
	Index int    `json:"index"`
	ETag  string `json:"etag"`
}
type ChunksETag []ChunkETag

func (c ChunksETag) Len() int {
	return len(c)
}

func (c ChunksETag) Less(i, j int) bool {
	return c[i].Index < c[j].Index
}

func (c ChunksETag) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func FromUploadID(ctx context.Context, cache cache.ICache, uploadID string, opts ...OptionFunc) (*Meta, error) {
	meta := &Meta{
		ctx:      ctx,
		Status:   InitializeStatus,
		UploadID: uploadID,
		cache:    cache,
	}
	for _, opt := range opts {
		opt(meta)
	}

	if bs, err := meta.cache.Get(uploadID); err == nil {
		var mcv MetaCacheValue
		if err := json.Unmarshal(bs, &mcv); err == nil {
			meta.Object = mcv.Object
			meta.Size = mcv.Size
			meta.ChunkSize = mcv.ChunkSize
			meta.Digest = mcv.Digest
			meta.Status = mcv.Status
			meta.chunks = mcv.Chunks
			meta.lastModify = mcv.LastModify
			fmt.Println("uploader is", mcv.Uploader)
			switch mcv.Uploader {
			case OSS:
				meta.uploader = &Oss{ctx: ctx}
			case COS:
				meta.uploader = &Cos{ctx: ctx}
			default:
				return nil, errors.New("undefined uploader" + strconv.FormatInt(int64(mcv.Uploader), 10))
			}
			meta.uploaderSrc = mcv.Uploader
			meta.uploader.Reset(meta)
			return meta, nil

		} else {
			return nil, err
		}

	} else {
		fmt.Println("get from cache", uploadID)
		return nil, err
	}
}
func FromDigest(ctx context.Context, cache cache.ICache, handle int8, digest string, opts ...OptionFunc) (*Meta, error) {
	meta := &Meta{
		ctx:    ctx,
		Status: InitializeStatus,
		cache:  cache,
	}
	if bs, err := meta.cache.Get(digest); err == nil {
		return FromUploadID(ctx, cache, string(bs))
	}

	for _, opt := range opts {
		opt(meta)
	}
	switch handle {
	case OSS:
		meta.uploader = &Oss{ctx: ctx}
	case COS:
		meta.uploader = &Cos{ctx: ctx}
	default:
		return nil, errors.New("undefined uploader")

	}
	meta.uploaderSrc = handle
	meta.uploader.Reset(meta)
	return meta, nil

}

func (m *Meta) Init() error {
	if requestID, err := m.uploader.Init(); err == nil {
		m.UploadID = requestID
		chunks := int(math.Ceil(float64(m.Size) / float64(m.ChunkSize)))
		fmt.Println("init size", chunks)
		m.chunks = []*ChunkETag{}
		m.Status = PendingStatus
		return m.Save()
	} else {
		return err
	}

}
func (m *Meta) Save() error {
	cacheValue := &MetaCacheValue{
		UploadID:   m.UploadID,
		Object:     m.Object,
		Size:       m.Size,
		ChunkSize:  m.ChunkSize,
		Chunks:     m.chunks,
		Digest:     m.Digest,
		Status:     m.Status,
		Uploader:   m.uploaderSrc,
		LastModify: time.Now().Unix(),
	}
	if bs, err := json.Marshal(cacheValue); err == nil {
		if err := m.cache.Set(m.UploadID, bs); err != nil {
			return err
		}
		fmt.Println("set upload_id cache", m.UploadID)
		goto cache
	} else {
		return err
	}
cache:
	return m.cache.Set(m.Digest, []byte(m.UploadID))
}
func (m *Meta) Touch() int8 {
	return m.Status
}
func (m *Meta) Pending() []*ChunkETag {
	return m.chunks
}
func (m *Meta) Append(ctx context.Context, index *ChunkETag) {
	m.chunks = append(m.chunks, index)
}
func (m *Meta) Url() string {
	return m.uploader.Url()
}
func (m *Meta) UploadPart(index int, chunk []byte) (*ChunkETag, error) {
	if etag, err := m.uploader.Upload(index, chunk); err == nil {
		chunkPart := ChunkETag{
			Index: index,
			ETag:  etag,
		}
		m.Append(m.ctx, &chunkPart)
		return &chunkPart, nil
	} else {
		return nil, err
	}
}
func (m *Meta) Merge(chunks []ChunkETag) error {
	var _chunks ChunksETag = chunks
	sort.Sort(_chunks)
	if err := m.uploader.Merge(_chunks); err == nil {
		m.Status = FulfilledStatus
		m.chunks = []*ChunkETag{}
		return m.Save()
	} else {
		return err
	}

}
