package texdatastringtable

import (
	"strings"
)

const StringTableNullTerminator = "\x00"

// TexData info lookup table
// Constructs from TexDataStringData & TexDataStringTable lump datas
type TexDataStringTable struct {
	data string
	lookupTable []int32
}

func (table *TexDataStringTable) GetString(stringID int) string {
	end := strings.Index(table.data[table.lookupTable[stringID]:], StringTableNullTerminator)
	if end == -1 {
		return table.data[table.lookupTable[stringID]:]
	}
	return strings.Split(table.data[table.lookupTable[stringID]:], StringTableNullTerminator)[0]
}

func (table *TexDataStringTable) AddOrFindString(s string) int {
	// garymcthack: Make this use an RBTree!
	for i := 0; i < len(table.lookupTable); i++ {
		end := strings.Index(table.data[table.lookupTable[i]:], StringTableNullTerminator)
		if end > 0 {
			if strings.ToLower(s) == strings.ToLower(table.data[table.lookupTable[i]:end]) {
				return i
			}
		} else {
			if strings.ToLower(s) == strings.ToLower(table.data[table.lookupTable[i]:]) {
				return i
			}
		}
	}

	// @TODO validate this
	// this may be invalid
	table.data += StringTableNullTerminator + s
	outOffset := len(table.data)
	table.lookupTable = append(table.lookupTable, int32(outOffset))

	return len(table.lookupTable) - 1
}


func NewTable(data string, lookupTable []int32) *TexDataStringTable{
	return &TexDataStringTable{
		data: data,
		lookupTable: lookupTable,
	}
}
