use std::cmp::Ordering;
use std::collections::BinaryHeap;

use proconio::input;

#[derive(PartialEq, Eq)]
struct Item(usize);

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(&other))
    }
}

impl Ord for Item {
    fn cmp(&self, other: &Self) -> Ordering {
        other.0.cmp(&self.0)
    }
}

enum Query {
    Query1(usize),
    Query2,
    Query3,
}

fn main() {
    input! { q: usize }

    let queries = (0..q).map(|_| {
        input! { query: usize }
        match query {
            1 => {
                input! { n: usize }
                Query::Query1(n)
            }
            2 => Query::Query2,
            3 => Query::Query3,
            _ => unreachable!(),
        }
    });

    let mut pq = BinaryHeap::new();
    for query in queries {
        match query {
            Query::Query1(n) => pq.push(Item(n)),
            Query::Query2 => {
                if let Some(&Item(n)) = pq.peek() {
                    println!("{n}");
                }
            }
            Query::Query3 => {
                pq.pop();
            }
        }
    }
}
