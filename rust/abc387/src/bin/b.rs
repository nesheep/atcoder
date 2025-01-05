use proconio::input;

fn main() {
    input! { x: usize }

    let ans = (1..=9)
        .flat_map(|a| (1..=9).map(move |b| if a * b == x { 0 } else { a * b }))
        .sum::<usize>();

    println!("{ans}");
}
