package game

type Color uint8

const (
	NONE_COLOR Color = iota
	ORANGE
	LIGHT_GREEN
	DARK_GREEN
	RADIUM
	PURPLE
	PINK
	DARK_BLUE
	LIGHT_BLUE
	RED
	YELLOW
	BROWN
	GREY
)


type Block struct {
	myColor Color
}

func GetBlock(color Color) *Block {
	return &Block{color}
}

func (b *Block) GetColor() Color {
	return b.myColor
}