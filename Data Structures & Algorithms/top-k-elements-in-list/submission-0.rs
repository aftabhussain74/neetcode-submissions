impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut freq = HashMap::new();

         // Count frequencies
        for num in nums {
            *freq.entry(num).or_insert(0) += 1;
        }
         // Sort by frequency
        let mut items: Vec<(i32, i32)> = freq.into_iter().collect();

        items.sort_by(|a, b| b.1.cmp(&a.1));

        // Take top k elements
        items
            .into_iter()
            .take(k as usize)
            .map(|(num, _)| num)
            .collect()
    }
}
