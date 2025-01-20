use std::ops::ControlFlow;

use proconio::input;

fn main() {
    input! { x: usize }

    let ans = (1..=50)
        .try_fold(1, |acc, i| {
            if acc >= x {
                ControlFlow::Break(i - 1)
            } else {
                ControlFlow::Continue(acc * i)
            }
        })
        .break_value()
        .unwrap();

    println!("{ans}");
}
