fn main() {
    println!("Hello, world!");
}

struct Solution;
impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut left = 0;
        let mut right = nums.len();

        while left < right {
            let mid = ((right - left) / 2) + left;
            println!("{} {} {}", &left, &right, &mid);

            if nums[mid] == target {
                return mid as i32;
            }

            if target > nums[mid] {
                left = mid + 1
            } else {
                right = mid
            }
        }

        -1 as i32
    }
}
