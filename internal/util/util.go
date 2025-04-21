package util

import (
	"cloud.google.com/go/firestore"
)

func ToFirestoreUpdateArray(props map[string]interface{}) []firestore.Update {
	updFirestore := make([]firestore.Update, len(props))
	for k, v := range props {
		obj := firestore.Update{Path: k, Value: v}
		updFirestore = append(updFirestore, obj)
	}
	return updFirestore

}
