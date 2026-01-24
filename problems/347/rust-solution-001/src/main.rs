// Use K to fetch the most frequent elements
// Count the amount of each element, store in a fast lookup
// To sort this to fetch the most frequent, bucket list (alt: max/min heap)

pub fn main() {}

struct Solution;

impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut freq = std::collections::HashMap::new();
        let len = nums.len();

        for n in nums {
            *freq.entry(n).or_insert(0) += 1;
        }

        let mut bucket: Vec<Vec<i32>> = vec![Vec::new(); len + 1];
        for (k, v) in freq {
            bucket[v as usize].push(k);
        }

        let mut res = Vec::with_capacity(len);
        for v in bucket.into_iter().rev() {
            if res.len() >= k as usize {
                break;
            }

            res.extend(v);
        }

        res[..k as usize].to_vec()
    }
}
