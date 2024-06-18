package tests

import (
	"CollabDoc/pkg/document"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateOperations(t *testing.T) {
	// Initialize StateSynchronizer
	ss := document.NewStateSynchronizer()

	// Create a document
	doc := ss.CreateDocument("doc1")

	// Define operations
	operations := []document.Operation{
		{DocID: "doc1", OpType: "insert", Pos: 0, Content: "Hello World", BaseVersion: doc.Version},
		{DocID: "doc1", OpType: "update", Pos: 6, Length: 5, Content: "Universe", BaseVersion: doc.Version + 1},
		{DocID: "doc1", OpType: "update", Pos: 6, Length: 8, Content: "Everyone", BaseVersion: doc.Version + 2},
		{DocID: "doc1", OpType: "update", Pos: 0, Length: 5, Content: "Hi", BaseVersion: doc.Version + 3},
		{DocID: "doc1", OpType: "update", Pos: 3, Length: 8, Content: "Galaxy", BaseVersion: doc.Version + 4},
		{DocID: "doc1", OpType: "update", Pos: 3, Length: 6, Content: "Beautiful Galaxy", BaseVersion: doc.Version + 5},
		{DocID: "doc1", OpType: "insert", Pos: 19, Content: "!", BaseVersion: doc.Version + 6},
		{DocID: "doc1", OpType: "update", Pos: 0, Length: 2, Content: "Hello", BaseVersion: doc.Version + 7},
		{DocID: "doc1", OpType: "delete", Pos: 6, Length: 10, BaseVersion: doc.Version + 8},
		{DocID: "doc1", OpType: "delete", Pos: 0, Length: 5, BaseVersion: doc.Version + 9},
		{DocID: "doc1", OpType: "backspace", Pos: 8, BaseVersion: doc.Version + 10},
		{DocID: "doc1", OpType: "backspace", Pos: 7, BaseVersion: doc.Version + 11}, // Backspace at position 8
	}

	// Apply operations
	for i, op := range operations {
		ss.UpdateDocument("doc1", op)
		finalDoc, exists := ss.GetDocument("doc1")
		assert.True(t, exists)
		t.Logf("After operation %d (%s): %s", i+1, op.OpType, finalDoc.Content)
	}

	// Retrieve the document
	finalDoc, exists := ss.GetDocument("doc1")
	assert.True(t, exists)

	// Expected final content: "Beautiful!"
	expectedContent := " Galax"

	// Verify final content
	assert.Equal(t, expectedContent, finalDoc.Content)

	// Log the final content for visibility
	t.Logf("Final content of the document: %s", finalDoc.Content)
}
