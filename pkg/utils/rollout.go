package utils

import "hash/fnv"

func RolloutBucket(userID string) int {

	h := fnv.New32a()

	_, _ = h.Write([]byte(userID))

	return int(h.Sum32() % 100)
}