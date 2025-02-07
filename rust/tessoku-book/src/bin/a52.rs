use std::collections::VecDeque;

use proconio::input;

enum Query {
    Query1(String),
    Query2,
    Query3,
}

fn main() {
    input! { q: usize }

    let queries = (0..q).map(|_| {
        input! { query: usize }
        match query {
            1 => {
                input! { s: String }
                Query::Query1(s)
            }
            2 => Query::Query2,
            3 => Query::Query3,
            _ => unreachable!(),
        }
    });

    let mut queue = VecDeque::new();
    for query in queries {
        match query {
            Query::Query1(s) => queue.push_back(s),
            Query::Query2 => {
                if let Some(s) = queue.front() {
                    println!("{s}");
                }
            }
            Query::Query3 => {
                queue.pop_front();
            }
        }
    }
}
