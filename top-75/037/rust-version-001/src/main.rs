fn main() {
    let intervals = vec![
        vec![1,2],
        vec![2,3],
        vec![3,4],
        vec![1,3],
    ];

    let res = Solution::erase_overlap_intervals(intervals);
    println!("Expected 1 got {}", res);
}

struct Solution;

// Time Complexity: O(n log n) + O(n) = O(n log n)
impl Solution {
    pub fn erase_overlap_intervals(mut intervals: Vec<Vec<i32>>) -> i32 {
        intervals.sort_unstable_by_key(|k| k[0]);

        let mut res = 0;
        let mut prev = 0;

        for i in 1..intervals.len(){
            if intervals[i][0] >= intervals[prev][1] {
                // non overlapping, just continue.
                prev = i;
            } else {
                res+=1;

                // overlapping, apply greedy and choose one.
                if intervals[i][1] <= intervals[prev][1] {
                    // keep the lowest value.
                    prev = i;
                }

                // else i keep the prev itself because its lowest.
            }
        }

        return res;
    }
}