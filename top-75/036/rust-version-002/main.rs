
struct Solution;

use std::cmp::max;

impl Solution {
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        // Sort based on the input
        intervals.sort_by_key(|v| v[0]);

        let mut res = Vec::new();
        res.push(intervals[0].clone());
        
        for i in 1..intervals.len(){
            let curr = &intervals[i];

            if let Some(last) = res.last_mut() {
                if curr[0] <= last[1] {
                    // Merge
                    last[1] = max(last[1], curr[1])
                } else {
                    res.push(intervals[i].clone());
                }
            }
        }

        res        
    }
}



fn main() {
    let intervals = vec![
        vec![1,3],
        vec![2,6],
        vec![8,10],
        vec![15,18]
    ];

    let res = Solution::merge(intervals);
    println!("Got: {:?}", res);
    println!("Expected {:?}", vec![[1,6],[8,10],[15,18]])
}
