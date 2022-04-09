use crate::cucumber::*;

use std::{collections::HashMap, fmt::Display};

type CucumberKey = (usize, usize);

pub struct Grid {
    cucumbers: HashMap<CucumberKey, Cucumber>,
    max_row: usize,
    max_col: usize,
}

impl Grid {
    pub fn new(max_row: usize, max_col: usize) -> Self {
        Self {
            cucumbers: HashMap::new(),
            max_row,
            max_col,
        }
    }

    pub fn take_step(&mut self) -> bool {
        let x = self.move_heard(CucumberDirection::East);
        let y = self.move_heard(CucumberDirection::South);
        x || y
    }

    fn move_heard(&mut self, direction: CucumberDirection) -> bool {
        let mut has_moved = false;

        let heard: Vec<CucumberKey> = self.filter_heard(direction);

        if heard.len() > 0 {
            has_moved = true;
        }

        for key in heard.into_iter() {
            self.move_cucumber(key)
        }

        has_moved
    }

    fn move_cucumber(&mut self, key: CucumberKey) {
        if let Some(cucumber) = self.cucumbers.get_mut(&key) {
            let old_key = (cucumber.row, cucumber.col);
            cucumber.take_step(self.max_row, self.max_col);

            if let Some(c) = self.cucumbers.remove(&old_key) {
                self.add_cucumber(c)
            }
        }
    }

    fn filter_heard(&mut self, direction: CucumberDirection) -> Vec<CucumberKey> {
        self.cucumbers
            .values()
            .filter_map(|c: &Cucumber| {
                if c.direction.eq(&direction) {
                    let next_step = c.next_step(self.max_row, self.max_col);
                    if self.cucumbers.get(&next_step).is_none() {
                        Some((c.row, c.col))
                    } else {
                        None
                    }
                } else {
                    None
                }
            })
            .collect()
    }

    pub fn add_cucumber(&mut self, cucumber: Cucumber) {
        self.cucumbers
            .insert((cucumber.row, cucumber.col), cucumber);
    }

    fn get_cucumber(&self, row: usize, col: usize) -> Option<&Cucumber> {
        self.cucumbers.get(&(row, col))
    }
}

impl Display for Grid {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        writeln!(f, "")?;
        for i in 0..self.max_row {
            for j in 0..self.max_col {
                let cucumber = self.get_cucumber(i, j);
                match cucumber {
                    Some(c) => write!(f, "{}", c)?,
                    None => write!(f, ".",)?,
                };
            }
            writeln!(f, "")?;
        }
        writeln!(f, "")
    }
}
