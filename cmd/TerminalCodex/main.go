package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ahhossain/TerminalCodex/internal/history"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	cursor  int
}

type commandFinishedMsg struct {
	output string
	err    error
}

func intialModel() model {
	log.Println("Getting command history")
	commands := history.GetHistory()
	log.Println("Retrieved command history")
	for i, command := range commands {
		log.Printf("Command %d: %s", i+1, command)
	}
	return model{
		choices: commands,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			log.Println("Detected enter")
			runCommand(m.choices[m.cursor])
			return m, runCommand(m.choices[m.cursor])

		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Which option would you like?\n\n"
	//log.Println("Loading commands into choices")
	for i, choice := range m.choices {
		log.Println("Loaded : ", choice)
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"
	//log.Println("Returning following string to view object", s)
	return s
}

func runCommand(command string) tea.Cmd {
	log.Println("Called RunCommand : ", command)
	return func() tea.Msg {
		cmd := exec.Command("powershell.exe", "-Command", command)
		output, err := cmd.CombinedOutput()
		log.Println(output)
		if err != nil {
			log.Println(err)
		}
		return commandFinishedMsg{
			output: string(output),
			err:    err,
		}
	}
}

func main() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.Println("")
	log.Println("")
	log.Println("Application started.")
	log.Println("Creating New program using NewProgram()")
	p := tea.NewProgram(intialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
	log.Println("Creating New program successfully")
}
