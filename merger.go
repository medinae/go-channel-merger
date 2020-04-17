package merger

// NewMerger return a new merger struct
func NewMerger() merger {
	return merger{}
}

type merger struct {
}

// Merge merges two integer typed channels into one
func (m merger) Merge(a, b <-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		defer close(merged)

		for b != nil || a != nil {
			select {
			case val, received := <-a:
				if !received {
					// assigning nil value to chan disable the related case in the select statement
					// and avoid loosing CPU resources without reasons when chan.
					a = nil
					continue
				}
				merged <- val

			case val, received := <-b:
				if !received {
					b = nil
					continue
				}
				merged <- val
			}
		}
	}()

	return merged
}
