// Package generated contains the types for schema 'devdb'.
package generated

// TestTable represents a row from 'devdb.test_table'.
type TestTable struct {
	ClientID   int    `json:"client_id"`   // client_id
	ClientName string `json:"client_name"` // client_name
	ClientAge  int    `json:"client_age"`  // client_age
}
