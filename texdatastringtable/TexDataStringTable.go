package texdatastringtable

import (
	"strings"
	"errors"
)

const StringTableNullTerminator = "\x00"

// TexData info lookup table
// Constructs from TexDataStringData & TexDataStringTable lump datas
type TexDataStringTable struct {
	data string
	lookupTable []int32
}

// GetString
// Find a string by StringID. StringID comes from TexDataStringData, and
// is a lookup to TexDataStringTable null escaped strings
func (table *TexDataStringTable) GetString(stringID int) (string,error) {
	// Ensure that we can't go out of bounds
	if stringID >= len(table.lookupTable) {
		return "", errors.New("invalid texdata string id")
	}
	if table.lookupTable[stringID] >= int32(len(table.data)) {
		return "", errors.New("texdata string data/table mismatch. string id lookup extends out of bounds")
	}
	end := strings.Index(table.data[table.lookupTable[stringID]:], StringTableNullTerminator)
	if end == -1 {
		return table.data[table.lookupTable[stringID]:],nil
	}
	return strings.Split(table.data[table.lookupTable[stringID]:], StringTableNullTerminator)[0],nil
}

// AddOrFindString
// Adds a new string to the Table, unless it exists
// Returns the stringID of the newly added string, or the existing one if found
func (table *TexDataStringTable) AddOrFindString(s string) (int,error) {
	// garymcthack: Make this use an RBTree!
	for i := 0; i < len(table.lookupTable); i++ {
		end := strings.Index(table.data[table.lookupTable[i]:], StringTableNullTerminator)
		if end > 0 {
			if strings.ToLower(s) == strings.ToLower(table.data[table.lookupTable[i]:int(table.lookupTable[i]) + end]) {
				return i, nil
			}
		} else {
			if strings.ToLower(s) == strings.ToLower(table.data[table.lookupTable[i]:]) {
				return i, nil
			}
		}
	}

	// @TODO validate this
	// this may be invalid
	table.data += StringTableNullTerminator + s
	outOffset := len(table.data)
	table.lookupTable = append(table.lookupTable, int32(outOffset))

	return len(table.lookupTable) - 1, nil
}

// NewTable
// Create a new table from the contents of the
// TexDataStringData and TexDataStringTable lumps
func NewTable(data string, lookupTable []int32) *TexDataStringTable{
	return &TexDataStringTable{
		data: data,
		lookupTable: lookupTable,
	}
}
