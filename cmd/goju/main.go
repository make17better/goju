package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/make17better/goju/internal/config"
	"github.com/make17better/goju/internal/learn"
	"github.com/make17better/goju/internal/lookup"
	"github.com/make17better/goju/internal/practise"
	"github.com/make17better/goju/internal/ui"
)

const (
	version  = "0.1.0"
	asciiArt = `
     ██████╗  ██████╗      ██╗██╗   ██╗
    ██╔════╝ ██╔═══██╗     ██║██║   ██║
    ██║  ███╗██║   ██║     ██║██║   ██║
    ██║   ██║██║   ██║██   ██║██║   ██║
    ╚██████╔╝╚██████╔╝╚█████╔╝╚██████╔╝
     ╚═════╝  ╚═════╝  ╚════╝  ╚═════╝ 
`
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Define flags
	helpFlag := flag.Bool("help", false, "Show help information")
	versionFlag := flag.Bool("version", false, "Show version information")
	practiseFlag := flag.Bool("practise", false, "Enter practice mode")
	learnFlag := flag.Bool("learn", false, "Enter learning mode")
	langFlag := flag.String("lang", "", "Set language (en, zh, zh-tw)")
	countFlag := flag.Int("count", cfg.Practice.DefaultCount, "Number of questions for practice")

	flag.Parse()

	// Handle flags
	if *helpFlag {
		printHelp()
		return
	}

	if *versionFlag {
		fmt.Printf("goju version %s\n", version)
		return
	}

	// Update language if specified
	if *langFlag != "" {
		cfg.Language = *langFlag
		if err := config.SaveConfig(cfg); err != nil {
			fmt.Printf("Error saving configuration: %v\n", err)
			os.Exit(1)
		}
	}

	// Default TUI mode
	if !*practiseFlag && !*learnFlag && len(flag.Args()) == 0 {
		tui := ui.NewTUI(cfg)
		if err := tui.Run(); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Handle specific modes
	if *practiseFlag {
		session := practise.NewPracticeSession(*countFlag, cfg.Practice.Categories)
		runPracticeSession(session)
		return
	}

	if *learnFlag {
		content := learn.GetLearningContent(learn.Hard, "both")
		fmt.Println(learn.FormatLearningContent(content, "both"))
		return
	}

	// Handle lookup mode
	if len(flag.Args()) > 0 {
		lookupType := flag.Args()[0]
		values := flag.Args()[1:]
		if len(values) == 0 {
			fmt.Println("Please provide characters to look up")
			os.Exit(1)
		}

		results := lookup.BatchLookup(lookupType, values)
		fmt.Println(lookup.FormatBatchLookup(results))
		return
	}
}

func runPracticeSession(session *practise.PracticeSession) {
	fmt.Println("Starting practice session...")
	fmt.Println("Type 'quit' to exit")

	for i := 0; i < session.Count; i++ {
		char := session.GetNextCharacter()
		session.Current.Character = char
		session.Current.StartTime = time.Now()

		fmt.Printf("\nQuestion %d/%d: What is the romaji for: %s\n", i+1, session.Count, char.Hiragana)
		fmt.Print("Answer: ")

		var answer string
		fmt.Scanln(&answer)

		if answer == "quit" {
			fmt.Println("\nPractice session ended")
			return
		}

		correct := session.CheckAnswer(answer)
		if !correct {
			session.RecordMistake(answer)
			fmt.Printf("Incorrect! The answer is: %s\n", char.Romaji)
			fmt.Println("Press Enter to continue...")
			fmt.Scanln()
		} else {
			session.CompleteQuestion(true)
		}
	}

	// Show results
	fmt.Printf("\nPractice session completed!\n")
	fmt.Printf("Accuracy: %.2f%%\n", session.GetAccuracy())
	if len(session.GetWeaknesses(5)) > 0 {
		fmt.Println("\nWeaknesses:")
		for _, weakness := range session.GetWeaknesses(5) {
			fmt.Printf("- %s (missed %d times)\n", weakness.Character.Romaji, weakness.Attempts)
		}
	}
}

func printHelp() {
	fmt.Printf("%s\n", asciiArt)
	fmt.Println("goju - Japanese fifty-sounds practice tool")
	fmt.Println("\nUsage:")
	fmt.Println("  goju [command] [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  hiragana    Look up hiragana characters")
	fmt.Println("  katakana    Look up katakana characters")
	fmt.Println("  romaji      Look up romaji")
	fmt.Println("\nOptions:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")
	fmt.Println("  -p, --practise Enter practice mode")
	fmt.Println("  -l, --learn    Enter learning mode")
	fmt.Println("  --lang         Set language (en, zh, zh-tw)")
	fmt.Println("  --count        Number of questions for practice")
	fmt.Println("\nExamples:")
	fmt.Println("  goju                    # Launch TUI")
	fmt.Println("  goju hiragana あ        # Look up hiragana")
	fmt.Println("  goju --practise         # Enter practice mode")
	fmt.Println("  goju --learn            # Enter learning mode")
}
