impl Solution {
    pub fn find_the_winner(n: i32, k: i32) -> i32 {
        // use an array to keep track of the friends
        // remove each of them using k until there is 1 left
        // use an index with module to track the position to remove the friends.

        let mut friends: Vec<i32> = Vec::from_iter(1..=n);

        println!("{:?}", friends);

        let mut idx: i32 = 0;
        let mut len: i32 = n;
        while friends.len() > 1 {
            len = friends.len() as i32;
            idx = (idx + k - 1) % len;
            friends.remove(idx as usize);
        }

        friends[0]
    }
}

struct Solution {}

fn main() {}
