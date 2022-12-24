func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mergedArr := merge(nums1, nums2)
    i := len(mergedArr) / 2
    if len(mergedArr) % 2 == 0 {
        return float64(float64(mergedArr[i] + mergedArr[i-1])/2)
    } else {
        return float64(mergedArr[i])
    }
}

func merge(n1 []int, n2 []int) []int {
    s := append(n1, n2...)
    sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
    return s
}
