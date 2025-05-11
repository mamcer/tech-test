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
}