
/// Contains all buffered data, so we don't have to regenerate anything in history.
#[derive(Clone)]
pub enum HistoryFrame {
    Home, //TODO
    Album(Vec<String>, String), //for singles, artists is empty and name is "Singles"
    Queue, //TODO
    Search, //TODO
    ArtistPage, //TODO
    NowPlaying, //TODO
    Lyrics, //TODO
}

pub struct History {
    left_frames: Vec<HistoryFrame>, //back in history, currently on end of this vec
    right_frames: Vec<HistoryFrame>, //forward in history
}
impl History {
    /// Returns the new frame (unless it errors), plus the sizes of both stacks.
    pub fn back(&mut self) -> Result<(HistoryFrame, usize, usize), (usize, usize)> {
        if self.left_frames.len() < 2 {
            return Err((self.left_frames.len(), self.right_frames.len()));
        }
        let old_cur_frame = self.left_frames.pop().unwrap(); //safe unwrap
        self.right_frames.push(old_cur_frame);
        Ok((self.left_frames.last().unwrap().clone(), self.left_frames.len(), self.right_frames.len()))
    }

    /// Returns the new frame (unless it errors), plus the sizes of both stacks.
    pub fn forward(&mut self) -> Result<(HistoryFrame, usize, usize), (usize, usize)> {
        if self.right_frames.len() < 1 {
            return Err((self.left_frames.len(), self.right_frames.len()));
        }
        let new_frame = self.right_frames.pop().unwrap(); //safe unwrap
        let cur_frame = new_frame.clone();
        self.left_frames.push(new_frame);
        Ok((cur_frame, self.left_frames.len(), self.right_frames.len()))
    }

    /// Add a new frame to the history.
    pub fn push(&mut self, frame: HistoryFrame) {
        self.left_frames.push(frame);
    }
}





