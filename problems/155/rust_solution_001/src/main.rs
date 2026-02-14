struct MinStack {
    stack: Vec<i32>,
    mins: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {
    fn new() -> Self {
        Self {
            stack: Vec::new(),
            mins: Vec::new(),
        }
    }

    fn push(&mut self, val: i32) {
        self.stack.push(val);

        match self.mins.last() {
            Some(last) => {
                if val <= *last {
                    self.mins.push(val);
                }
            }
            None => self.mins.push(val),
        }
    }

    fn pop(&mut self) {
        match self.stack.pop() {
            Some(last) => {
                if let Some(lastMin) = self.mins.last() {
                    if *lastMin == last {
                        self.mins.pop();
                    }
                }
            }
            None => {}
        }
    }

    fn top(&self) -> i32 {
        *self.stack.last().unwrap_or(&0)
    }

    fn get_min(&self) -> i32 {
        *self.mins.last().unwrap_or(&0)
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * let obj = MinStack::new();
 * obj.push(val);
 * obj.pop();
 * let ret_3: i32 = obj.top();
 * let ret_4: i32 = obj.get_min();
 */
fn main() {
    println!("Hello, world!");
}
