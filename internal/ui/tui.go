package ui

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/make17better/goju/internal/config"
	"github.com/make17better/goju/internal/learn"
	"github.com/make17better/goju/internal/practise"
	"github.com/rivo/tview"
)

// TUI represents the terminal user interface
type TUI struct {
	app        *tview.Application
	pages      *tview.Pages
	config     *config.Config
	history    []practise.PracticeResult
	weaknesses []practise.Mistake
}

// NewTUI creates a new TUI instance
func NewTUI(cfg *config.Config) *TUI {
	tui := &TUI{
		app:    tview.NewApplication(),
		pages:  tview.NewPages(),
		config: cfg,
	}

	// Initialize the main menu
	tui.initMainMenu()

	return tui
}

// Run starts the TUI application
func (t *TUI) Run() error {
	return t.app.SetRoot(t.pages, true).Run()
}

// initMainMenu initializes the main menu
func (t *TUI) initMainMenu() {
	// Create the main menu layout
	menu := tview.NewFlex().SetDirection(tview.FlexRow)

	// Add ASCII art
	asciiArt := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(`
     ██████╗  ██████╗      ██╗██╗   ██╗
    ██╔════╝ ██╔═══██╗     ██║██║   ██║
    ██║  ███╗██║   ██║     ██║██║   ██║
    ██║   ██║██║   ██║██   ██║██║   ██║
    ╚██████╔╝╚██████╔╝╚█████╔╝╚██████╔╝
     ╚═════╝  ╚═════╝  ╚════╝  ╚═════╝ 
`)

	// Add mode selection
	modes := tview.NewList().
		AddItem("Learn", "Study the fifty-sounds", 'l', func() {
			t.showLearnMenu()
		}).
		AddItem("Practice", "Test your knowledge", 'p', func() {
			t.showPracticeMenu()
		})

	// Add language selection
	languages := tview.NewList().
		AddItem("English", "", 'e', nil).
		AddItem("简体中文", "", 's', nil).
		AddItem("繁體中文", "", 't', nil)

	// Add history and weaknesses if available
	if len(t.history) > 0 {
		modes.AddItem("History", "View practice history", 'h', func() {
			t.showHistory()
		})
	}

	if len(t.weaknesses) > 0 {
		modes.AddItem("Weaknesses", "Focus on difficult characters", 'w', func() {
			t.showWeaknesses()
		})
	}

	// Add quit option
	modes.AddItem("Quit", "Exit the application", 'q', func() {
		t.app.Stop()
	})

	// Add components to the menu
	menu.AddItem(asciiArt, 0, 1, false)
	menu.AddItem(modes, 0, 1, true)
	menu.AddItem(languages, 0, 1, false)

	t.pages.AddPage("main", menu, true, true)
}

// showLearnMenu shows the learning mode menu
func (t *TUI) showLearnMenu() {
	menu := tview.NewFlex().SetDirection(tview.FlexRow)

	// Add difficulty selection
	difficulties := tview.NewList().
		AddItem("Easy (清音)", "Basic sounds only", 'e', func() {
			t.showLearningContent(learn.Easy)
		}).
		AddItem("Normal (清音・浊音)", "Basic and voiced sounds", 'n', func() {
			t.showLearningContent(learn.Normal)
		}).
		AddItem("Hard (五十音)", "All sounds", 'h', func() {
			t.showLearningContent(learn.Hard)
		}).
		AddItem("Back", "Return to main menu", 'b', func() {
			t.pages.SwitchToPage("main")
		})

	menu.AddItem(difficulties, 0, 1, true)
	t.pages.AddPage("learn", menu, true, false)
	t.pages.SwitchToPage("learn")
}

// showLearningContent displays the learning content
func (t *TUI) showLearningContent(difficulty learn.Difficulty) {
	content := learn.GetLearningContent(difficulty, "both")
	text := tview.NewTextView().
		SetText(learn.FormatLearningContent(content, "both")).
		SetScrollable(true)

	// Add back button
	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
		t.pages.SwitchToPage("learn")
	})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(text, 0, 1, false).
		AddItem(backButton, 1, 0, true)

	t.pages.AddPage("content", layout, true, false)
	t.pages.SwitchToPage("content")
}

// showPracticeMenu shows the practice mode menu
func (t *TUI) showPracticeMenu() {
	menu := tview.NewFlex().SetDirection(tview.FlexRow)

	// Add practice options
	options := tview.NewList().
		AddItem("Start Practice", "Begin a new practice session", 's', func() {
			t.startPractice()
		}).
		AddItem("Back", "Return to main menu", 'b', func() {
			t.pages.SwitchToPage("main")
		})

	menu.AddItem(options, 0, 1, true)
	t.pages.AddPage("practice", menu, true, false)
	t.pages.SwitchToPage("practice")
}

// startPractice starts a new practice session
func (t *TUI) startPractice() {
	session := practise.NewPracticeSession(t.config.Practice.DefaultCount, t.config.Practice.Categories)
	question := tview.NewTextView().SetText("")
	input := tview.NewInputField().SetLabel("Answer: ")

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(question, 0, 1, false).
		AddItem(input, 1, 0, true)

	t.pages.AddPage("practice_session", layout, true, false)
	t.pages.SwitchToPage("practice_session")

	// Set up practice session
	t.setupPracticeSession(session, question, input)
}

// setupPracticeSession sets up the practice session UI
func (t *TUI) setupPracticeSession(session *practise.PracticeSession, question *tview.TextView, input *tview.InputField) {
	// Get first character
	char := session.GetNextCharacter()
	session.Current.Character = char
	session.Current.StartTime = time.Now()

	// Update question display
	question.SetText(fmt.Sprintf("What is the romaji for: %s", char.Hiragana))

	// Handle input
	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			answer := input.GetText()
			correct := session.CheckAnswer(answer)

			if !correct {
				session.RecordMistake(answer)
				// Show correct answer
				question.SetText(fmt.Sprintf("Incorrect! The answer is: %s", char.Romaji))
				// Wait for any key press
				t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
					if event.Key() == tcell.KeyRune {
						// Continue to next question
						t.setupPracticeSession(session, question, input)
					}
					return event
				})
			} else {
				session.CompleteQuestion(true)
				// Continue to next question
				t.setupPracticeSession(session, question, input)
			}
		}
	})
}

// showHistory shows the practice history
func (t *TUI) showHistory() {
	// TODO: Implement history view
}

// showWeaknesses shows the weakness analysis
func (t *TUI) showWeaknesses() {
	// TODO: Implement weaknesses view
}
