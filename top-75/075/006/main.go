// Need to count the number of occorrencesx
// Store in something that is sortable to fetch the most freq
// Bucket the size of nums

func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int, len(nums))
    for _, num := range nums {
        freqMap[num]++
    }

    bucket := make([][]int, len(nums)+1)
    for num, freq := range freqMap{
        bucket[freq] = append(bucket[freq], num)
    }

    res := make([]int, 0, k)
    for i := len(nums); i >= 0 && len(res) < k; i--{
        if len(bucket[i]) != 0 {
            res = append(res, bucket[i]...)
        }
    }


    return res[:k]
}
