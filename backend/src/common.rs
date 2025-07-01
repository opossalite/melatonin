use std::path::PathBuf;



pub fn expand_tilde(path: &str) -> Option<PathBuf> {
    if let Some(stripped) = path.strip_prefix("~/") {
        if let Some(home_dir) = home::home_dir() {
            return Some(home_dir.join(stripped));
        }
    }
    Some(PathBuf::from(path)) // Not starting with ~, just return as-is
}


