use std::collections::HashMap;

use proconio::input;

enum Query {
    Query1(String, usize),
    Query2(String),
}

fn main() {
    input! { q: usize }

    let queries = (0..q).map(|_| {
        input! { query: usize }
        match query {
            1 => {
                input! { x: String, y: usize }
                Query::Query1(x, y)
            }
            2 => {
                input! { x: String }
                Query::Query2(x)
            }
            _ => unreachable!(),
        }
    });

    let mut map = HashMap::new();
    for query in queries {
        match query {
            Query::Query1(x, y) => {
                map.insert(x, y);
            }
            Query::Query2(x) => {
                if let Some(y) = map.get(&x) {
                    println!("{y}");
                }
            }
        }
    }
}
