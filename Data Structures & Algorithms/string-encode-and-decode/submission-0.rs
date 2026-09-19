impl Solution {
    pub fn encode(strs: Vec<String>) -> String {
        let mut encoded = String::new();

        for s in strs {
            encoded.push_str(&s.len().to_string());
            encoded.push('#');
            encoded.push_str(&s);
        }

        encoded
    }

    pub fn decode(s: String) -> Vec<String> {
        let bytes = s.as_bytes();
        let mut result = Vec::new();
        let mut i = 0;

        while i < bytes.len() {
            // Find '#'
            let mut j = i;

            while bytes[j] != b'#' {
                j += 1;
            }

            // Parse length
            let len: usize = s[i..j].parse().unwrap();

            // Start of actual string
            let start = j + 1;

            // Extract exactly `len` bytes
            let word = s[start..start + len].to_string();

            result.push(word);

            // Move to next encoded string
            i = start + len;
        }

        result
    }
}
