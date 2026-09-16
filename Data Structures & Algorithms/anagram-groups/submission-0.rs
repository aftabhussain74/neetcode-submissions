impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut groups: HashMap<[i32; 26], Vec<String>> = HashMap::new();

        for word in strs {
            let mut count = [0; 26];

            for ch in word.bytes() {
                count[(ch - b'a') as usize] += 1;
            }

            groups.entry(count).or_default().push(word);
        }
        groups.into_values().collect()
    }
}
