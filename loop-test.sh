#!/bin/bash

echo -e "Loop Test Begins\n"
succ=1

for i in {1..100}
do 
  echo -e "\n------------------Running test #$i------------------------\n"
  if ! make TEST_REGEX=TestRaftLogsConsistent specific-test; then
    echo -e "\n!!!!!!!Error: TestRaftLogsConsistent #$i failed !!!!!!!\n"
    succ=0
    break
  fi

  if ! make TEST_REGEX=TestRaftLogsCorrectlyOverwritten specific-test; then
    echo -e "\n!!!!!!!Error: TestRaftLogsCorrectlyOverwritten #$i failed !!!!!!!\n"
    succ=0
    break
  fi
  echo -e "\n------------------End test #$i------------------------\n"
done

if [ "$succ" -eq 1 ]; then
  echo -e "\n\nAll Tests Passed! Congratulations!"
