use std::fs::read_to_string;

mod cucumber;
mod grid;

use crate::cucumber::Cucumber;
use crate::grid::Grid;

fn main() {
    let data = read_to_string("./inputs/input").unwrap();

    let rows: Vec<&str> = data.trim().split("\n").collect();

    let max_row = rows.len();
    let max_col = rows
        .get(0)
        .unwrap_or_else(|| &" ")
        .split("")
        .filter(|c| !c.is_empty())
        .count();

    let mut grid = Grid::new(max_row, max_col);

    for (row_index, &row) in rows.iter().enumerate() {
        for (col_index, cucumber_char) in row.trim().split("").filter(|c| !c.is_empty()).enumerate()
        {
            if let Some(c) = Cucumber::new(cucumber_char, row_index, col_index) {
                grid.add_cucumber(c);
            }
        }
    }

    let mut i = 1;
    loop {
        if !grid.take_step() {
            break;
        }
        i += 1
    }

    println!("{}", i);
}
