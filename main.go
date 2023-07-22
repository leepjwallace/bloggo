package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Save notes in a hidden directory in the user's home directory
var (
	notesDir = filepath.Join(os.Getenv("HOME"), ".bloggo")
	logger   = log.New(os.Stdout, "", 0) // 0 removes date and time from cmd
)

func main() {
	if len(os.Args) < 2 {
		logger.Fatal("Expected a command")
	}

	// Create notes directory if it does not exist
	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		os.Mkdir(notesDir, 0755)
	}

	switch os.Args[1] {
	case "create":
		newNote()
	case "get":
		openNote()
	case "list":
		listNotes()
	case "delete":
		deleteNote()
	case "help":
		printHelp()
	default:
		logger.Fatal("Unknown command")
	}
}

func newNote() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	filename := os.Args[2]
	// filename := os.Args[2] + ".txt"

	if len(filename) > 20 {
		logger.Fatal("Filename should be less than 20 characters")
	}

	path := filepath.Join(notesDir, filename)

	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		logger.Fatal(err)
	}
}

func listNotes() {
	files, err := os.ReadDir(notesDir)
	if err != nil {
		logger.Fatal(err)
	}

	if len(files) == 0 {
		logger.Fatal("No notes found")
	}

	for _, file := range files {
		logger.Println(file.Name())
	}
}

func openNote() {
	if len(os.Args) < 3 {
		logger.Fatal("Expected a name for the note, usage: 'bloggo open <note_name>'")
	}

	filename := os.Args[2]
	path := filepath.Join(notesDir, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Fatal("No note found with this name")
	}

	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		logger.Fatal(err)
	}
}

func deleteNote() {
	if len(os.Args) < 3 {
		logger.Fatal("Expected a name for the note, usage: 'bloggo delete <note_name>'")
	}

	filename := os.Args[2]
	path := filepath.Join(notesDir, filename)

	err := os.Remove(path)

	if err != nil {
		logger.Fatal("Failed to delete the note")
	} else {
		fmt.Printf("Note \"%s\" deleted successfully.\n", filename)
	}
}

func printHelp() {
	fmt.Println(`
bloggo is a command line note taking app.

Usage:
	bloggo <command> [arguments]

Available Commands:
	create  Create a new note. Use 'bloggo new <note_name>'.
	list    List all notes.
	get     Open an existing note. Use 'bloggo open <note_name>'.
	delete  Delete an existing note. Use 'bloggo delete <note_name>'.
	help    Show help.
	`)
}
