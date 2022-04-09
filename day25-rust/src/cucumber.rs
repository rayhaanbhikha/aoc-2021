use std::fmt::Display;

#[derive(PartialEq, Eq)]
pub enum CucumberDirection {
    East,
    South,
    Unknown,
}

pub struct Cucumber {
    pub row: usize,
    pub col: usize,
    pub direction: CucumberDirection,
}

impl Cucumber {
    pub fn new(cucumber: &str, row_index: usize, col_index: usize) -> Option<Self> {
        let direction = match cucumber {
            ">" => CucumberDirection::East,
            "v" => CucumberDirection::South,
            _ => CucumberDirection::Unknown,
        };

        if let CucumberDirection::Unknown = direction {
            return None;
        }

        Some(Self {
            row: row_index,
            col: col_index,
            direction,
        })
    }

    pub fn take_step(&mut self, max_row: usize, max_col: usize) {
        let (next_row, next_col) = self.next_step(max_row, max_col);
        self.row = next_row;
        self.col = next_col;
    }

    pub fn next_step(&self, max_row: usize, max_col: usize) -> (usize, usize) {
        match self.direction {
            CucumberDirection::South => ((self.row + 1) % max_row, self.col),
            CucumberDirection::East => ((self.row, (self.col + 1) % max_col)),
            _ => (self.row, self.col),
        }
    }
}

impl Display for Cucumber {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let char_to_display = match self.direction {
            CucumberDirection::East => '>',
            CucumberDirection::South => 'v',
            CucumberDirection::Unknown => 'X',
        };
        write!(f, "{}", char_to_display)
    }
}
