package tab

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Style is a struct that holds the style of a tab
type Style struct {
	Color lipgloss.Color
	Icon  string
	Name  string
}

// Tab is an interface outlining the methods that a tab should implement including bubbletea's model methods
type Tab interface {
	// general fields
	Title() string
	Style() Style
	SetSize(width, height int) Tab
	GetKeyBinds() []key.Binding

	// bubbletea methods
	Init() tea.Cmd
	Update(msg tea.Msg) (Tab, tea.Cmd)
	View() string
}

// NewTab returns a tea.Cmd which sends a message to the main model to create a new tab
func NewTab(sender Tab, title string) tea.Cmd {
	return func() tea.Msg {
		return NewTabMessage{
			Sender: sender,
			Title:  title,
		}
	}
}

// NewTabMessage is a tea.Msg that signals that a new tab should be created
type NewTabMessage struct {
	Sender Tab
	Title  string
}
