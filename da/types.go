package da

import (
	"github.com/misseven0/opencc/cedar"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *cedar.Cedar
	Values [][]string
}
