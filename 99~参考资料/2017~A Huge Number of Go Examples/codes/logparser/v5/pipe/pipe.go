
package pipe

// Iterator yields a record.
type Iterator interface {
	Each(func(Record) error) error
}

// Consumer consumes records from an iterator.
type Consumer interface {
	Consume(Iterator) error
}

// Transform represents both a record consumer and producer.
// It has an input and output.
// It takes a single record and provides an iterator for all the records.
type Transform interface {
	Iterator // producer: should never return on yield().err == nil
	Consumer
}
