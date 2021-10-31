package disruptor

import "sync"

type compositeReader []Reader

func (cr compositeReader) Read() {
	var waiter sync.WaitGroup
	waiter.Add(len(cr))

	for _, item := range cr {
		go func(reader Reader) {
			reader.Read()
			waiter.Done()
		}(item)
	}

	waiter.Wait()
}

func (cr compositeReader) Close() error {
	for _, item := range cr {
		_ = item.Close()
	}

	return nil
}
