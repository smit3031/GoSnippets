# GoSnippets
Just Documenting Go Learnings
class Solution {
    public int totalFruit(int[] fruits) {
        HashMap<Integer, Integer> fruitCount = new HashMap<>();
        int maxFruits = 0;
        int start = 0;

        for (int end = 0; end < fruits.length; end++) {
            int currentFruit = fruits[end];
            fruitCount.put(currentFruit, fruitCount.getOrDefault(currentFruit, 0) + 1);

            while (fruitCount.size() > 2) {
                int leftFruit = fruits[start];
                fruitCount.put(leftFruit, fruitCount.get(leftFruit) - 1);
                if (fruitCount.get(leftFruit) == 0) {
                    fruitCount.remove(leftFruit);
                }
                start++;
            }

            maxFruits = Math.max(maxFruits, end - start + 1);
        }

        return maxFruits;
    }
}