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