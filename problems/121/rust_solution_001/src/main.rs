fn main() {
    println!("Hello, world!");
}

// start at day 1 (idx 0)
// if the next day is less then the previous one, change it
// keep the window checking the next ones

struct Solution;

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut l = 0;
        let mut max_profit = 0;

        for r in 0..prices.len() {
            let profit = prices[r] - prices[l];
            max_profit = max_profit.max(profit);

            if prices[r] < prices[l] {
                l = r
            }
        }

        max_profit
    }
}
