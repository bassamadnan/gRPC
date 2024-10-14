package utils

import (
	crdt "docs/crdt"
	dpb "docs/pkg/proto/docs"
)

func GetDocumentProto(doc crdt.Document) *dpb.Document {

	characters := make([]*dpb.Character, 0, len(doc.Characters))
	for _, charDoc := range doc.Characters {
		characters = append(characters, &dpb.Character{
			Id:         charDoc.ID,
			Visible:    charDoc.Visible,
			Value:      charDoc.Value,
			IdPrevious: charDoc.IDPrevious,
			IdNext:     charDoc.IDNext,
		})
	}
	return &dpb.Document{
		Document: characters,
	}
}

func GetDocument(doc *dpb.Document) *crdt.Document {
	characters := make([]crdt.Character, 0, len(doc.Document))
	for _, charProto := range doc.Document {
		characters = append(characters, crdt.Character{
			ID:         charProto.Id,
			Visible:    charProto.Visible,
			Value:      charProto.Value,
			IDPrevious: charProto.IdPrevious,
			IDNext:     charProto.IdNext,
		})
	}
	return &crdt.Document{
		Characters: characters,
	}
}

// func GetMessage(msg *dpb.Message) crdt.
