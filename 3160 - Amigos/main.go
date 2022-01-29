package main

import (
	"bufio"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const size = 32 << 10
	w := bufio.NewWriterSize(os.Stdout, size)
	r := bufio.NewReaderSize(os.Stdin, size)

	friendList := readFriendList(r)
	newFriendList := readFriendList(r)
	friend, err := r.ReadString('\n')
	check(err)

	friend = strings.TrimSpace(friend)

	finalFriendList := make([]string, 0, len(friendList)+len(newFriendList)+1)

	if friend == "nao" {
		finalFriendList = append(finalFriendList, friendList...)
		finalFriendList = append(finalFriendList, newFriendList...)
	} else {
		const FriendNotFound = -1
		idx := FriendNotFound
		for i, f := range friendList {
			if f == friend {
				idx = i
				break
			}
		}

		if idx == FriendNotFound {
			panic("friend not found")
		}

		finalFriendList = append(finalFriendList, friendList[:idx]...)
		finalFriendList = append(finalFriendList, newFriendList...)
		finalFriendList = append(finalFriendList, friendList[idx:]...)
	}

	length := len(finalFriendList)
	if length > 0 {
		w.WriteString(finalFriendList[0])
		for i := 1; i < length; i++ {
			w.WriteByte(' ')
			w.WriteString(finalFriendList[i])
		}
		w.WriteByte('\n')
		w.Flush()
	}
}

func readFriendList(r *bufio.Reader) []string {
	line, err := r.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return strings.Fields(line)
}
