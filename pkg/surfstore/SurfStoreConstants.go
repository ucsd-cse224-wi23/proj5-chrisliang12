package surfstore

const DEFAULT_META_FILENAME string = "index.db"

const (
	TOMBSTONE_HASHVALUE string = "0"
	EMPTYFILE_HASHVALUE string = "-1"
)

const (
	FILENAME_INDEX  int = 0
	VERSION_INDEX   int = 1
	HASH_LIST_INDEX int = 2
)

const (
	CONFIG_DELIMITER string = ","
	HASH_DELIMITER   string = " "
)
