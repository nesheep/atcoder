use proconio::input;

fn main() {
    input! { n: usize }

    let ans = (1..=n).map(|x| {
        (1..=n).map(move |y| {
            if x.min(n + 1 - x).min(y).min(n + 1 - y) % 2 == 0 {
                '.'
            } else {
                '#'
            }
        })
    });

    ans.for_each(|l| println!("{}", l.collect::<String>()));
}
