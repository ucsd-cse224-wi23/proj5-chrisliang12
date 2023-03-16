package surfstore

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
)

// Implement the logic for a client syncing with the server here.
func ClientSync(client RPCClient) {
	// read local files
	localFiles, err := ioutil.ReadDir(client.BaseDir)
	if err != nil {
		log.Println("Error: reading basedir files. ", err)
	}

	localIndex, err := LoadMetaFromMetaFile(client.BaseDir)
	if err != nil {
		log.Println("Error: loading meta from local index. ", err)
	}

	remoteIndex := make(map[string]*FileMetaData)
	if err := client.GetFileInfoMap(&remoteIndex); err != nil {
		log.Println("Error: loading remote index", err)
	}

	localDirMap := make(map[string][]string)
	uploadFail := false
	for _, fileInfo := range localFiles {
		if fileInfo.Name() == "index.db" {
			continue
		}

		// read file data
		file, err := os.Open(client.BaseDir + "/" + fileInfo.Name())
		if err != nil {
			log.Println("Error: opening file (", fileInfo.Name(), ") | err msg: ", err)
		}
		numBlock := int(math.Ceil(float64(fileInfo.Size()) / float64(client.BlockSize)))
		// compute hashlist
		for i := 0; i < numBlock; i++ {
			buf := make([]byte, client.BlockSize)
			len, err := file.Read(buf)
			if err != nil {
				log.Println("Error: reading file (", fileInfo.Name(), ") | err msg: ", err)
			}
			buf = buf[:len]
			h := GetBlockHashString(buf)
			localDirMap[fileInfo.Name()] = append(localDirMap[fileInfo.Name()], h)
		}

		localMetaData, localIndexOk := localIndex[fileInfo.Name()]
		remoteMetaData, remoteIndexOk := remoteIndex[fileInfo.Name()]

		log.Println("localIndexOk: ", localIndexOk, " | remoteIndexOk: ", remoteIndexOk)

		if localIndexOk && remoteIndexOk {
			log.Println("<", fileInfo.Name(), "> is in localindex and remote index")
			// not a new file
			localVersion := localMetaData.Version
			remoteVersion := remoteMetaData.Version
			if reflect.DeepEqual(localDirMap[fileInfo.Name()], localMetaData.BlockHashList) {
				// local index matches the local file
				if localVersion > remoteVersion {
					// if local version is higher, need to update remote
					err := upload(client, localMetaData, numBlock)
					if err != nil {
						log.Println("Error: uploading <", fileInfo.Name(), "> | msg: ", err)
						uploadFail = true
					}
				}
			} else {
				// local file has a uncommitted change

				// if local index ver == remote index ver, then upload
				if localVersion == remoteVersion {
					err := upload(client, &FileMetaData{Filename: fileInfo.Name(), Version: localVersion, BlockHashList: localDirMap[fileInfo.Name()]}, numBlock)
					if err != nil {
						log.Println("Error: uploading <", fileInfo.Name(), "> | msg: ", err)
						uploadFail = true
					}
					localIndex[fileInfo.Name()].BlockHashList = localDirMap[fileInfo.Name()]
					localIndex[fileInfo.Name()].Version++
				}
			}
		} else {
			// new file
			log.Println("<", fileInfo.Name(), "> is a new file")
			if !remoteIndexOk {
				var tempFileMetaData *FileMetaData
				if localIndexOk {
					tempFileMetaData = localMetaData
				} else {
					tempFileMetaData = &FileMetaData{Filename: fileInfo.Name(), Version: 0, BlockHashList: localDirMap[fileInfo.Name()]}
				}
				err := upload(client, tempFileMetaData, numBlock)
				if err != nil {
					uploadFail = true
				}
				if err := client.GetFileInfoMap(&remoteIndex); err != nil {
					log.Println("Error: loading remote index", err)
				}
				// PrintMetaMap(remoteIndex)
				if err != nil {
					log.Println("Error: uploading <", fileInfo.Name(), "> | msg: ", err)
				}
				if !localIndexOk {
					localIndex[fileInfo.Name()] = tempFileMetaData
				}
			}
		}
	}

	if uploadFail {
		log.Fatal("Upload Fail")
	}

	// check for deletion
	for filename, localMetaData := range localIndex {
		if _, ok := localDirMap[filename]; !ok {
			// empty file
			if len(localMetaData.BlockHashList) == 0 {
				continue
			}

			if localMetaData.BlockHashList[0] == "0" && len(localMetaData.BlockHashList) == 1 {
				continue
			}

			var latestVersion int32
			err := client.UpdateFile(&FileMetaData{Filename: filename, Version: int32(localMetaData.Version + 1), BlockHashList: []string{"0"}}, &latestVersion)
			if err != nil {
				log.Fatal(err)
			}
			if latestVersion == localMetaData.Version+1 {
				localIndex[filename].Version++
				localIndex[filename].BlockHashList = []string{"0"}
			}

		}
	}

	if err := client.GetFileInfoMap(&remoteIndex); err != nil {
		log.Println("Error: loading remote index", err)
	}

	// check for remote update
	for filename, remoteMetaData := range remoteIndex {
		if localMetaData, ok := localIndex[filename]; ok {
			if localMetaData.Version < remoteMetaData.Version {
				err := download(client, localMetaData, remoteMetaData)
				if err != nil {
					log.Println("Error: downloadindg <", filename, "> | msg: ", err)
				}
				localIndex[filename] = remoteMetaData
			}
		} else {
			// PrintMetaMap(remoteIndex)
			err := download(client, localMetaData, remoteMetaData)
			if err != nil {
				log.Println("Error: downloadindg <", filename, "> | msg: ", err)
			}
			localIndex[filename] = remoteMetaData
		}
	}
	log.Println("local index:")
	// PrintMetaMap(localIndex)

	WriteMetaFile(localIndex, client.BaseDir)
}

func reverseBlockStoreMap(blockStoreMap *map[string][]string, blockToServer *map[string]string) {
	for serverAddr, blockHashes := range *blockStoreMap {
		for _, hash := range blockHashes {
			(*blockToServer)[hash] = serverAddr
		}
	}
}

func upload(client RPCClient, meta *FileMetaData, numBlock int) error {
	var latestVer int32
	path, _ := filepath.Abs(ConcatPath(client.BaseDir, meta.Filename))

	file, err := os.Open(path)
	if err != nil {
		log.Println("Error: opening file | msg: ", err)
		return err
	}
	defer file.Close()

	// get blockStoreMap
	blockStoreMap := make(map[string][]string)
	err = client.GetBlockStoreMap(meta.BlockHashList, &blockStoreMap)
	if err != nil {
		log.Println("Error: calling GetBlockStoreMap | msg: ", err)
		return err
	}

	// reverse blockStoreMap for convinient lookup
	blockToServer := make(map[string]string)
	reverseBlockStoreMap(&blockStoreMap, &blockToServer)

	// upload bocks first
	for i := 0; i < numBlock; i++ {
		buf := make([]byte, client.BlockSize)
		len, err := file.Read(buf)
		if err != nil {
			log.Println("Error: reading file | msg: ", err)
			return err
		}
		buf = buf[:len]

		currHash := GetBlockHashString(buf)
		var success bool
		log.Println("curr block addr: ", blockToServer[currHash])
		err = client.PutBlock(&Block{BlockData: buf, BlockSize: int32(len)}, blockToServer[currHash], &success)
		if err != nil {
			log.Println("Error: calling PutBlock() | msg: ", err)
			return err
		}
	}

	// then update metastore
	err = client.UpdateFile(&FileMetaData{Filename: meta.Filename, Version: meta.Version + 1, BlockHashList: meta.BlockHashList}, &latestVer)
	if err != nil {
		log.Println("Error: calling UpdateFile() | msg: ", err)
		return err
	}
	if latestVer > 0 {
		meta.Version = latestVer
	}
	log.Println("latestVer: ", latestVer)
	return nil
}

func download(client RPCClient, localMeta *FileMetaData, remoteMeta *FileMetaData) error {
	path, _ := filepath.Abs(ConcatPath(client.BaseDir, remoteMeta.Filename))
	file, err := os.Create(path)
	if err != nil {
		log.Println("Error: creating file | msg: ", err)
	}
	defer file.Close()

	// check remote deletion
	if remoteMeta.BlockHashList[0] == "0" && len(remoteMeta.BlockHashList) == 1 {
		if err := os.Remove(path); err != nil {
			log.Println("Error: deleting file | msg: ", err)
			return err
		}
		return nil
	}

	blockStoreMap := make(map[string][]string)
	err = client.GetBlockStoreMap(remoteMeta.BlockHashList, &blockStoreMap)
	if err != nil {
		log.Println("Error: calling GetBlockStoreMap")
		return err
	}

	blockToServer := make(map[string]string)
	reverseBlockStoreMap(&blockStoreMap, &blockToServer)

	// get block
	var buf string
	for _, hash := range remoteMeta.BlockHashList {
		var block Block
		err := client.GetBlock(hash, blockToServer[hash], &block)
		if err != nil {
			log.Println("Error: calling GetBlock() | msg: ", err)
			return err
		}

		buf += string(block.BlockData)
	}

	file.WriteString(buf)

	return nil
}
