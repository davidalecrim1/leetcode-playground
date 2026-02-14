struct Solution;

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut left = 0;
        let mut right = height.len();

        let mut max_area = 0;
        while left < right {
            let distance = height[right] - height[left];
            let min_height = height[right].min(height[left]);

            max_area = max_area.max(min_height * distance);

            if height[right] < height[left] {
                right -= 1
            } else {
                left += 1
            }
        }

        max_area
    }
}

pub fn main() {}
