impl Solution {
    pub fn has_duplicate(nums: Vec<i32>) -> bool {
        let mut seen = HashSet::new();
        nums.into_iter().any(|num| !seen.insert(num))
    }
}
