package react

// reactor manages the creation of cells
type reactor struct{}

func New() Reactor {
	return &reactor{}
}

// cell handles both Input and Compute logic
type cell struct {
	value     int
	reactor   *reactor
	deps      []*cell
	children  []*cell
	compute   func() int
	callbacks map[int]func(int)
	nextID    int
}

type canceler struct {
	cell *cell
	id   int
}

func (c *canceler) Cancel() {
	delete(c.cell.callbacks, c.id)
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	if c.value == value {
		return
	}
	c.value = value

	// Track original values of all affected compute cells to decide on callbacks
	changedCells := make(map[*cell]int)
	c.propagate(changedCells)

	// Fire callbacks only if the value actually changed after the full propagation
	for cc, oldVal := range changedCells {
		if cc.value != oldVal {
			for _, cb := range cc.callbacks {
				cb(cc.value)
			}
		}
	}
}

// propagate updates downstream cells recursively
func (c *cell) propagate(changedCells map[*cell]int) {
	for _, child := range c.children {
		if _, ok := changedCells[child]; !ok {
			changedCells[child] = child.value
		}
		child.value = child.compute()
		child.propagate(changedCells)
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	if c.callbacks == nil {
		c.callbacks = make(map[int]func(int))
	}
	id := c.nextID
	c.nextID++
	c.callbacks[id] = callback
	return &canceler{cell: c, id: id}
}

// Factory Methods

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{value: initial, reactor: r}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	d := dep.(*cell)
	c := &cell{
		reactor: r,
		deps:    []*cell{d},
		compute: func() int { return compute(d.Value()) },
	}
	c.value = c.compute()
	d.children = append(d.children, c)
	return c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	d1, d2 := dep1.(*cell), dep2.(*cell)
	c := &cell{
		reactor: r,
		deps:    []*cell{d1, d2},
		compute: func() int { return compute(d1.Value(), d2.Value()) },
	}
	c.value = c.compute()
	d1.children = append(d1.children, c)
	d2.children = append(d2.children, c)
	return c
}
