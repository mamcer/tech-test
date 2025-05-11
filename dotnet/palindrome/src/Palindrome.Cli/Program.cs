namespace Palindrome.Cli;

public class Program
{
    public static void Main(string[] args)
    {
        var word = "racecar";
        var isPalindrome = IsPalindrome(word);
        if (isPalindrome)
        {
            Console.WriteLine($"'{word}' is a palindrome.");
        }
        else
        {
            Console.WriteLine($"'{word}' is not a palindrome.");
        }
    }

    /// <summary>
    /// Determines whether the specified word is a palindrome.
    /// </summary>
    /// <param name="word">The word to check for palindrome properties.</param>
    /// <returns>
    /// <c>true</c> if the specified word is a palindrome; otherwise, <c>false</c>.
    /// </returns>
    /// <remarks>
    /// A palindrome is a word that reads the same backward as forward. 
    /// This method performs a case-sensitive comparison.
    /// </remarks>
    public static bool IsPalindrome(string word)
    {
        int left = 0;
        int right = word.Length - 1;

        while (left < right)
        {
            if (word[left] != word[right])
            {
                return false;
            }
            left++;
            right--;
        }

        return true;
    }
}