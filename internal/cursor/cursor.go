package cursor

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
)

type Cursor struct {
	position   int
	DebugLabel *widget.Label
}

func (c *Cursor) Inc() {
	c.position++
	c.updateDebugLabel()
}

func (c *Cursor) Dec() {
	c.position--
	c.updateDebugLabel()
}

func (c *Cursor) CanBackspace() bool {
	return c.position > 0
}

func (c *Cursor) CurrentPosition() int {
	return c.position
}

func (c *Cursor) updateDebugLabel() {
	c.DebugLabel.Text = fmt.Sprintf("Cursor: %d", c.position)
}

func CreateCursor() *Cursor {
	return &Cursor{position: 0, DebugLabel: widget.NewLabel("Cursor: 0")}
}
