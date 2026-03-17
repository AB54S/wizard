package tui

import (
  "fmt"
  "log"
  
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/bubbles/textinput"
)

var _ tea.Model = &InitPromptModel{}

const (
  histKey = `History File Path`
)

type InitPromptModel struct {
  inputs  map[string]textinput.Model
  done    bool
  cfgPath string
}

func (i InitPromptModel) Init() tea.Cmd {
  return textinput.Blink
}

func (i InitPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "esq":
      return i, tea.Quit
    case "enter":
      i.done = true
      return i, tea.Quit
    }
  }
  cmd := i.updateInputs(msg)
  return i, cmd
}

func (i *InitPromptModel) updateInputs(msg tea.Msg) tea.Cmd {
  cmds := make([]tea.Cmd, 0)
  for k := range i.inputs {
    if i.inputs[k].Focused() {
      m, cmd := i.inputs[k].Update(msg)
      i.inputs[k] = m
      cmds = append(cmds, cmd)
    }
  }
  return tea.Batch(cmds...)
}

func (i InitPromptModel) View() string {
  if i.done {
    v := i.inputs[histKey]
    if v.Value() == "" {
      v.SetValue(v.Placeholder)
    }
    config := &cfg.Config{
      HistoryFile: v.Value(),
    }
    err := cfg.ToFile(i.cfgPath, config)
    if err != nil {
      return err.Error()
    }
    return "Initialization complete! \n"
  }

  output := string.Builder{}
  // Write input to screen
  for k, v := range i.inputs {
    output.WriteString(k + "\n")
    output.WriteString(v.View())
  }
  return output.String()
}


func Start() {
  p := tea.NewProgram(initialModel())
  if _, err := p.Run(); err != nil {
    log.Fatalf(err)
  }
}

type tickMsg struct{}
type errMsg error

type Model struct {
    textinput textinput.Model
    err				error
}

func initialModel() Model {
  ti := textinput.New()
  ti.Placeholder = "Wizard"
  ti.Focus()
  ti.CharLimit = 156
  ti.Width = 20

  return Model{
    textinput: ti,
    err: 			 nil,
  }
}

func (m Model) Init() tea.Cmd {
  return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.Type {
    case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
      return m, tea.Quit
    }
  case errMsg:
    m.err = msg
    return m, nil
  }
  m.textinput, cmd = m.textinput.Update(msg)
  return m, cmd
}

func (m Model) View() string {
  return fmt.Sprintf(
        "What’s your favorite Pokémon?\n\n%s\n\n%s",
        m.textinput.View(),
        "(esc to quit)",
  ) + "\n"
}