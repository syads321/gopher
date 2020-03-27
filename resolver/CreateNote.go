package resolver

import (
	types "eaciit/gopher/types"
)

func (r *RootResolver) CreateNote(args types.CreateNoteArgs) (*NoteResolver, error) {
	// Find user:
	var note *types.Note
	for _, user := range users {
		// Create a note with a note ID of n-010:
		note = &types.Note{UserID: "n-010", Data: args.Note.Data}
		user.Notes = append(user.Notes, note) // Push note.
	}
	// Return note:
	return &NoteResolver{note}, nil
}
