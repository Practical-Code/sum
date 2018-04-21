package storage

import (
	"fmt"
	"log"
	"os"
	"sync"

	pb "github.com/evilsocket/sum/proto"
)

// Records is a thread safe data structure used to
// index and manage vectors (records).
type Records struct {
	sync.RWMutex
	dataPath string
	index    map[uint64]*pb.Record
	nextId   uint64
}

// LoadRecords loads and indexes raw protobuf records from
// the data files found in a given path.
func LoadRecords(dataPath string) (*Records, error) {
	dataPath, files, err := ListPath(dataPath)
	if err != nil {
		return nil, err
	}

	records := make(map[uint64]*pb.Record)
	nfiles := len(files)
	maxId := uint64(0)

	if nfiles > 0 {
		log.Printf("Loading %d data files from %s ...", len(files), dataPath)
		for _, fileName := range files {
			record := new(pb.Record)
			if err := Load(fileName, record); err != nil {
				return nil, err
			}
			records[record.Id] = record
			if record.Id > maxId {
				maxId = record.Id
			}
		}
	}

	return &Records{
		dataPath: dataPath,
		index:    records,
		nextId:   maxId + 1,
	}, nil
}

func (r *Records) pathFor(record *pb.Record) string {
	return r.dataPath + fmt.Sprintf("/%d.dat", record.Id)
}

// ForEach will execute a callback for every record object
// in this container, passing raw *pb.Record objects to it
func (r *Records) ForEach(cb func(record *pb.Record)) {
	r.RLock()
	defer r.RUnlock()
	for _, record := range r.index {
		cb(record)
	}
}

// Size returns the number of records currently loaded.
func (r *Records) Size() uint64 {
	r.RLock()
	defer r.RUnlock()
	return uint64(len(r.index))
}

// NextId sets the next integer to use as a record idetifier.
func (r *Records) NextId(next uint64) {
	r.Lock()
	defer r.Unlock()
	r.nextId = next
}

// Create takes a raw *pb.Record object, idexes that record
// by its unique identifier and flushed to disk.
func (r *Records) Create(record *pb.Record) error {
	r.Lock()
	defer r.Unlock()

	record.Id = r.nextId
	r.nextId++

	// make sure the id is unique
	if _, found := r.index[record.Id]; found == true {
		return fmt.Errorf("Record identifier %d violates the unicity constraint.", record.Id)
	}

	r.index[record.Id] = record

	return Flush(record, r.pathFor(record))
}

// Update takes a raw *pb.Record object and updates the one
// with the specified id with its contents.
func (r *Records) Update(record *pb.Record) error {
	r.Lock()
	defer r.Unlock()

	stored, found := r.index[record.Id]
	if found == false {
		return fmt.Errorf("Record %d not found.", record.Id)
	}

	if record.Meta != nil {
		stored.Meta = record.Meta
	}

	if record.Data != nil {
		stored.Data = record.Data
	}

	return Flush(stored, r.pathFor(stored))
}

// Find returns a *pb.Record object given
// its identifier, or nil if not found.
func (r *Records) Find(id uint64) *pb.Record {
	r.RLock()
	defer r.RUnlock()

	record, found := r.index[id]
	if found == true {
		return record
	}
	return nil
}

// Delete removes a record from the index given its
// identifier, it returns the deleted raw *pb.Record
// object, or nil if not found.
func (r *Records) Delete(id uint64) *pb.Record {
	r.Lock()
	defer r.Unlock()

	record, found := r.index[id]
	if found == false {
		return nil
	}

	delete(r.index, id)

	os.Remove(r.pathFor(record))

	return record
}
