fn main() {
    println!("Hello, world!");
}

// What I care is the max size (X) possible and 1. This is the window I need to verify.
// There is a point where the value K becomes valid and it's valid until X (max).

// Option 1: Brute for this starting at 1 to X
// Option 2: Binary Search on this range
struct Solution;
impl Solution {
    pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
        let mut left = 1;
        let mut right = piles.iter().max().unwrap().clone();

        while left < right {
            let mid = left + (right - left) / 2;

            if Self::is_valid(&piles, h, mid) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        left
    }

    pub fn is_valid(piles: &Vec<i32>, h: i32, k: i32) -> bool {
        let mut actual = 0;

        for &p in piles.iter() {
            let div = (p + k - 1) / k; // rounding up e.g. 3.2 to 4.
            actual += div;
        }

        actual <= h
    }
}
