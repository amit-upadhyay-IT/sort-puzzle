package game

type Tube struct {
	capacity int
	blocks   []*Block
}

// GetTube constructs the tube in order the first color goes down the tube
// last color passed is kept at the top of tube
func GetTube(capacity int, colors... Color) *Tube {

	// validate the length of colors should not exceed capacity
	if len(colors) > capacity {
		// TODO: return error in such case
		return nil
	}

	blks := []*Block{}
	for _, col := range colors {
		blks = append(blks, GetBlock(col))
	}

	return &Tube{
		capacity: capacity,
		blocks:   blks,
	}
}

func (t *Tube) GetCapacity() int {
	return t.capacity
}

func (t *Tube) GetLength() int {
	return len(t.blocks)
}

func (t *Tube) IsEmpty() bool {
	return t.GetLength() == 0
}

func (t *Tube) IsFull() bool {
	return t.GetLength() == t.capacity
}

func (t *Tube) GetEmptyLength() int {
	return t.capacity - t.GetLength()
}

// IsTubeSorted returns true if tube is filled completely with all
// blocks having same color
func (t *Tube) IsTubeSorted() bool {

	if t.IsEmpty() {
		// an empty tube is sorted
		return true
	}

	if !t.IsFull() && !t.IsEmpty() {
		return false
	}

	baseColor := t.blocks[0].GetColor()

	for _, blk := range t.blocks {
		if blk.GetColor() != baseColor {
			return false
		}
	}

	return true
}

func (t *Tube) HasSameColor() bool {

	if t.IsEmpty() {
		return true
	}

	baseColor := t.blocks[0].GetColor()

	for _, blk := range t.blocks {
		if blk.GetColor() != baseColor {
			return false
		}
	}
	return true
}

// GetTopBlock will return the top block present in the tube
// NOTE, this can throw nil, caller should be responsible for nil check
func (t *Tube) GetTopBlock() *Block {

	if t.IsEmpty() {
		return nil
	}

	return t.blocks[len(t.blocks)-1]
}

// GetTopBlocksSet is suppose to return list of blocks on top having same color
func (t *Tube) GetTopBlocksSet() []*Block {

	var blocks []*Block

	prevColor := t.blocks[len(t.blocks)-1].GetColor()
	for i := len(t.blocks)-1; i >= 0; i-- {
		// if previous block color is not matching the current block color, we need to stop accumulating blocks
		if prevColor != t.blocks[i].GetColor() {
			break
		}
		// if top color(prev color in this construct) matches current color add it to block stack
		if prevColor == t.blocks[i].GetColor() {
			blocks = append(blocks, t.blocks[i])
		}
		prevColor = t.blocks[i].GetColor()
	}

	return blocks
}

func (t *Tube) Push(blocks []*Block) bool {

	if t.IsFull() {
		// TODO: return error as well
		return false
	}

	if t.GetEmptyLength() < len(blocks) {
		// TODO: return error as well
		return false
	}

	t.blocks = append(t.blocks, blocks...)

	return true
}

func (t *Tube) Pop() (poppedBlocks []*Block, isOpSuccess bool) {

	if t.IsEmpty() {
		// TODO: return error
		return nil, false
	}

	isOpSuccess = true
	poppedBlocks = t.GetTopBlocksSet()

	// remove the block
	t.blocks = t.blocks[:len(t.blocks)-len(poppedBlocks)]

	return
}