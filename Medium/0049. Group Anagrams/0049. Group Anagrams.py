class Solution:

    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:

        anagrams = dict()
        for word in strs:
            key = "".join(sorted(word))
            if key not in anagrams:
                anagrams[key] = list()
            anagrams[key].append(word)

        return anagrams.values()

