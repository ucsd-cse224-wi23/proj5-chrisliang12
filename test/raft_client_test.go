package SurfTest

import (
	"fmt"
	"os"
	"testing"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	//	"time"
)

// A creates and syncs with a file. B creates and syncs with same file. A syncs again.
func TestSyncTwoClientsSameFileLeaderFailure(t *testing.T) {
	t.Logf("client1 syncs with file1. client2 syncs with file1 (different content). client1 syncs again.")
	cfgPath := "./config_files/3nodes.txt"
	test := InitTest(cfgPath)
	defer EndTest(test)
	test.Clients[0].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	worker1 := InitDirectoryWorker("test0", SRC_PATH)
	worker2 := InitDirectoryWorker("test1", SRC_PATH)
	defer worker1.CleanUp()
	defer worker2.CleanUp()

	// clients add different files
	file1 := "multi_file1.txt"
	file2 := "multi_file1.txt"
	err := worker1.AddFile(file1)
	if err != nil {
		t.FailNow()
	}
	err = worker2.AddFile(file2)
	if err != nil {
		t.FailNow()
	}
	err = worker2.UpdateFile(file2, "update text")
	if err != nil {
		t.FailNow()
	}

	// client1 syncs
	err = SyncClient("localhost:8080", "test0", BLOCK_SIZE, cfgPath)
	if err != nil {
		t.Fatalf("Sync failed")
	}

	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	test.Clients[0].Crash(test.Context, &emptypb.Empty{})
	test.Clients[1].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})

	// client2 syncs
	err = SyncClient("localhost:8080", "test1", BLOCK_SIZE, cfgPath)
	if err != nil {
		t.Fatalf("Sync failed")
	}

	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})

	// client1 syncs
	err = SyncClient("localhost:8080", "test0", BLOCK_SIZE, cfgPath)
	if err != nil {
		t.Fatalf("Sync failed")
	}

	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})
	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})

	workingDir, _ := os.Getwd()

	// check client1
	_, err = os.Stat(workingDir + "/test0/" + META_FILENAME)
	if err != nil {
		t.Fatalf("Could not find meta file for client1")
	}

	fileMeta1, err := LoadMetaFromDB(workingDir + "/test0/")
	if err != nil {
		t.Fatalf("Could not load meta file for client1")
	}
	if len(fileMeta1) != 1 {
		t.Fatalf("Wrong number of entries in client1 meta file")
	}
	if fileMeta1 == nil || fileMeta1[file1].Version != 1 {
		t.Fatalf("Wrong version for file1 in client1 metadata.")
	}

	c, e := SameFile(workingDir+"/test0/multi_file1.txt", SRC_PATH+"/multi_file1.txt")
	if e != nil {
		t.Fatalf("Could not read files in client base dirs.")
	}
	if !c {
		t.Fatalf("file1 should not change at client1")
	}

	// check client2
	_, err = os.Stat(workingDir + "/test1/" + META_FILENAME)
	if err != nil {
		t.Fatalf("Could not find meta file for client2")
	}

	fileMeta2, err := LoadMetaFromDB(workingDir + "/test1/")
	if err != nil {
		t.Fatalf("Could not load meta file for client2")
	}
	if len(fileMeta2) != 1 {
		t.Fatalf("Wrong number of entries in client2 meta file")
	}
	if fileMeta2 == nil || fileMeta2[file1].Version != 1 {
		t.Fatalf("Wrong version for file1 in client2 metadata.")
	}

	c, e = SameFile(workingDir+"/test1/multi_file1.txt", SRC_PATH+"/multi_file1.txt")
	if e != nil {
		t.Fatalf("Could not read files in client base dirs.")
	}
	if !c {
		t.Fatalf("wrong file2 contents at client2")
	}
}

func TestSyncTwoClientsClusterFailure(t *testing.T) {
	t.Logf("client1 syncs with file1. client2 syncs. majority of the cluster crashes. client2 syncs again.")
	cfgPath := "./config_files/3nodes.txt"
	test := InitTest(cfgPath)
	defer EndTest(test)
	test.Clients[0].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	worker1 := InitDirectoryWorker("test0", SRC_PATH)
	worker2 := InitDirectoryWorker("test1", SRC_PATH)
	defer worker1.CleanUp()
	defer worker2.CleanUp()

	// clients add different files
	file1 := "multi_file1.txt"
	if err := worker1.AddFile(file1); err != nil {
		t.FailNow()
	}

	// client1 sync with file1
	if err := SyncClient("localhost:8080", "test0", BLOCK_SIZE, cfgPath); err != nil {
		t.Fatalf("Client1 Sync failed")
	}

	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	// client2 syncs
	err := worker1.UpdateFile(file1, "update text")

	err = SyncClient("localhost:8080", "test1", BLOCK_SIZE, cfgPath)
	if err != nil {
		t.Fatalf("Client2 Sync Failed")
	}
	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	test.Clients[0].Crash(test.Context, &emptypb.Empty{})
	test.Clients[2].Crash(test.Context, &emptypb.Empty{})

	test.Clients[1].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})

	// client2 syncs again, should fail
	err = SyncClient("localhost:8080", "test1", BLOCK_SIZE, cfgPath)
	if err == nil {
		t.Fatalf("SyncClient should fail")
	}
}

func TestRaftLogsCorrectlyOverwritten(t *testing.T) {
	t.Logf("leader1 gets several requests while all other nodes are crashed. leader1 crashes. all other nodes are restored. leader2 gets a request. leader1 is restored.")
	cfgPath := "./config_files/3nodes.txt"
	test := InitTest(cfgPath)
	defer EndTest(test)
	test.Clients[0].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	worker1 := InitDirectoryWorker("test0", SRC_PATH)
	worker2 := InitDirectoryWorker("test1", SRC_PATH)
	defer worker1.CleanUp()
	defer worker2.CleanUp()

	file1 := "multi_file1.txt"
	file2 := "multi_file2.txt"
	err := worker1.AddFile(file1)
	if err != nil {
		t.FailNow()
	}

	err = worker1.AddFile(file2)
	if err != nil {
		t.FailNow()
	}

	err = worker2.AddFile(file1)
	if err != nil {
		t.FailNow()
	}

	err = worker2.AddFile(file2)
	if err != nil {
		t.FailNow()
	}

	err = worker2.UpdateFile(file1, "abcdefg")
	if err != nil {
		t.FailNow()
	}

	err = worker2.UpdateFile(file2, "akdlsjfieqw")
	if err != nil {
		t.FailNow()
	}

	test.Clients[1].Crash(test.Context, &emptypb.Empty{})
	test.Clients[2].Crash(test.Context, &emptypb.Empty{})

	fmt.Println("\n-------- leader 1 start sync--------")
	err = SyncClient("localhost:8080", "test0", BLOCK_SIZE, cfgPath)
	if err == nil {
		t.Fatalf("Sync should fail")
	}
	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})
	fmt.Println("-------- leader 1 end sync-------- \n ")

	internalState_before, err := test.Clients[0].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println("\n------leader1 log (all other nodes crash): \n", internalState_before.Log, "\n ")

	// leader 1 crashes
	test.Clients[0].Crash(test.Context, &emptypb.Empty{})

	// time.Sleep(time.Second)
	// all other nodes restore
	test.Clients[1].Restore(test.Context, &emptypb.Empty{})
	test.Clients[2].Restore(test.Context, &emptypb.Empty{})

	// set leader2
	test.Clients[1].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})

	// leader1 is restored
	test.Clients[0].Restore(test.Context, &emptypb.Empty{})

	// leader 2 get several requests
	fmt.Println("\n-------- leader 2 start sync--------")
	err = SyncClient("localhost:8080", "test1", BLOCK_SIZE, cfgPath)
	if err != nil {
		t.Fatal("Sync failed ", err)
	}

	test.Clients[1].SendHeartbeat(test.Context, &emptypb.Empty{})
	fmt.Println("-------- leader 2 end sync-------- \n ")

	internalStateLeader1, err := test.Clients[0].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	internalStateLeader2, err := test.Clients[1].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println("\n------leader1 log (after leader2 sync): \n", internalStateLeader1.Log, "\n ")

	if len(internalStateLeader1.Log) != len(internalStateLeader2.Log) {
		t.Fatalf("log inconsistent!")
	}

	for i := 0; i < len(internalStateLeader1.Log); i++ {
		if internalStateLeader1.Log[i].Term != internalStateLeader2.Log[i].Term {
			t.Fatalf("log inconsistent!")
		}
	}
}

func TestRaftNewLeaderPushesUpdates(t *testing.T) {
	t.Logf("leader1 gets a request while the majority of the cluster is down. leader1 crashes. the other nodes come back. leader2 is elected")
	cfgPath := "./config_files/5nodes.txt"
	test := InitTest(cfgPath)
	defer EndTest(test)
	test.Clients[0].SetLeader(test.Context, &emptypb.Empty{})
	test.Clients[0].SendHeartbeat(test.Context, &emptypb.Empty{})

	worker1 := InitDirectoryWorker("test0", SRC_PATH)
	worker2 := InitDirectoryWorker("test1", SRC_PATH)
	defer worker1.CleanUp()
	defer worker2.CleanUp()

	file1 := "multi_file1.txt"
	err := worker1.AddFile(file1)
	if err != nil {
		t.FailNow()
	}

	test.Clients[2].Crash(test.Context, &emptypb.Empty{})
	test.Clients[3].Crash(test.Context, &emptypb.Empty{})
	test.Clients[4].Crash(test.Context, &emptypb.Empty{})

	err = SyncClient("localhost:8080", "test0", BLOCK_SIZE, cfgPath)
	if err == nil {
		t.Fatalf("Sync should failed")
	}

	test.Clients[0].Crash(test.Context, &emptypb.Empty{})

	test.Clients[2].Restore(test.Context, &emptypb.Empty{})
	test.Clients[3].Restore(test.Context, &emptypb.Empty{})
	test.Clients[4].Restore(test.Context, &emptypb.Empty{})

	test.Clients[1].SetLeader(test.Context, &emptypb.Empty{})

	err = SyncClient("localhost:8080", "test0", BLOCK_SIZE, cfgPath)
	if err != nil {
		t.Fatalf("Sync failed")
	}

	internalStateLeader1, err := test.Clients[0].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	internalStateLeader2, err := test.Clients[1].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	internalStateLeader3, err := test.Clients[2].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	internalStateLeader4, err := test.Clients[3].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	internalStateLeader5, err := test.Clients[4].GetInternalState(test.Context, &emptypb.Empty{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	if internalStateLeader1.Term != 1 {
		t.Fatalf("leader1 has incorrect meta state")
	}

	if internalStateLeader2.Term != 2 {
		t.Fatalf("leader1 has incorrect meta state")
	}

	if internalStateLeader3.Term != 2 {
		t.Fatalf("leader1 has incorrect meta state")
	}
	if internalStateLeader4.Term != 2 {
		t.Fatalf("leader1 has incorrect meta state")
	}
	if internalStateLeader5.Term != 2 {
		t.Fatalf("leader1 has incorrect meta state")
	}

	// if len(internalStateLeader1.MetaMap.FileInfoMap) != 0 {
	// 	t.Fatalf("leader1 should not commit the data")
	// }
	// if len(internalStateLeader2.MetaMap.FileInfoMap) != 0 {
	// 	t.Fatalf("leader2 should not commit the data")
	// }
	// if len(internalStateLeader3.MetaMap.FileInfoMap) != 0 {
	// 	t.Fatalf("leader3 should not commit the data")
	// }
	// if len(internalStateLeader4.MetaMap.FileInfoMap) != 0 {
	// 	t.Fatalf("leader4 should not commit the data")
	// }
	// if len(internalStateLeader5.MetaMap.FileInfoMap) != 0 {
	// 	t.Fatalf("leader5 should not commit the data")
	// }
}
